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
	Url        string `json:"url"`
	Encoding   string `json:"encoding"`
	TargetType string `json:"targetType"`
}

type SubtitleContents struct {
	Text               string `json:"text"`
	ContentType        string `json:"contentType"`
	ContentDisposition string `json:"contentDisposition"`
}
