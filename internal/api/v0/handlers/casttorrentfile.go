package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/api/v0/responses"
	"github.com/nyakaspeter/raven-torrent/internal/dlnacast"
	"github.com/nyakaspeter/raven-torrent/pkg/utils"
)

func CastTorrentFile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var videoURL = ""
		var subtitlesURL = ""
		var videoTitle = "video"
		var host = utils.GetLocalIP()

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

		if params["video"] != nil && params["video"][0] != "" {
			videoURL = params["video"][0]
		}

		if params["subtitle"] != nil && params["subtitle"][0] != "" {
			subtitlesURL = params["subtitle"][0]
		}

		if params["title"] != nil && params["title"][0] != "" {
			videoTitle = params["title"][0]
		}

		transportURL, controlURL, err := dlnacast.GetAvTransportUrl(string(location))

		if err != nil || videoURL == "" {
			http.Error(w, failedCastingToDevice(), http.StatusNotFound)
			return
		}

		videoURL = strings.Replace(videoURL, "localhost", host, 1)
		subtitlesURL = strings.Replace(subtitlesURL, "localhost", host, 1)

		whereToListen := fmt.Sprintf("%s:%d", host, dlnacast.ServerPort)
		callbackURL := fmt.Sprintf("http://%s/callback", whereToListen)

		newPayload := &dlnacast.TVPayload{
			TransportURL: transportURL,
			ControlURL:   controlURL,
			CallbackURL:  callbackURL,
			VideoURL:     videoURL,
			SubtitlesURL: subtitlesURL,
			VideoTitle:   videoTitle,
		}

		serverStarted := make(chan struct{})

		if dlnacast.Server == nil {
			srv := dlnacast.CreateServer(whereToListen)
			dlnacast.Server = &srv
		} else {
			dlnacast.TvPayload.SendtoTV("Stop")
			dlnacast.Server.StopServer()
		}

		dlnacast.TvPayload = newPayload

		go func() {
			dlnacast.Server.StartServer(serverStarted, newPayload)
		}()
		// Wait for HTTP server to properly initialize
		<-serverStarted

		err = newPayload.SendtoTV("Play1")

		if err != nil || videoURL == "" {
			http.Error(w, failedCastingToDevice(), http.StatusNotFound)
			return
		}

		io.WriteString(w, responses.SuccessMessage())
	}
}

func failedCastingToDevice() string {
	message := responses.MessageResponse{
		Success: false,
		Message: "Casting to device failed.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
