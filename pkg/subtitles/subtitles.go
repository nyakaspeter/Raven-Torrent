package subtitles

import (
	"context"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/TheForgotten69/go-opensubtitles/opensubtitles"
	subs "github.com/martinlindhe/subtitles"
	"github.com/nyakaspeter/raven-torrent/internal/settings"
	"github.com/nyakaspeter/raven-torrent/pkg/subtitles/types"
	"github.com/nyakaspeter/raven-torrent/pkg/utils"
)

func GetSubtitles(movie types.MediaParams, languages []string) []types.SubtitleFile {
	client := opensubtitles.NewClient(nil, "", opensubtitles.Credentials{
		Username: *settings.OpenSubtitlesUser,
		Password: *settings.OpenSubtitlesPassword,
	}, *settings.OpenSubtitlesKey)
	client.UserAgent = "Raven Torrent API"

	client, err := client.Connect()
	if err != nil {
		return []types.SubtitleFile{}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	// Fallback language always English
	if len(languages) == 0 {
		languages = append(languages, "eng")
	}

	subtitles, _, err := client.Find.Subtitles(ctx, &opensubtitles.SubtitlesOptions{Type: "movie", Query: movie.Title, ImdbID: movie.ImdbId, MovieHash: movie.FileHash, Languages: strings.Join(languages, ",")})
	if err != nil || len(subtitles.Items) == 0 {
		return []types.SubtitleFile{}
	}

	return subtitleFilesList(subtitles.Items, languages[0])
}

func GetSubtitlesForEpisode(show types.MediaParams, episode types.EpisodeParams, languages []string) []types.SubtitleFile {
	client := opensubtitles.NewClient(nil, "", opensubtitles.Credentials{
		Username: *settings.OpenSubtitlesUser,
		Password: *settings.OpenSubtitlesPassword,
	}, *settings.OpenSubtitlesKey)
	client.UserAgent = "Raven Torrent API"

	client, err := client.Connect()
	if err != nil {
		return []types.SubtitleFile{}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	feature, _, err := client.Find.Features(ctx, &opensubtitles.FeatureOptions{Query: show.Title, ImdbID: show.ImdbId, Type: "tvshow"})
	if err != nil || len(feature.Items) == 0 || feature.Items[0].Type != "tvshow" {
		return []types.SubtitleFile{}
	}

	episodeFeatureId := 0

	for _, season := range feature.Items[0].Attributes.Seasons {
		if season.SeasonNumber == int(episode.Season) {
			for _, ep := range season.Episodes {
				if ep.EpisodeNumber == int(episode.Episode) {
					episodeFeatureId = ep.FeatureID
					break
				}
			}
		}
	}

	// Fallback language always English
	if len(languages) == 0 {
		languages = append(languages, "eng")
	}

	subtitles, _, err := client.Find.Subtitles(ctx, &opensubtitles.SubtitlesOptions{ID: episodeFeatureId, MovieHash: show.FileHash, Languages: strings.Join(languages, ",")})
	if err != nil || len(subtitles.Items) == 0 {
		return []types.SubtitleFile{}
	}

	return subtitleFilesList(subtitles.Items, languages[0])
}

func GetSubtitleContents(params types.SubtitleParams) types.SubtitleContents {
	fileId, err := strconv.Atoi(params.FileId)
	if err != nil {
		return types.SubtitleContents{}
	}

	client := opensubtitles.NewClient(nil, "", opensubtitles.Credentials{
		Username: *settings.OpenSubtitlesUser,
		Password: *settings.OpenSubtitlesPassword,
	}, *settings.OpenSubtitlesKey)
	client.UserAgent = "Raven Torrent API"

	client, err = client.Connect()
	if err != nil {
		return types.SubtitleContents{}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	subtitleDownload, _, err := client.Download.Download(ctx, &opensubtitles.DownloadOptions{FileID: fileId})
	if err != nil {
		return types.SubtitleContents{}
	}

	var downloadClient http.Client
	resp, err := downloadClient.Get(subtitleDownload.Link)
	if err != nil {
		return types.SubtitleContents{}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return types.SubtitleContents{}
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.SubtitleContents{}
	}
	bodyString := string(bodyBytes)

	contents := types.SubtitleContents{}

	subtitle, err := subs.NewFromSRT(bodyString)
	if err != nil {
		return types.SubtitleContents{}
	}

	if params.TargetType == "srt" {
		contents.Text = subtitle.RemoveAds().AsSRT()
		contents.ContentType = "text/plain; charset=utf-8"
		contents.ContentDisposition = "filename=subtitle.srt"
		return contents
	} else if params.TargetType == "vtt" {
		contents.Text = subtitle.RemoveAds().AsVTT()
		contents.ContentType = "text/vtt; charset=utf-8"
		contents.ContentDisposition = "filename=subtitle.vtt"
		return contents
	} else {
		return types.SubtitleContents{}
	}
}

func subtitleFilesList(subtitles []opensubtitles.SubtitleData, firstLanguage string) []types.SubtitleFile {
	sortSubtitleFiles(subtitles, firstLanguage)

	var results []types.SubtitleFile

	for _, sub := range subtitles {
		if sub.Type == "subtitle" {
			for _, file := range sub.Attributes.Files {
				baseLink := "http://" + utils.GetLocalIP() + ":" + strconv.Itoa(*settings.Port) + "/subtitle/" + strconv.Itoa(file.FileID)

				result := types.SubtitleFile{
					Lang:         sub.Attributes.Language,
					SubtitleName: file.FileName,
					ReleaseName:  sub.Attributes.Release,
					SubData:      baseLink + "/srt",
					VttData:      baseLink + "/vtt",
				}

				results = append(results, result)
			}
		}
	}

	return results
}

func sortSubtitleFiles(subtitles []opensubtitles.SubtitleData, lang string) {
	sort.Slice(subtitles, func(i, j int) bool {
		return subtitles[i].Attributes.Language == lang
	})
}
