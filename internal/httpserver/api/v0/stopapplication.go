package v0

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func StopApplication(quitSignal chan os.Signal) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, serverStopped())
		quitSignal <- os.Kill
	}
}

func serverStopped() string {
	message := MessageResponse{
		Success: true,
		Message: "Server stopped.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
