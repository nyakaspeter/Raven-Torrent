package rarbg

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	. "github.com/nyakaspeter/raven-torrent/pkg/torrents/output"
)

type sessionToken struct {
	Token     string
	StartTime time.Time
	MaxTime   float64
	WaitTime  int64
}

type apiTokenResponse struct {
	Token string `json:"token"`
}

type apiMovieResponse struct {
	TorrentResults []struct {
		Title    string `json:"title"`
		Category string `json:"category"`
		Download string `json:"download"`
		Seeders  int64  `json:"seeders"`
		Leechers int64  `json:"leechers"`
		Size     int64  `json:"size"`
	} `json:"torrent_results"`
	Error string `json:"error"`
}

type apiShowResponse struct {
	TorrentResults []struct {
		Title       string `json:"title"`
		Category    string `json:"category"`
		Download    string `json:"download"`
		Seeders     int64  `json:"seeders"`
		Leechers    int64  `json:"leechers"`
		Size        int64  `json:"size"`
		EpisodeInfo struct {
			SeasonNum string `json:"seasonnum"`
			EpNum     string `json:"epnum"`
			Title     string `json:"title"`
		} `json:"episode_info"`
	} `json:"torrent_results"`
	Error string `json:"error"`
}

var token = sessionToken{
	Token:     "",
	StartTime: time.Now(),
	MaxTime:   890,  // Seconds
	WaitTime:  2100, // Milliseconds
}

var tryCount = 0

func GetMovieTorrentsByImdbId(imdb string, ch chan<- []MovieTorrent) {
	if getToken() {
		time.Sleep(time.Millisecond * time.Duration(token.WaitTime))
		req, err := http.NewRequest("GET", "https://torrentapi.org/pubapi_v2.php?mode=search&app_id=whiteraven&format=json_extended&category=14;48;17;44;45;47;50;51;52;46;54&limit=50&min_seeders=1&sort=seeders&search_imdb="+imdb+"&token="+token.Token, nil)
		if err != nil {
			ch <- []MovieTorrent{}
			return
		}

		//req.Header.Set("User-Agent", UserAgent)
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			ch <- []MovieTorrent{}
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			ch <- []MovieTorrent{}
			return
		}

		response := apiMovieResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			ch <- []MovieTorrent{}
			return
		}

		if response.Error != "" && tryCount < 3 {
			//fmt.Println(response.Error, tryCount, token.Token)
			tryCount++
			GetMovieTorrentsByImdbId(imdb, ch)
		} else {
			tryCount = 0
		}

		if len(response.TorrentResults) == 0 {
			ch <- []MovieTorrent{}
			return
		}

		outputMovieData := []MovieTorrent{}

		for _, thistorrent := range response.TorrentResults {
			temp := MovieTorrent{
				Hash:     GetInfoHashFromMagnetLink(thistorrent.Download),
				Quality:  GuessQualityFromString(thistorrent.Title),
				Size:     strconv.FormatInt(thistorrent.Size, 10),
				Provider: "RARBG",
				Lang:     "en",
				Title:    thistorrent.Title,
				Seeds:    strconv.FormatInt(thistorrent.Seeders, 10),
				Peers:    strconv.FormatInt(thistorrent.Leechers, 10),
				Magnet:   thistorrent.Download,
			}
			outputMovieData = append(outputMovieData, temp)
		}

		ch <- outputMovieData
		return
	} else {
		ch <- []MovieTorrent{}
		return
	}
}

func GetShowTorrentsByImdbId(imdb string, season string, episode string, ch chan<- []ShowTorrent) {
	if getToken() {
		query := ""
		if season != "0" {
			if len(season) == 1 {
				query += "s0" + season
			} else {
				query += "s" + season
			}
		}
		if episode != "0" {
			if len(episode) == 1 {
				query += "e0" + episode
			} else {
				query += "e" + episode
			}
		}

		time.Sleep(time.Millisecond * time.Duration(token.WaitTime))
		req, err := http.NewRequest("GET", "https://torrentapi.org/pubapi_v2.php?mode=search&app_id=whiteraven&format=json_extended&category=18;41;49&limit=25&min_seeders=1&sort=seeders&search_imdb="+imdb+"&search_string="+query+"&token="+token.Token, nil)
		if err != nil {
			ch <- []ShowTorrent{}
			return
		}

		//req.Header.Set("User-Agent", UserAgent)
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			ch <- []ShowTorrent{}
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			ch <- []ShowTorrent{}
			return
		}

		response := apiShowResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			ch <- []ShowTorrent{}
			return
		}

		if response.Error != "" && tryCount < 3 {
			//fmt.Println(response.Error, tryCount, token.Token)
			tryCount++
			GetShowTorrentsByImdbId(imdb, season, episode, ch)
		} else {
			tryCount = 0
		}

		if len(response.TorrentResults) == 0 {
			ch <- []ShowTorrent{}
			return
		}

		outputShowData := []ShowTorrent{}

		for _, thistorrent := range response.TorrentResults {
			temp := ShowTorrent{
				Hash:     GetInfoHashFromMagnetLink(thistorrent.Download),
				Quality:  GuessQualityFromString(thistorrent.Title),
				Size:     strconv.FormatInt(thistorrent.Size, 10),
				Provider: "RARBG",
				Lang:     "en",
				Title:    thistorrent.Title,
				Seeds:    strconv.FormatInt(thistorrent.Seeders, 10),
				Peers:    strconv.FormatInt(thistorrent.Leechers, 10),
				Season:   thistorrent.EpisodeInfo.SeasonNum,
				Episode:  thistorrent.EpisodeInfo.EpNum,
				Magnet:   thistorrent.Download,
			}
			outputShowData = append(outputShowData, temp)
		}

		ch <- outputShowData
		return
	} else {
		ch <- []ShowTorrent{}
		return
	}
}

func getToken() bool {
	var url = "https://torrentapi.org/pubapi_v2.php?get_token=get_token&app_id=whiteraven"

	if token.Token == "" || (token.Token != "" && time.Since(token.StartTime).Seconds() > token.MaxTime) {
		//req.Header.Set("User-Agent", UserAgent)
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client := &http.Client{Transport: tr, Timeout: 10 * time.Second}

		res, err := client.Get(url)
		if err != nil {
			return false
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return false
		}

		var apiResponse apiTokenResponse
		json.Unmarshal(body, &apiResponse)

		if apiResponse.Token != "" {
			token.Token = apiResponse.Token
			token.StartTime = time.Now()
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}
