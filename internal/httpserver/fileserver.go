package httpserver

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
	"github.com/nyakaspeter/raven-torrent/pkg/subtitles"
	subtitlestypes "github.com/nyakaspeter/raven-torrent/pkg/subtitles/types"
)

func ServeTorrentFile() func(w http.ResponseWriter, r *http.Request) {
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
					torrentclient.ServeTorrentFile(w, r, file)

					//stop downloading the file when no connections left
					if torrentclient.DecreaseConnections(path, t) <= 0 {
						torrentclient.StopFileDownload(file)
					}
				} else {
					http.Error(w, "Invalid path", http.StatusNotFound)
					return
				}
			} else {
				log.Println("Unknown torrent:", vars["hash"])
				http.Error(w, "Unknown torrent", http.StatusNotFound)
				return
			}
		} else {
			log.Println(err)
			http.Error(w, "Invalid path", http.StatusNotFound)

			return
		}
	}
}

func ServeSubtitleFile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		params := subtitlestypes.SubtitleParams{}
		params.FileId = vars["fileId"]
		params.TargetType = vars["type"]

		contents := subtitles.GetSubtitleContents(params)

		if contents.Text == "" {
			http.Error(w, "Failed to load subtitle", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Disposition", contents.ContentDisposition)
		w.Header().Set("Content-Type", contents.ContentType)
		io.WriteString(w, contents.Text)
	}
}
