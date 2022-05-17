package types

type TorrentFile struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Length string `json:"length"`
}

type TorrentInfo struct {
	Hash  string        `json:"hash"`
	Files []TorrentFile `json:"files"`
}
