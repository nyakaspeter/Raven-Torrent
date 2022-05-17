package eztv

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/nyakaspeter/raven-torrent/pkg/torrents/types"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/utils"
)

type apiShowResponse struct {
	TorrentCount int64 `json:"torrents_count"`
	Limit        int64 `json:"limit"`
	Page         int64 `json:"page"`
	Torrents     []struct {
		Hash       string `json:"hash"`
		Filename   string `json:"filename"`
		Season     string `json:"season"`
		Episode    string `json:"episode"`
		SizeBytes  string `json:"size_bytes"`
		Title      string `json:"title"`
		Seeds      int64  `json:"seeds"`
		Peers      int64  `json:"peers"`
		MagnetUrl  string `json:"magnet_url"`
		TorrentUrl string `json:"torrent_url"`
	} `json:"torrents"`
}

func GetShowTorrentsByImdbId(imdb string, season string, episode string, ch chan<- []types.ShowTorrent) {
	id := make([]string, 1)
	id[0] = strings.TrimPrefix(imdb, "tt")

	req, err := http.NewRequest("GET", "https://eztv.re/api/get-torrents?imdb_id="+id[0]+"&limit=100&page=1", nil)
	if err != nil {
		ch <- []types.ShowTorrent{}
		return
	}

	//req.Header.Set("User-Agent", UserAgent)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		ch <- []types.ShowTorrent{}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- []types.ShowTorrent{}
		return
	}

	firstresponse := apiShowResponse{}
	err = json.Unmarshal(body, &firstresponse)
	if err != nil {
		ch <- []types.ShowTorrent{}
		return
	}

	if firstresponse.TorrentCount == 0 {
		ch <- []types.ShowTorrent{}
		return
	}

	response := []apiShowResponse{}
	response = append(response, firstresponse)

	var maxpage int
	maxpage = int(firstresponse.TorrentCount / firstresponse.Limit)
	if firstresponse.TorrentCount%firstresponse.Limit != 0 {
		maxpage += 1
	}

	innerCh := make(chan apiShowResponse)

	for i := 2; i <= maxpage; i++ {
		go scrapeData(id[0], i, innerCh)
	}

	for i := 2; i <= maxpage; i++ {
		temp := <-innerCh
		response = append(response, temp)
	}

	outputShowData := []types.ShowTorrent{}

	for _, responsedata := range response {
		for _, thistorrent := range responsedata.Torrents {
			if (thistorrent.Season == season || season == "0") && (thistorrent.Episode == episode || episode == "0") {

				quality := utils.GuessQualityFromString(thistorrent.Filename)

				temp := types.ShowTorrent{
					Hash:     thistorrent.Hash,
					Quality:  quality,
					Size:     thistorrent.SizeBytes,
					Season:   thistorrent.Season,
					Episode:  thistorrent.Episode,
					Title:    thistorrent.Title,
					Provider: "EZTV",
					Seeds:    strconv.FormatInt(thistorrent.Seeds, 10),
					Peers:    strconv.FormatInt(thistorrent.Peers, 10),
					Magnet:   thistorrent.MagnetUrl,
					Torrent:  thistorrent.TorrentUrl,
				}
				outputShowData = append(outputShowData, temp)
			}
		}
	}

	ch <- outputShowData
}

func scrapeData(imdb string, page int, innerCh chan<- apiShowResponse) {
	response := apiShowResponse{}

	req, err := http.NewRequest("GET", "https://eztv.re/api/get-torrents?imdb_id="+imdb+"&limit=100&page="+strconv.Itoa(page), nil)
	if err != nil {
		innerCh <- response
	}

	//req.Header.Set("User-Agent", UserAgent)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		innerCh <- response
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		innerCh <- response
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		innerCh <- response
	}

	if response.TorrentCount == 0 {
		innerCh <- response
	}

	innerCh <- response
}
