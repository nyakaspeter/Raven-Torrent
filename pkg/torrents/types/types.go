package types

type MovieTorrent struct {
	Hash     string `json:"hash"`
	Quality  string `json:"quality"`
	Size     string `json:"size"`
	Provider string `json:"provider"`
	Lang     string `json:"lang"`
	Title    string `json:"title"`
	Seeds    string `json:"seeds"`
	Peers    string `json:"peers"`
	Magnet   string `json:"magnet"`
	Torrent  string `json:"torrent"`
}

type ShowTorrent struct {
	Hash     string `json:"hash"`
	Quality  string `json:"quality"`
	Season   string `json:"season"`
	Episode  string `json:"episode"`
	Size     string `json:"size"`
	Provider string `json:"provider"`
	Lang     string `json:"lang"`
	Title    string `json:"title"`
	Seeds    string `json:"seeds"`
	Peers    string `json:"peers"`
	Magnet   string `json:"magnet"`
	Torrent  string `json:"torrent"`
}

type MovieParams struct {
	ImdbId     string `json:"imdbId"`
	SearchText string `json:"searchText"`
}

type ShowParams struct {
	ImdbId     string `json:"imdbId"`
	SearchText string `json:"searchText"`
	Season     string `json:"season"`
	Episode    string `json:"episode"`
}

type SourceParams struct {
	Jackett  JackettParams  `json:"jackett"`
	Yts      YtsParams      `json:"yts"`
	Itorrent ItorrentParams `json:"itorrent"`
	X1337x   X1337xParams   `json:"x1337x"`
	Eztv     EztvParams     `json:"eztv"`
}

type JackettParams struct {
	Enabled bool   `json:"enabled"`
	Address string `json:"apiAddress"`
	ApiKey  string `json:"apiKey"`
}

type YtsParams struct {
	Enabled bool `json:"enabled"`
}

type ItorrentParams struct {
	Enabled bool `json:"enabled"`
}

type X1337xParams struct {
	Enabled bool `json:"enabled"`
}

type EztvParams struct {
	Enabled bool `json:"enabled"`
}
