package v0

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// @Router /about [get]
// @Summary Get application details
// @Description
// @Tags General
// @Success 200 {object} MessageResponse
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

	log.Println("Returning server info.")

	return string(messageString)
}
