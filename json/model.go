package json

// PushNotificationRequest represents a pushNotification request
type PushNotificationRequest struct {
	Data PushNotificationRequestData `json:"data" `
}

// PushNotificationRequestData represents the a pushNotification request data
type PushNotificationRequestData struct {
	Type       string                     `json:"type" `
	Attributes PushNotificationAttributes `json:"attributes" `
}

// PushNotificationAttributes contains attributes for push notification
type PushNotificationAttributes struct {
	Data     map[string]string `json:"data"`
	Title    string            `json:"title"`
	Body     string            `json:"body"`
	ImageURL string            `json:"imageURL"`
	Tokens   []string          `json:"tokens"`
}
