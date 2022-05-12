package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/nyakaspeter/raven-torrent/internal/api/v0/responses"
)

func About(version string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, serverInfo(version))
	}
}

func serverInfo(version string) string {
	message := responses.MessageResponse{
		Success: true,
		Message: "Raven Torrent v" + version,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
