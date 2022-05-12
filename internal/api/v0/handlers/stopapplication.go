package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/nyakaspeter/raven-torrent/internal/api/v0/responses"
)

func StopApplication(procQuit chan bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, serverStopped())
		if err == nil {
			go func() {
				time.Sleep(1 * time.Nanosecond)
				procQuit <- true
			}()
		} else {
			go func() {
				time.Sleep(1 * time.Nanosecond)
				procQuit <- true
			}()
		}
	}
}

func serverStopped() string {
	message := responses.MessageResponse{
		Success: true,
		Message: "Server stopped.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
