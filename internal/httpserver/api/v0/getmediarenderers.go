package v0

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/nyakaspeter/raven-torrent/pkg/dlnacast"
	dlnacasttypes "github.com/nyakaspeter/raven-torrent/pkg/dlnacast/types"
)

type MediaRenderersResponse struct {
	Success bool                        `json:"success"`
	Results []dlnacasttypes.MediaDevice `json:"results"`
}

// @Router /mediarenderers [get]
// @Summary Get list of available casting targets
// @Description
// @Tags Media playback
// @Success 200 {object} MediaRenderersResponse
// @Failure 404 {object} MessageResponse
func GetMediaRenderers() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Looking for media renderers...")
		devices := dlnacast.GetMediaDevices()

		if len(devices) == 0 {
			log.Println("No media renderers found.")
			http.Error(w, noMediaRenderersFound(), http.StatusNotFound)
			return
		}

		log.Println("Media renderers found.")
		io.WriteString(w, mediaRenderersList(devices))
	}
}

func mediaRenderersList(renderers []dlnacasttypes.MediaDevice) string {
	message := MediaRenderersResponse{
		Success: true,
		Results: renderers,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func noMediaRenderersFound() string {
	message := MessageResponse{
		Success: false,
		Message: "No media renderers found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
