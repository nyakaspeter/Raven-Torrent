package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nyakaspeter/raven-torrent/internal/api/v0/responses"
)

func NotFound() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, resourceNotFound(), http.StatusNotFound)
	})
}

func resourceNotFound() string {
	message := responses.MessageResponse{
		Success: false,
		Message: "The resource you requested could not be found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
