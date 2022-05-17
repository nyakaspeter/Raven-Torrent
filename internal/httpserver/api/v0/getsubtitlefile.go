package v0

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/pkg/subtitles"
	subtitlestypes "github.com/nyakaspeter/raven-torrent/pkg/subtitles/types"
)

func GetSubtitleFile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if subtitleUrl, err := base64.StdEncoding.DecodeString(vars["base64path"]); err == nil {
			params := subtitlestypes.SubtitleParams{}
			params.Url = string(subtitleUrl)
			params.Encoding = vars["encode"]
			params.TargetType = vars["subtype"]

			contents := subtitles.GetSubtitleContents(params)

			if contents.Text == "" {
				http.Error(w, failedToLoadSubtitle(), http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Disposition", contents.ContentDisposition)
			w.Header().Set("Content-Type", contents.ContentType)
			io.WriteString(w, contents.Text)
		} else {
			http.Error(w, failedToLoadSubtitle(), http.StatusNotFound)
		}
	}
}

func failedToLoadSubtitle() string {
	message := MessageResponse{
		Success: false,
		Message: "Failed to load the subtitle.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
