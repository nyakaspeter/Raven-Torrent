package tvmaze

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/nyakaspeter/raven-torrent/pkg/mediainfo/types"
)

type tvmazeIdResponse struct {
	Id int `json:"id"`
}

func GetEpisodes(showId types.ShowIds) ([]types.TvMazeEpisode, error) {
	if showId.TvMazeId == "" {
		if showId.TvdbId != "" {
			showId.TvMazeId = getTvMazeId("tvdb", showId.TvdbId)
		} else if showId.ImdbId != "" {
			showId.TvMazeId = getTvMazeId("imdb", showId.ImdbId)
		}
	}

	if showId.TvMazeId == "" {
		return []types.TvMazeEpisode{}, errors.New("id not found")
	}

	requesturl := "https://api.tvmaze.com/shows/" + showId.TvMazeId + "/episodes"

	req, err := http.NewRequest("GET", requesturl, nil)
	if err != nil {
		return []types.TvMazeEpisode{}, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return []types.TvMazeEpisode{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []types.TvMazeEpisode{}, err
	}

	var results []types.TvMazeEpisode
	err = json.Unmarshal(body, &results)
	if err != nil {
		return []types.TvMazeEpisode{}, err
	}

	return results, nil
}

func getTvMazeId(qtype string, id string) string {
	requesturl := ""

	if qtype == "tvdb" {
		requesturl = "https://api.tvmaze.com/lookup/shows?thetvdb=" + id
	} else {
		requesturl = "https://api.tvmaze.com/lookup/shows?imdb=" + id
	}

	req, err := http.NewRequest("GET", requesturl, nil)
	if err != nil {
		return ""
	}

	//req.Header.Set("User-Agent", UserAgent)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	var message tvmazeIdResponse
	err = json.Unmarshal(body, &message)
	if err != nil || string(body) == "null" || message.Id == 0 {
		return ""
	}

	return strconv.Itoa(message.Id)
}
