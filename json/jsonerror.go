package json

// ErrorData represents ErrorData data
type ErrorData struct {
	Data Error `json:"data"`
}

// Error is object for application errors response
type Error struct {
	Code        string `json:"errorCode"`
	Description string `json:"errorDescription,omitempty"`
	Status      string `json:"httpStatus"`
}
