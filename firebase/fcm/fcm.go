package fcm

import (
	"context"
	"strings"

	"firebase.google.com/go/v4/errorutils"
	"firebase.google.com/go/v4/messaging"
	firebase "github.com/vivek080/Go-PushNotifications/firebase"
	jsonV1 "github.com/vivek080/Go-PushNotifications/json"
)

type partialErr []error

// CheckAndSendMessage check if the notifications has to sent to single or mutliple devices.
func CheckAndSendMessage(msg *jsonV1.PushNotificationAttributes) error {
	fcmClient := firebase.FcmClient()
	if len(msg.Tokens) == 1 {
		_, err := Send(fcmClient, toFcmMessage(msg))
		if err != nil {
			return err
		}
		return nil
	}

	err := SendMulticast(fcmClient, toFcmMultiMessage(msg))
	if err != nil {
		return err
	}

	return nil
}

// toFcmMessage converts the input data in fcm message format.
func toFcmMessage(msg *jsonV1.PushNotificationAttributes) *messaging.Message {
	message := &messaging.Message{}
	message.Data = msg.Data
	message.Notification = &messaging.Notification{
		Body:     msg.Body,
		Title:    msg.Title,
		ImageURL: msg.ImageURL,
	}
	message.Token = msg.Tokens[0]

	return message
}

// toFcmMultiMessage converts the input data in fcm multicast message format.
func toFcmMultiMessage(msg *jsonV1.PushNotificationAttributes) *messaging.MulticastMessage {
	message := &messaging.MulticastMessage{}
	message.Tokens = msg.Tokens
	message.Data = msg.Data
	message.Notification = &messaging.Notification{
		Body:     msg.Body,
		Title:    msg.Title,
		ImageURL: msg.ImageURL,
	}
	return message
}

// Send sends a notification to a single device
func Send(client *messaging.Client, msg *messaging.Message) (string, error) {
	return client.Send(context.Background(), msg)
}

// SendMulticast sends a notification to multiple devices. If a partial error occours, it
// tries to resend the failed messages once more if it's eligible for retries.
func SendMulticast(client *messaging.Client, msg *messaging.MulticastMessage) error {
	br, err := client.SendMulticast(context.Background(), msg)
	if err != nil {
		return err
	}
	if br.SuccessCount == len(msg.Tokens) {
		return nil
	}

	failedTokens := []string{}
	// parital error occurred
	for i, resp := range br.Responses {
		if !resp.Success {
			if errorutils.IsUnavailable(resp.Error) {
				failedTokens = append(failedTokens, msg.Tokens[i])
			}
		}
	}
	msg.Tokens = failedTokens
	br, err = client.SendMulticast(context.Background(), msg)
	if err != nil {
		return err
	}
	var errs partialErr
	// we still have partial error. return now
	if br.FailureCount > 0 {
		for _, resp := range br.Responses {
			if !resp.Success {
				errs = append(errs, resp.Error)
			}
		}
	}

	return errs
}

// Error returns the error message
func (e partialErr) Error() string {
	s := []string{}
	for _, err := range e {
		s = append(s, err.Error())
	}

	return strings.Join(s, ",")
}

// IsPartialError returns true if this was a partial error
func IsPartialError(err error) bool {
	_, ok := err.(partialErr)
	return ok
}
