package v0

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
)

func GetTorrentFile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if d, err := base64.StdEncoding.DecodeString(vars["base64path"]); err == nil {
			if t, ok := torrentclient.ActiveTorrents[vars["hash"]]; ok {

				idx := torrentclient.GetFileIndexByPath(string(d), t.Torrent.Files())
				if idx != -1 {
					file := t.Torrent.Files()[idx]

					path := file.DisplayPath()
					log.Println("Downloading torrent:", vars["hash"])

					torrentclient.IncreaseConnections(path, t)

					/*log.Println("Calculate Opensubtitles hash...")
					fileHash := calculateOpensubtitlesHash(file)
					log.Println("Opensubtitles hash calculated:", fileHash)*/

					torrentclient.ServeTorrentFile(w, r, file)
					//stop downloading the file when no connections left
					if torrentclient.DecreaseConnections(path, t) <= 0 {
						torrentclient.StopFileDownload(file)
					}
				} else {
					http.Error(w, invalidBase64Path(), http.StatusNotFound)
					return
				}
			} else {
				http.Error(w, torrentNotFound(), http.StatusNotFound)
				log.Println("Unknown torrent:", vars["hash"])
				return
			}
		} else {
			http.Error(w, invalidBase64Path(), http.StatusNotFound)
			log.Println(err)
			return
		}
	}
}

func invalidBase64Path() string {
	message := MessageResponse{
		Success: false,
		Message: "Invalid base64 path.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
