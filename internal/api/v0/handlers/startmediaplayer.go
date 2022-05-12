package handlers

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/api/v0/responses"
)

func StartMediaPlayer() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		path, err := base64.StdEncoding.DecodeString(vars["base64path"])
		if err != nil {
			http.Error(w, failedToOpenMediaPlayer(), http.StatusNotFound)
			return
		}

		args, err := base64.StdEncoding.DecodeString(vars["base64args"])
		if err != nil {
			http.Error(w, failedToOpenMediaPlayer(), http.StatusNotFound)
			return
		}

		splitArgs := strings.Split(string(args), ",")

		cmd := exec.Command(string(path), splitArgs...)
		err = cmd.Run()
		if err != nil {
			http.Error(w, failedToOpenMediaPlayer(), http.StatusNotFound)
			return
		}

		io.WriteString(w, responses.SuccessMessage())
	}
}

func failedToOpenMediaPlayer() string {
	message := responses.MessageResponse{
		Success: false,
		Message: "Failed to open media player.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
