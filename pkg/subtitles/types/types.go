package types

type MediaParams struct {
	ImdbId   string `json:"imdbId"`
	Title    string `json:"title"`
	FileHash string `json:"fileHash"`
	FileSize int64  `json:"fileSize"`
}

type EpisodeParams struct {
	Season  int64 `json:"season"`
	Episode int64 `json:"episode"`
}

type SubtitleParams struct {
	FileId     string `json:"url"`
	TargetType string `json:"targetType"`
}

type SubtitleFile struct {
	Lang         string `json:"lang"`
	SubtitleName string `json:"subtitlename"`
	ReleaseName  string `json:"releasename"`
	SubData      string `json:"subdata"`
	VttData      string `json:"vttdata"`
}

type SubtitleContents struct {
	Text               string `json:"text"`
	ContentType        string `json:"contentType"`
	ContentDisposition string `json:"contentDisposition"`
}
