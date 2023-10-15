package v0

import (
	"encoding/json"
	"log"
	"net/http"
)

type MessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

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

func successMessage() string {
	message := MessageResponse{
		Success: true,
		Message: "OK",
	}

	messageString, _ := json.Marshal(message)

	log.Println("Success.")

	return string(messageString)
}
