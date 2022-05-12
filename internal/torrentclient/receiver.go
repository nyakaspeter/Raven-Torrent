package torrentclient

import (
	"log"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
)

var receiverEnabled bool = false
var receiverResponse string = ""

func SetReceivedMagnet(magnet string) string {
	if receiverEnabled {
		log.Println("Received magnet link:", magnet)

		receiverResponse = magnet
		receiverEnabled = false
		return "ok"
	} else {
		return ""
	}
}

func SetReceivedTorrent(mi *metainfo.MetaInfo) string {
	if receiverEnabled {
		spec := torrent.TorrentSpecFromMetaInfo(mi)
		log.Println("Received torrent file:", spec.DisplayName)

		receivedTorrent = mi
		receiverResponse = spec.DisplayName
		receiverEnabled = false
		return "ok"
	} else {
		return ""
	}
}

func CheckReceiver(todo string) string {
	if todo == "start" {
		receiverEnabled = true
		receiverResponse = ""
		receivedTorrent = nil
		return "{\"response\":\"ok\"}"
	} else if todo == "check" {
		// Wait 3 second because Long Polling
		time.Sleep(3 * time.Second)
		return "{\"received\":\"" + receiverResponse + "\"}"
	} else if todo == "stop" {
		receiverEnabled = false
		return "{\"response\":\"ok\"}"
	} else {
		return "{\"response\":\"unknown\"}"
	}
}
