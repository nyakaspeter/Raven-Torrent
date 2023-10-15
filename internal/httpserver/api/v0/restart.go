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

// @Router /restart/downrate/{downrate}/uprate/{uprate} [get]
// @Summary Restart torrent client with new bandwith limits
// @Description
// @Tags Torrent client
// @Param downrate path int true " "
// @Param uprate path int true " "
// @Success 200 {object} MessageResponse
// @Failure 404 {object} MessageResponse
func RestartTorrentClient(quitSignal chan os.Signal) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		log.Println("Restarting torrent client:", vars)

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

		io.WriteString(w, torrentClientRestarted())
	}
}

func torrentClientRestarted() string {
	message := MessageResponse{
		Success: true,
		Message: "Restarted torrent client.",
	}

	messageString, _ := json.Marshal(message)

	log.Println("Restarted torrent client.")

	return string(messageString)
}

func failedToSetLimits() string {
	message := MessageResponse{
		Success: false,
		Message: "Failed to set speed limits.",
	}

	messageString, _ := json.Marshal(message)

	log.Println("Failed to set speed limits.")

	return string(messageString)
}
