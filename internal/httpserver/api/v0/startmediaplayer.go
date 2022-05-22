package v0

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/pkg/mediaplayer"
	mediaplayertypes "github.com/nyakaspeter/raven-torrent/pkg/mediaplayer/types"
)

// @Router /startplayer/{base64path}/{base64args} [get]
// @Summary Launch media player application
// @Description
// @Tags Media playback
// @Param base64path path string true " "
// @Param base64args path string true " "
// @Success 200 {object} MessageResponse
// @Failure 404 {object} MessageResponse
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

		params := mediaplayertypes.MediaPlayerParams{}
		params.ExecutablePath = string(path)
		params.ExecutableArgs = string(args)

		err = mediaplayer.StartMediaPlayer(params)
		if err != nil {
			http.Error(w, failedToOpenMediaPlayer(), http.StatusNotFound)
			return
		}

		io.WriteString(w, successMessage())
	}
}

func failedToOpenMediaPlayer() string {
	message := MessageResponse{
		Success: false,
		Message: "Failed to open media player.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
