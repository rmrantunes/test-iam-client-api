package service

import (
	"encoding/json"
	"net/http"
)

type ErrorHandlerInput struct {
	Message        []string
	HttpStatusCode int
}

func StdHttpError(w http.ResponseWriter, input *ErrorHandlerInput) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(input.HttpStatusCode)
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"errorMessages": input.Message,
	})
}
