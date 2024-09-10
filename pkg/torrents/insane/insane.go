package insane

import (
	"fmt"
	"math"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/nyakaspeter/raven-torrent/internal/settings"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/types"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/utils"
	"golang.org/x/net/publicsuffix"
)

func GetMovieTorrentsByImdbId(imdb string, user string, password string, ch chan<- []types.MovieTorrent) {
	client, err := login(user, password)
	if err != nil {
		ch <- []types.MovieTorrent{}
		return
	}

	params := url.Values{
		"page":       {"0"},
		"search":     {imdb},
		"searchsort": {"imdb"},
		"searchtype": {"desc"},
		"torart":     {"tor"},
		"cat[]":      {"41", "27", "44", "42", "25", "45"},
	}

	link := fmt.Sprintf("https://newinsane.info/browse.php?%s", params.Encode())

	resp, err := client.Get(link)
	if err != nil {
		ch <- []types.MovieTorrent{}
		return
	}

	defer resp.Body.Close()

	torrents, err := parseMovieTorrents(resp)
	if err != nil {
		ch <- []types.MovieTorrent{}
		return
	}

	ch <- torrents
}

func GetMovieTorrentsByText(searchText string, user string, password string, ch chan<- []types.MovieTorrent) {
	client, err := login(user, password)
	if err != nil {
		ch <- []types.MovieTorrent{}
		return
	}

	params := url.Values{
		"page":       {"0"},
		"search":     {searchText},
		"searchsort": {"normal"},
		"searchtype": {"desc"},
		"torart":     {"tor"},
		"cat[]":      {"41", "27", "44", "42", "25", "45"},
	}

	link := fmt.Sprintf("https://newinsane.info/browse.php?%s", params.Encode())

	resp, err := client.Get(link)
	if err != nil {
		ch <- []types.MovieTorrent{}
		return
	}

	defer resp.Body.Close()

	torrents, err := parseMovieTorrents(resp)
	if err != nil {
		ch <- []types.MovieTorrent{}
		return
	}

	ch <- torrents
}

func GetShowTorrentsByImdbId(imdb string, season string, episode string, user string, password string, ch chan<- []types.ShowTorrent) {
	client, err := login(user, password)
	if err != nil {
		ch <- []types.ShowTorrent{}
		return
	}

	params := url.Values{
		"page":       {"0"},
		"search":     {imdb},
		"searchsort": {"imdb"},
		"searchtype": {"desc"},
		"torart":     {"tor"},
		"cat[]":      {"8", "40", "47", "7", "39", "46"},
	}

	link := fmt.Sprintf("https://newinsane.info/browse.php?%s", params.Encode())

	resp, err := client.Get(link)
	if err != nil {
		ch <- []types.ShowTorrent{}
		return
	}

	defer resp.Body.Close()

	torrents, err := parseShowTorrents(resp)
	if err != nil {
		ch <- []types.ShowTorrent{}
		return
	}

	matchingTorrents := []types.ShowTorrent{}
	for _, torrent := range torrents {
		if (torrent.Season == season || season == "0") && (torrent.Episode == episode || episode == "0") {
			matchingTorrents = append(matchingTorrents, torrent)
		}
	}

	ch <- matchingTorrents
}

func GetShowTorrentsByText(searchText string, season string, episode string, user string, password string, ch chan<- []types.ShowTorrent) {
	query := searchText + " "

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

	client, err := login(user, password)
	if err != nil {
		ch <- []types.ShowTorrent{}
		return
	}

	params := url.Values{
		"page":       {"0"},
		"search":     {query},
		"searchsort": {"normal"},
		"searchtype": {"desc"},
		"torart":     {"tor"},
		"cat[]":      {"8", "40", "47", "7", "39", "46"},
	}

	link := fmt.Sprintf("https://newinsane.info/browse.php?%s", params.Encode())

	resp, err := client.Get(link)
	if err != nil {
		ch <- []types.ShowTorrent{}
		return
	}

	defer resp.Body.Close()

	torrents, err := parseShowTorrents(resp)
	if err != nil {
		ch <- []types.ShowTorrent{}
		return
	}

	matchingTorrents := []types.ShowTorrent{}
	for _, torrent := range torrents {
		if (torrent.Season == season || season == "0") && (torrent.Episode == episode || episode == "0") {
			matchingTorrents = append(matchingTorrents, torrent)
		}
	}

	ch <- matchingTorrents
}

func login(user string, password string) (*http.Client, error) {
	insaneUser := *settings.InsaneUser
	if user != "" {
		insaneUser = user
	}

	insanePassword := *settings.InsanePassword
	if password != "" {
		insanePassword = password
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Jar: jar,
	}

	data := url.Values{
		"username": {insaneUser},
		"password": {insanePassword},
	}

	req, err := http.NewRequest("POST", "https://newinsane.info/login.php", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("login failed with status: %s", resp.Status)
	}

	return client, nil
}

func parseMovieTorrents(resp *http.Response) ([]types.MovieTorrent, error) {
	torrents := []types.MovieTorrent{}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("tr.torrentrow").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("a.torrentname").Attr("title")
		size := parseSize(s.Find("td.size").Text())
		seeds, _ := strconv.Atoi(s.Find("td.data > a:nth-of-type(1)").Text())
		peers, _ := strconv.Atoi(s.Find("td.data > a:nth-of-type(2)").Text())
		torrent, _ := s.Find("a.downloadicon").Attr("href")

		if name == "" || torrent == "" {
			return
		}

		torrents = append(torrents, types.MovieTorrent{
			Title:    name,
			Provider: "INSANE",
			Lang:     utils.GuessLanguageFromString(name),
			Quality:  utils.GuessQualityFromString(name),
			Size:     strconv.FormatInt(size, 10),
			Seeds:    strconv.Itoa(seeds),
			Peers:    strconv.Itoa(peers),
			Torrent:  torrent,
		})
	})

	return torrents, nil
}

func parseShowTorrents(resp *http.Response) ([]types.ShowTorrent, error) {
	torrents := []types.ShowTorrent{}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("tr.torrentrow").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("a.torrentname").Attr("title")
		size := parseSize(s.Find("td.size").Text())
		seeds, _ := strconv.Atoi(s.Find("td.data > a:nth-of-type(1)").Text())
		peers, _ := strconv.Atoi(s.Find("td.data > a:nth-of-type(2)").Text())
		torrent, _ := s.Find("a.downloadicon").Attr("href")

		if name == "" || torrent == "" {
			return
		}

		season, episode := utils.GuessSeasonEpisodeNumberFromString(name)

		torrents = append(torrents, types.ShowTorrent{
			Title:    name,
			Provider: "INSANE",
			Lang:     utils.GuessLanguageFromString(name),
			Quality:  utils.GuessQualityFromString(name),
			Size:     strconv.FormatInt(size, 10),
			Seeds:    strconv.Itoa(seeds),
			Peers:    strconv.Itoa(peers),
			Torrent:  torrent,
			Season:   season,
			Episode:  episode,
		})
	})

	return torrents, nil
}

func parseSize(sizeStr string) int64 {
	size := strings.Replace(strings.TrimSpace(sizeStr), ",", ".", 1)
	var bytes float64

	if strings.HasSuffix(size, "TiB") {
		bytes, _ = strconv.ParseFloat(strings.TrimSuffix(size, "TiB"), 64)
		bytes *= math.Pow(1024, 4)
	} else if strings.HasSuffix(size, "GiB") {
		bytes, _ = strconv.ParseFloat(strings.TrimSuffix(size, "GiB"), 64)
		bytes *= math.Pow(1024, 3)
	} else if strings.HasSuffix(size, "MiB") {
		bytes, _ = strconv.ParseFloat(strings.TrimSuffix(size, "MiB"), 64)
		bytes *= math.Pow(1024, 2)
	} else if strings.HasSuffix(size, "KiB") {
		bytes, _ = strconv.ParseFloat(strings.TrimSuffix(size, "KiB"), 64)
		bytes *= 1024
	} else if strings.HasSuffix(size, "B") {
		bytes, _ = strconv.ParseFloat(strings.TrimSuffix(size, "B"), 64)
	}

	return int64(math.Ceil(bytes))
}
