package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	jsonV1 "github.com/vivek080/Go-PushNotifications/json"
)

// SendHTTPErrorResponse sends a HTTP response with a single error.
func SendHTTPErrorResponse(w http.ResponseWriter, status int, code string, description string) {

	var jerr jsonV1.ErrorData
	jerr.Data.Status = strconv.Itoa(status)
	jerr.Data.Code = code
	jerr.Data.Description = description

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(jerr)
}
