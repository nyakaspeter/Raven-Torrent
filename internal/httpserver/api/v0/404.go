package v0

import (
	"encoding/json"
	"net/http"
)

func NotFound() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, resourceNotFound(), http.StatusNotFound)
	})
}

func resourceNotFound() string {
	message := MessageResponse{
		Success: false,
		Message: "The resource you requested could not be found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
