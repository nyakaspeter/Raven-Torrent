package types

type MediaDevice struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type MediaParams struct {
	VideoUrl    string `json:"videoUrl"`
	SubtitleUrl string `json:"subtitleUrl"`
	Title       string `json:"title"`
}
