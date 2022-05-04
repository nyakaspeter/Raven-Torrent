package main

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	//"log"
	"github.com/anacrolix/torrent"
	"github.com/dustin/go-humanize"
	"github.com/oz/osdb"
	out "github.com/silentmurdock/wrserver/providers/output"
)

type messageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type torrentFilesResponse struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Length string `json:"length"`
}

type torrentListResponse struct {
	Name   string `json:"name"`
	Hash   string `json:"hash"`
	Length string `json:"length"`
}

type torrentFilesResultsResponse struct {
	Success bool                   `json:"success"`
	Hash    string                 `json:"hash"`
	Results []torrentFilesResponse `json:"results"`
}

type torrentListResultsResponse struct {
	Success bool                  `json:"success"`
	Results []torrentListResponse `json:"results"`
}

type torrentStatsResponse struct {
	Success     bool   `json:"success"`
	DownSpeed   string `json:"downspeed"`
	DownData    string `json:"downdata"`
	DownPercent string `json:"downpercent"`
	FullData    string `json:"fulldata"`
	Peers       string `json:"peers"`
}

type subtitleFilesResponse struct {
	Lang         string `json:"lang"`
	SubtitleName string `json:"subtitlename"`
	ReleaseName  string `json:"releasename"`
	SubFormat    string `json:"subformat"`
	SubEncoding  string `json:"subencoding"`
	SubData      string `json:"subdata"`
	VttData      string `json:"vttdata"`
}

type subtitleFilesResultsResponse struct {
	Success bool                    `json:"success"`
	Results []subtitleFilesResponse `json:"results"`
}

type movieMagnetLinksResponse struct {
	Success bool                    `json:"success"`
	Results []out.OutputMovieStruct `json:"results"`
}

type showMagnetLinksResponse struct {
	Success bool                   `json:"success"`
	Results []out.OutputShowStruct `json:"results"`
}

type mediaRenderersResponse struct {
	Success bool            `json:"success"`
	Results []mediaRenderer `json:"results"`
}

type mediaRenderer struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func serverInfo() string {
	message := messageResponse{
		Success: true,
		Message: "White Raven Server v" + version,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func serverStop() string {
	message := messageResponse{
		Success: true,
		Message: "Server stopped.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func restartTorrentClient() string {
	message := messageResponse{
		Success: true,
		Message: "Restart torrent client.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func torrentFilesList(address string, infohash string, files []*torrent.File) string {
	sortFiles(files)

	var results []torrentFilesResponse

	for _, f := range files {
		result := torrentFilesResponse{
			Name:   f.DisplayPath(),
			Url:    "http://" + address + "/api/get/" + f.Torrent().InfoHash().String() + "/" + base64.StdEncoding.EncodeToString([]byte(f.DisplayPath())),
			Length: strconv.FormatInt(f.FileInfo().Length, 10),
		}

		results = append(results, result)
	}

	message := torrentFilesResultsResponse{
		Success: true,
		Hash:    infohash,
		Results: results,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func onlyOneTorrent() string {
	message := messageResponse{
		Success: false,
		Message: "Only one torrent stream allowed at a time.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func failedToAddTorrent() string {
	message := messageResponse{
		Success: false,
		Message: "Failed to add torrent.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func deleteTorrent() string {
	message := messageResponse{
		Success: true,
		Message: "Torrent deleted.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func deleteAllTorrent() string {
	message := messageResponse{
		Success: true,
		Message: "All torrents have been deleted.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func torrentNotFound() string {
	message := messageResponse{
		Success: false,
		Message: "Torrent not found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func noActiveTorrentFound() string {
	message := messageResponse{
		Success: false,
		Message: "No active torrents found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func showAllTorrent() string {
	var results []torrentListResponse

	for _, thistorrent := range torrents {
		result := torrentListResponse{
			Name:   thistorrent.torrent.Name(),
			Hash:   thistorrent.torrent.InfoHash().String(),
			Length: strconv.FormatInt(thistorrent.torrent.Length(), 10),
		}

		results = append(results, result)
	}

	message := torrentListResultsResponse{
		Success: true,
		Results: results,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func downloadStats(address string, torr *torrent.Torrent) string {
	currentProgress := torr.BytesCompleted()

	torrWorkTime := time.Now()
	torrDivTime := torrWorkTime.Sub(torrents[torr.InfoHash().String()].prevtime).Seconds()
	if uint64(torrDivTime) <= 0 {
		torrDivTime = 1
	}
	torrents[torr.InfoHash().String()].prevtime = torrWorkTime

	downloadSpeed := humanize.Bytes(uint64(currentProgress-torrents[torr.InfoHash().String()].progress)/uint64(torrDivTime)) + "/s"
	torrents[torr.InfoHash().String()].progress = currentProgress

	complete := humanize.Bytes(uint64(currentProgress))
	percent := humanize.FormatFloat("#.", float64(currentProgress)/float64(torr.Info().TotalLength())*100)
	size := humanize.Bytes(uint64(torr.Info().TotalLength()))
	peers := strconv.Itoa(torr.Stats().ActivePeers) + "/" + strconv.Itoa(torr.Stats().TotalPeers)

	//log.Println("Download speed:", downloadSpeed, "Downloaded data:", complete, "Total length:", size)
	//log.Println("Active peers:", torr.Stats().ActivePeers, "Total peers", torr.Stats().TotalPeers, "Percent:", percent)

	message := torrentStatsResponse{
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

func resourceNotFound() string {
	message := messageResponse{
		Success: false,
		Message: "The resource you requested could not be found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func invalidBase64Path() string {
	message := messageResponse{
		Success: false,
		Message: "Invalid base64 path.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func failedToConnectToOpenSubtitles() string {
	message := messageResponse{
		Success: false,
		Message: "Failed to connect to opensubtitles.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func noSubtitlesFound() string {
	message := messageResponse{
		Success: false,
		Message: "No subtitles found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func subtitleFilesList(address string, files osdb.Subtitles, lang string) string {
	sortSubtitleFiles(files, lang)

	var results []subtitleFilesResponse

	for _, f := range files {
		if f.SubFormat == "srt" {
			workSubFileName := strings.ReplaceAll(f.SubFileName, "\"", "")
			workSubFileName = strings.ReplaceAll(workSubFileName, "\\", "")

			workMovieReleaseName := strings.ReplaceAll(f.MovieReleaseName, "\"", "")
			workMovieReleaseName = strings.ReplaceAll(workMovieReleaseName, "\\", "")

			result := subtitleFilesResponse{
				Lang:         f.ISO639,
				SubtitleName: workSubFileName,
				ReleaseName:  workMovieReleaseName,
				SubFormat:    f.SubFormat,
				SubEncoding:  f.SubEncoding,
				SubData:      "http://" + address + "/api/getsubtitle/" + base64.URLEncoding.EncodeToString([]byte(f.ZipDownloadLink)) + "/encode/" + f.SubEncoding + "/subtitle.srt",
				VttData:      "http://" + address + "/api/getsubtitle/" + base64.URLEncoding.EncodeToString([]byte(f.ZipDownloadLink)) + "/encode/" + f.SubEncoding + "/subtitle.vtt",
			}

			results = append(results, result)
		}
	}

	message := subtitleFilesResultsResponse{
		Success: true,
		Results: results,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func failedToLoadSubtitle() string {
	message := messageResponse{
		Success: false,
		Message: "Failed to load the subtitle.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func displayMovieMagnetLinks(results []out.OutputMovieStruct) string {
	message := movieMagnetLinksResponse{
		Success: true,
		Results: results,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func displayShowMagnetLinks(results []out.OutputShowStruct) string {
	message := showMagnetLinksResponse{
		Success: true,
		Results: results,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func noMagnetLinksFound() string {
	message := messageResponse{
		Success: false,
		Message: "No magnet links found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func outputTmdbData(data string) string {
	return "{\"success\":true,\"results\":[" + data + "]}"
}

func noTmdbDataFound() string {
	message := messageResponse{
		Success: false,
		Message: "No TMDB data found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func outputTvMazeData(data string) string {
	return "{\"success\":true,\"results\":" + data + "}"
}

func noTvMazeDataFound() string {
	message := messageResponse{
		Success: false,
		Message: "No TVMaze data found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func noMediaRenderersFound() string {
	message := messageResponse{
		Success: false,
		Message: "No media renderers found.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func mediaRenderersList(renderers []mediaRenderer) string {
	message := mediaRenderersResponse{
		Success: true,
		Results: renderers,
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func castingFailed() string {
	message := messageResponse{
		Success: false,
		Message: "Casting to device failed.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}

func playerFailed() string {
	message := messageResponse{
		Success: false,
		Message: "Failed to open media player.",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
