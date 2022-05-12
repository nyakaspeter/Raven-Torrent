package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/martinlindhe/subtitles"
	"github.com/nyakaspeter/raven-torrent/internal/api/v0/responses"
	"github.com/nyakaspeter/raven-torrent/pkg/utils"
)

func GetSubtitleFile(useragent string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if subtitleurl, err := base64.StdEncoding.DecodeString(vars["base64path"]); err == nil {

			zipContent, err := utils.FetchZip(string(subtitleurl), useragent)
			if err != nil {
				http.Error(w, failedToLoadSubtitle(), http.StatusNotFound)
				return
			}

			for _, f := range zipContent.File {
				if strings.HasSuffix(strings.ToLower(f.Name), ".srt") {
					fileHandler, err := f.Open()
					if err != nil {
						http.Error(w, failedToLoadSubtitle(), http.StatusNotFound)
						return
					}
					data, err := ioutil.ReadAll(fileHandler)
					if err != nil {
						http.Error(w, failedToLoadSubtitle(), http.StatusNotFound)
						return
					}
					fileHandler.Close()

					// Remove UTF BOM
					if data[0] == 0xef && data[1] == 0xbb && data[2] == 0xbf {
						data = bytes.Trim(data, "\xef\xbb\xbf")
					}

					srt := utils.DecodeData(data, vars["encode"])

					subtitle, err := subtitles.NewFromSRT(srt)
					if err != nil {
						http.Error(w, failedToLoadSubtitle(), http.StatusNotFound)
						return
					}

					if vars["subtype"] == "srt" {
						w.Header().Set("Content-Disposition", "filename=subtitle.srt")
						w.Header().Set("Content-Type", "text/plain; charset=utf-8")
						io.WriteString(w, subtitle.RemoveAds().AsSRT())
					} else if vars["subtype"] == "vtt" {
						w.Header().Set("Content-Disposition", "filename=subtitle.vtt")
						w.Header().Set("Content-Type", "text/vtt; charset=utf-8")
						io.WriteString(w, subtitle.RemoveAds().AsVTT())
					} else {
						http.Error(w, failedToLoadSubtitle(), http.StatusNotFound)
						return
					}

					break
				}
			}
		} else {
			http.Error(w, failedToLoadSubtitle(), http.StatusNotFound)
		}
	}
}

func failedToLoadSubtitle() string {
	message := responses.MessageResponse{
		Success: false,
		Message: "Failed to load the subtitle.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
