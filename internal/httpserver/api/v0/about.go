package v0

import (
	"encoding/json"
	"io"
	"net/http"
)

func About(version string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, serverInfo(version))
	}
}

func serverInfo(version string) string {
	message := MessageResponse{
		Success: true,
		Message: "Raven Torrent v" + version,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
