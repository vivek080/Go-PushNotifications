package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
	fcm "github.com/vivek080/Go-PushNotifications/firebase/fcm"
	jsonV1 "github.com/vivek080/Go-PushNotifications/json"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		sendPushNotification(w, r)
		return
	default:
		w.WriteHeader(400)
		return
	}
}

// sendPushNotification sends a Push Notification to the user devices.
func sendPushNotification(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var request jsonV1.PushNotificationRequest

	w.Header().Add("X-Request-ID", r.Header.Get("X-Request-ID"))
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())
		SendHTTPErrorResponse(w, http.StatusBadRequest, "Unable to read the request", err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		log.Print(err.Error())
		SendHTTPErrorResponse(w, http.StatusBadRequest, "Unable to unmarshal the request", err.Error())
		return
	}

	er := fcm.CheckAndSendMessage(&request.Data.Attributes)
	if er != nil {
		log.Print("Unable to Send Push Notification ", "err", er)
		SendHTTPErrorResponse(w, http.StatusInternalServerError, "Unable to send push notification", er.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
