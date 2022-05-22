package v0

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/dustin/go-humanize"
	"github.com/gorilla/mux"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
)

type TorrentStatsResponse struct {
	Success     bool   `json:"success"`
	DownSpeed   string `json:"downspeed"`
	DownData    string `json:"downdata"`
	DownPercent string `json:"downpercent"`
	FullData    string `json:"fulldata"`
	Peers       string `json:"peers"`
}

// @Router /stats/{hash} [get]
// @Summary Get torrent download stats
// @Description
// @Tags Torrent client
// @Param hash path string true " "
// @Success 200 {object} TorrentStatsResponse
// @Failure 404 {object} MessageResponse
func GetTorrentStats() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if t, ok := torrentclient.ActiveTorrents[vars["hash"]]; ok {
			log.Println("Check torrent stats:", vars["hash"])
			io.WriteString(w, downloadStats(r.Host, t.Torrent))
		} else {
			http.Error(w, torrentNotFound(), http.StatusNotFound)
			log.Println("Unknown torrent:", vars["hash"])
		}
	}
}

func downloadStats(address string, torr *torrent.Torrent) string {
	currentProgress := torr.BytesCompleted()

	torrWorkTime := time.Now()
	torrDivTime := torrWorkTime.Sub(torrentclient.ActiveTorrents[torr.InfoHash().String()].Prevtime).Seconds()
	if uint64(torrDivTime) <= 0 {
		torrDivTime = 1
	}
	torrentclient.ActiveTorrents[torr.InfoHash().String()].Prevtime = torrWorkTime

	downloadSpeed := humanize.Bytes(uint64(currentProgress-torrentclient.ActiveTorrents[torr.InfoHash().String()].Progress)/uint64(torrDivTime)) + "/s"
	torrentclient.ActiveTorrents[torr.InfoHash().String()].Progress = currentProgress

	complete := humanize.Bytes(uint64(currentProgress))
	percent := humanize.FormatFloat("#.", float64(currentProgress)/float64(torr.Info().TotalLength())*100)
	size := humanize.Bytes(uint64(torr.Info().TotalLength()))
	peers := strconv.Itoa(torr.Stats().ActivePeers) + "/" + strconv.Itoa(torr.Stats().TotalPeers)

	//log.Println("Download speed:", downloadSpeed, "Downloaded data:", complete, "Total length:", size)
	//log.Println("Active peers:", torr.Stats().ActivePeers, "Total peers", torr.Stats().TotalPeers, "Percent:", percent)

	message := TorrentStatsResponse{
		Success:     true,
		DownSpeed:   downloadSpeed,
		DownData:    complete,
		DownPercent: percent,
		FullData:    size,
		Peers:       peers,
	}

	messageString, _ := json.Marshal(message)

	// Wait 3 second because Long Polling
	time.Sleep(3 * time.Second)

	return string(messageString)
}
