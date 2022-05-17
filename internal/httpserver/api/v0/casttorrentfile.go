package v0

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/pkg/dlnacast"
	dlnacasttypes "github.com/nyakaspeter/raven-torrent/pkg/dlnacast/types"
)

func CastTorrentFile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		media := dlnacasttypes.MediaParams{}
		media.Title = "video"

		location, err := base64.StdEncoding.DecodeString(vars["base64location"])
		if err != nil {
			http.Error(w, failedCastingToDevice(), http.StatusNotFound)
			return
		}

		query, err := base64.StdEncoding.DecodeString(vars["base64query"])
		if err != nil {
			http.Error(w, failedCastingToDevice(), http.StatusNotFound)
			return
		}

		params, err := url.ParseQuery(string(query))
		if err != nil {
			http.Error(w, failedCastingToDevice(), http.StatusNotFound)
			return
		}

		if params["video"] != nil {
			media.VideoUrl = params["video"][0]
		}

		if params["subtitle"] != nil {
			media.SubtitleUrl = params["subtitle"][0]
		}

		if params["title"] != nil {
			media.Title = params["title"][0]
		}

		if media.VideoUrl == "" {
			http.Error(w, failedCastingToDevice(), http.StatusNotFound)
			return
		}

		err = dlnacast.CastMediaToDevice(media, string(location))
		if err != nil {
			http.Error(w, failedCastingToDevice(), http.StatusNotFound)
			return
		}

		io.WriteString(w, SuccessMessage())
	}
}

func failedCastingToDevice() string {
	message := MessageResponse{
		Success: false,
		Message: "Casting to device failed.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
