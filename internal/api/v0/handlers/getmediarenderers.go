package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/koron/go-ssdp"
	"github.com/nyakaspeter/raven-torrent/internal/api/v0/responses"
	"github.com/nyakaspeter/raven-torrent/internal/dlnacast"
)

type MediaRenderer struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type MediaRenderersResponse struct {
	Success bool            `json:"success"`
	Results []MediaRenderer `json:"results"`
}

func GetMediaRenderers() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		devices := []MediaRenderer{}
		list, err := ssdp.Search(ssdp.All, 1, "")

		if err != nil || len(list) == 0 {
			http.Error(w, noMediaRenderersFound(), http.StatusNotFound)
			log.Println("No media renderers found.")
		}

		for _, srv := range list {
			if srv.Type == "urn:schemas-upnp-org:service:AVTransport:1" {
				duplicate := false
				for _, device := range devices {
					if device.Location == srv.Location {
						duplicate = true
					}
				}

				if !duplicate {
					if friendlyName, err := dlnacast.GetDeviceFriendlyName(srv.Location); err != nil {
						devices = append(devices, MediaRenderer{Name: srv.Server, Location: srv.Location})
					} else {
						devices = append(devices, MediaRenderer{Name: friendlyName, Location: srv.Location})
					}
				}
			}
		}

		if len(devices) > 0 {
			io.WriteString(w, mediaRenderersList(devices))
		} else {
			http.Error(w, noMediaRenderersFound(), http.StatusNotFound)
			log.Println("No media renderers found.")
		}
	}
}

func mediaRenderersList(renderers []MediaRenderer) string {
	message := MediaRenderersResponse{
		Success: true,
		Results: renderers,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func noMediaRenderersFound() string {
	message := responses.MessageResponse{
		Success: false,
		Message: "No media renderers found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
