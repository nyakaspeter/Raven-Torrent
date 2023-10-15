package v0

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// @Router /stop [get]
// @Summary Shut down the application
// @Description
// @Tags General
// @Success 200 {object} MessageResponse
func StopApplication(quitSignal chan os.Signal) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, serverStopping())
		quitSignal <- os.Kill
	}
}

func serverStopping() string {
	message := MessageResponse{
		Success: true,
		Message: "Server stopping.",
	}

	messageString, _ := json.Marshal(message)

	log.Println("Stopping server.")

	return string(messageString)
}
