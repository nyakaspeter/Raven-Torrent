package pt

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/anacrolix/torrent/metainfo"
	. "github.com/nyakaspeter/raven-torrent/pkg/torrents/output"
)

type apiMovieResponse struct {
	Torrents struct {
		Lang struct {
			Quality2160p struct {
				Url      string `json:"url"`
				Size     string `json:"size"`
				Provider string `json:"provider"`
				Seed     int64  `json:"seed"`
				Peer     int64  `json:"peer"`
			} `json:"2160p"`
			Quality1080p struct {
				Url      string `json:"url"`
				Size     string `json:"size"`
				Provider string `json:"provider"`
				Seed     int64  `json:"seed"`
				Peer     int64  `json:"peer"`
			} `json:"1080p"`
			Quality720p struct {
				Url      string `json:"url"`
				Size     string `json:"size"`
				Provider string `json:"provider"`
				Seed     int64  `json:"seed"`
				Peer     int64  `json:"peer"`
			} `json:"720p"`
			Quality480p struct {
				Url      string `json:"url"`
				Size     string `json:"size"`
				Provider string `json:"provider"`
				Seed     int64  `json:"seed"`
				Peer     int64  `json:"peer"`
			} `json:"480p"`
			Quality360p struct {
				Url      string `json:"url"`
				Size     string `json:"size"`
				Provider string `json:"provider"`
				Seed     int64  `json:"seed"`
				Peer     int64  `json:"peer"`
			} `json:"360p"`
		} `json:"en"`
	} `json:"torrents"`
	Title string `json:"title"`
}

type apiShowResponse struct {
	Episodes []struct {
		Torrents struct {
			Quality2160p struct {
				Provider string `json:"provider"`
				Url      string `json:"url"`
				Seeds    int64  `json:"seeds"`
				Peers    int64  `json:"peers"`
			} `json:"2160p"`
			Quality1080p struct {
				Provider string `json:"provider"`
				Url      string `json:"url"`
				Seeds    int64  `json:"seeds"`
				Peers    int64  `json:"peers"`
			} `json:"1080p"`
			Quality720p struct {
				Provider string `json:"provider"`
				Url      string `json:"url"`
				Seeds    int64  `json:"seeds"`
				Peers    int64  `json:"peers"`
			} `json:"720p"`
			Quality480p struct {
				Provider string `json:"provider"`
				Url      string `json:"url"`
				Seeds    int64  `json:"seeds"`
				Peers    int64  `json:"peers"`
			} `json:"480p"`
			Quality360p struct {
				Provider string `json:"provider"`
				Url      string `json:"url"`
				Seeds    int64  `json:"seeds"`
				Peers    int64  `json:"peers"`
			} `json:"360p"`
		} `json:"torrents"`
		Episode int64 `json:"episode"`
		Season  int64 `json:"season"`
	} `json:"episodes"`
	Title string `json:"title"`
}

const apiUrl = "https://popcorn-time.ga"

func GetMovieTorrentsByImdbId(imdb string, ch chan<- []MovieTorrent) {
	req, err := http.NewRequest("GET", apiUrl+"/movie/"+imdb, nil)
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

	outputMovieData := []MovieTorrent{}

	if response.Torrents.Lang.Quality360p.Url != "" {
		data, err := metainfo.ParseMagnetURI(response.Torrents.Lang.Quality360p.Url)
		if err == nil {
			temp := MovieTorrent{
				Hash:     data.InfoHash.String(),
				Quality:  "360p",
				Size:     response.Torrents.Lang.Quality360p.Size,
				Provider: response.Torrents.Lang.Quality360p.Provider,
				Lang:     "en",
				Title:    response.Title,
				Seeds:    strconv.FormatInt(response.Torrents.Lang.Quality360p.Seed, 10),
				Peers:    strconv.FormatInt(response.Torrents.Lang.Quality360p.Peer, 10),
				Magnet:   response.Torrents.Lang.Quality360p.Url,
			}
			outputMovieData = append(outputMovieData, temp)
		}
	}

	if response.Torrents.Lang.Quality480p.Url != "" {
		data, err := metainfo.ParseMagnetURI(response.Torrents.Lang.Quality480p.Url)
		if err == nil {
			temp := MovieTorrent{
				Hash:     data.InfoHash.String(),
				Quality:  "480p",
				Size:     response.Torrents.Lang.Quality480p.Size,
				Provider: response.Torrents.Lang.Quality480p.Provider,
				Lang:     "en",
				Title:    response.Title,
				Seeds:    strconv.FormatInt(response.Torrents.Lang.Quality480p.Seed, 10),
				Peers:    strconv.FormatInt(response.Torrents.Lang.Quality480p.Peer, 10),
				Magnet:   response.Torrents.Lang.Quality480p.Url,
			}
			outputMovieData = append(outputMovieData, temp)
		}
	}

	if response.Torrents.Lang.Quality720p.Url != "" {
		data, err := metainfo.ParseMagnetURI(response.Torrents.Lang.Quality720p.Url)
		if err == nil {
			temp := MovieTorrent{
				Hash:     data.InfoHash.String(),
				Quality:  "720p",
				Size:     response.Torrents.Lang.Quality720p.Size,
				Provider: response.Torrents.Lang.Quality720p.Provider,
				Lang:     "en",
				Title:    response.Title,
				Seeds:    strconv.FormatInt(response.Torrents.Lang.Quality720p.Seed, 10),
				Peers:    strconv.FormatInt(response.Torrents.Lang.Quality720p.Peer, 10),
				Magnet:   response.Torrents.Lang.Quality720p.Url,
			}
			outputMovieData = append(outputMovieData, temp)
		}
	}

	if response.Torrents.Lang.Quality1080p.Url != "" {
		data, err := metainfo.ParseMagnetURI(response.Torrents.Lang.Quality1080p.Url)
		if err == nil {
			temp := MovieTorrent{
				Hash:     data.InfoHash.String(),
				Quality:  "1080p",
				Size:     response.Torrents.Lang.Quality1080p.Size,
				Provider: response.Torrents.Lang.Quality1080p.Provider,
				Lang:     "en",
				Title:    response.Title,
				Seeds:    strconv.FormatInt(response.Torrents.Lang.Quality1080p.Seed, 10),
				Peers:    strconv.FormatInt(response.Torrents.Lang.Quality1080p.Peer, 10),
				Magnet:   response.Torrents.Lang.Quality1080p.Url,
			}
			outputMovieData = append(outputMovieData, temp)
		}
	}

	if response.Torrents.Lang.Quality2160p.Url != "" {
		data, err := metainfo.ParseMagnetURI(response.Torrents.Lang.Quality2160p.Url)
		if err == nil {
			temp := MovieTorrent{
				Hash:     data.InfoHash.String(),
				Quality:  "2160p",
				Size:     response.Torrents.Lang.Quality2160p.Size,
				Provider: response.Torrents.Lang.Quality2160p.Provider,
				Lang:     "en",
				Title:    response.Title,
				Seeds:    strconv.FormatInt(response.Torrents.Lang.Quality2160p.Seed, 10),
				Peers:    strconv.FormatInt(response.Torrents.Lang.Quality2160p.Peer, 10),
				Magnet:   response.Torrents.Lang.Quality2160p.Url,
			}
			outputMovieData = append(outputMovieData, temp)
		}
	}

	ch <- outputMovieData
}

func GetShowTorrentsByImdbId(imdb string, season string, episode string, ch chan<- []ShowTorrent) {
	req, err := http.NewRequest("GET", apiUrl+"/show/"+imdb, nil)
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

	outputShowData := []ShowTorrent{}

	for _, thisepisode := range response.Episodes {
		s := strconv.FormatInt(thisepisode.Season, 10)
		e := strconv.FormatInt(thisepisode.Episode, 10)

		if (s == season || season == "0") && (e == episode || episode == "0") {
			title := response.Title + " "

			if s != "0" {
				if len(s) == 1 {
					title += "S0" + s
				} else {
					title += "S" + s
				}
			}
			if e != "0" {
				if len(e) == 1 {
					title += "E0" + e
				} else {
					title += "E" + e
				}
			}

			if thisepisode.Torrents.Quality360p.Url != "" {
				data, err := metainfo.ParseMagnetURI(thisepisode.Torrents.Quality360p.Url)
				if err == nil {
					temp := ShowTorrent{
						Hash:     data.InfoHash.String(),
						Quality:  "360p",
						Size:     "0",
						Season:   strconv.FormatInt(thisepisode.Season, 10),
						Episode:  strconv.FormatInt(thisepisode.Episode, 10),
						Provider: thisepisode.Torrents.Quality360p.Provider,
						Title:    title,
						Seeds:    strconv.FormatInt(thisepisode.Torrents.Quality360p.Seeds, 10),
						Peers:    strconv.FormatInt(thisepisode.Torrents.Quality360p.Peers, 10),
						Magnet:   thisepisode.Torrents.Quality360p.Url,
					}
					outputShowData = append(outputShowData, temp)
				}
			}

			if thisepisode.Torrents.Quality480p.Url != "" {
				data, err := metainfo.ParseMagnetURI(thisepisode.Torrents.Quality480p.Url)
				if err == nil {
					temp := ShowTorrent{
						Hash:     data.InfoHash.String(),
						Quality:  "480p",
						Size:     "0",
						Season:   strconv.FormatInt(thisepisode.Season, 10),
						Episode:  strconv.FormatInt(thisepisode.Episode, 10),
						Provider: thisepisode.Torrents.Quality480p.Provider,
						Title:    title,
						Seeds:    strconv.FormatInt(thisepisode.Torrents.Quality480p.Seeds, 10),
						Peers:    strconv.FormatInt(thisepisode.Torrents.Quality480p.Peers, 10),
						Magnet:   thisepisode.Torrents.Quality480p.Url,
					}
					outputShowData = append(outputShowData, temp)
				}
			}

			if thisepisode.Torrents.Quality720p.Url != "" {
				data, err := metainfo.ParseMagnetURI(thisepisode.Torrents.Quality720p.Url)
				if err == nil {
					temp := ShowTorrent{
						Hash:     data.InfoHash.String(),
						Quality:  "720p",
						Size:     "0",
						Season:   strconv.FormatInt(thisepisode.Season, 10),
						Episode:  strconv.FormatInt(thisepisode.Episode, 10),
						Provider: thisepisode.Torrents.Quality720p.Provider,
						Title:    title,
						Seeds:    strconv.FormatInt(thisepisode.Torrents.Quality720p.Seeds, 10),
						Peers:    strconv.FormatInt(thisepisode.Torrents.Quality720p.Peers, 10),
						Magnet:   thisepisode.Torrents.Quality720p.Url,
					}
					outputShowData = append(outputShowData, temp)
				}
			}

			if thisepisode.Torrents.Quality1080p.Url != "" {
				data, err := metainfo.ParseMagnetURI(thisepisode.Torrents.Quality1080p.Url)
				if err == nil {
					temp := ShowTorrent{
						Hash:     data.InfoHash.String(),
						Quality:  "1080p",
						Size:     "0",
						Season:   strconv.FormatInt(thisepisode.Season, 10),
						Episode:  strconv.FormatInt(thisepisode.Episode, 10),
						Provider: thisepisode.Torrents.Quality1080p.Provider,
						Title:    title,
						Seeds:    strconv.FormatInt(thisepisode.Torrents.Quality1080p.Seeds, 10),
						Peers:    strconv.FormatInt(thisepisode.Torrents.Quality1080p.Peers, 10),
						Magnet:   thisepisode.Torrents.Quality1080p.Url,
					}
					outputShowData = append(outputShowData, temp)
				}
			}

			if thisepisode.Torrents.Quality2160p.Url != "" {
				data, err := metainfo.ParseMagnetURI(thisepisode.Torrents.Quality2160p.Url)
				if err == nil {
					temp := ShowTorrent{
						Hash:     data.InfoHash.String(),
						Quality:  "2160p",
						Size:     "0",
						Season:   strconv.FormatInt(thisepisode.Season, 10),
						Episode:  strconv.FormatInt(thisepisode.Episode, 10),
						Provider: thisepisode.Torrents.Quality2160p.Provider,
						Title:    title,
						Seeds:    strconv.FormatInt(thisepisode.Torrents.Quality2160p.Seeds, 10),
						Peers:    strconv.FormatInt(thisepisode.Torrents.Quality2160p.Peers, 10),
						Magnet:   thisepisode.Torrents.Quality2160p.Url,
					}
					outputShowData = append(outputShowData, temp)
				}
			}
		}
	}

	ch <- outputShowData
}
