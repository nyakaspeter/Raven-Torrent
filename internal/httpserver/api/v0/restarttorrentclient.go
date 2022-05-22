package v0

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/settings"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
)

func RestartTorrentClient(quitSignal chan os.Signal) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		downrate, err := strconv.Atoi(vars["downrate"])
		if err != nil {
			http.Error(w, failedToSetLimits(), http.StatusBadRequest)
		}

		uprate, err := strconv.Atoi(vars["uprate"])
		if err != nil {
			http.Error(w, failedToSetLimits(), http.StatusBadRequest)
		}

		torrentclient.StopTorrentClient()

		*settings.DownloadRate = downrate
		*settings.UploadRate = uprate

		_, err = torrentclient.StartTorrentClient()
		if err != nil {
			log.Println("Failed to restart torrent client.")
			quitSignal <- os.Kill
		}

		log.Println("Restarted torrent client.")
		io.WriteString(w, torrentClientRestarted())
	}
}

func torrentClientRestarted() string {
	message := MessageResponse{
		Success: true,
		Message: "Restarted torrent client.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func failedToSetLimits() string {
	message := MessageResponse{
		Success: false,
		Message: "Failed to set speed limits.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
