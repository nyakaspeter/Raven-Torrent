package types

import (
	"time"

	"github.com/anacrolix/torrent"
)

// Torrent lock structure
type TorrentLeaf struct {
	Torrent     *torrent.Torrent
	Progress    int64          // Downoad stats measurement
	Prevtime    time.Time      // Previous time for progress calculation
	FileClients map[string]int // Count active connections
}

type TorrentFile struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Length string `json:"length"`
}

type TorrentInfo struct {
	Name   string        `json:"name"`
	Hash   string        `json:"hash"`
	Length string        `json:"length"`
	Files  []TorrentFile `json:"files"`
}
