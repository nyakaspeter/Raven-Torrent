package torrents

import (
	"sort"
	"strconv"
	"strings"

	"github.com/nyakaspeter/raven-torrent/pkg/torrents/eztv"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/itorrent"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/jackett"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/types"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/x1337x"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents/yts"
)

func GetMovieTorrents(movie types.MovieParams, sources types.SourceParams) []types.MovieTorrent {
	output := []types.MovieTorrent{}
	ch := make(chan []types.MovieTorrent)

	count := 0
	if sources.Jackett.Enabled {
		go jackett.GetMovieTorrentsByText(movie.SearchText, sources.Jackett.Address, sources.Jackett.ApiKey, ch)
		count++
	}
	if sources.X1337x.Enabled {
		go x1337x.GetMovieTorrentsByText(movie.SearchText, ch)
		count++
	}
	if movie.ImdbId != "" {
		if sources.Jackett.Enabled {
			go jackett.GetMovieTorrentsByImdbId(movie.ImdbId, sources.Jackett.Address, sources.Jackett.ApiKey, ch)
			count++
		}
		if sources.Yts.Enabled {
			go yts.GetMovieTorrentsByImdbId(movie.ImdbId, ch)
			count++
		}
		if sources.Itorrent.Enabled {
			go itorrent.GetMovieTorrentsByImdbId(movie.ImdbId, ch)
			count++
		}
	}

	for count > 0 {
		results := <-ch
		for _, result := range results {
			duplicate := false
			for _, outResult := range output {
				if (outResult.Hash != "" && strings.EqualFold(outResult.Hash, result.Hash)) ||
					(outResult.Torrent != "" && outResult.Provider == result.Provider && outResult.Title == result.Title) {
					duplicate = true
					if outResult.Size == "0" && result.Size != "0" {
						outResult.Size = result.Size
						outResult.Title = result.Title
					}
				}
			}

			if !duplicate {
				output = append(output, result)
			}
		}
		count--
	}

	// Sort by seeds in descending order
	sort.Slice(output, func(i, j int) bool {
		si, _ := strconv.ParseInt(output[i].Seeds, 10, 64)
		sj, _ := strconv.ParseInt(output[j].Seeds, 10, 64)
		return si > sj
	})

	return output
}

func GetShowTorrents(show types.ShowParams, sources types.SourceParams) []types.ShowTorrent {
	output := []types.ShowTorrent{}
	ch := make(chan []types.ShowTorrent)

	count := 0
	if sources.Jackett.Enabled {
		go jackett.GetShowTorrentsByText(show.SearchText, show.Season, show.Episode, sources.Jackett.Address, sources.Jackett.ApiKey, ch)
		count++
	}
	if sources.X1337x.Enabled {
		go x1337x.GetShowTorrentsByText(show.SearchText, show.Season, show.Episode, ch)
		count++
	}
	if show.ImdbId != "" {
		if sources.Jackett.Enabled {
			go jackett.GetShowTorrentsByImdbId(show.ImdbId, show.Season, show.Episode, sources.Jackett.Address, sources.Jackett.ApiKey, ch)
			count++
		}
		if sources.Eztv.Enabled {
			go eztv.GetShowTorrentsByImdbId(show.ImdbId, show.Season, show.Episode, ch)
			count++
		}
		if sources.Itorrent.Enabled {
			go itorrent.GetShowTorrentsByImdbId(show.ImdbId, show.Season, show.Episode, ch)
			count++
		}
	}

	for count > 0 {
		results := <-ch
		for _, result := range results {
			duplicate := false
			for _, outResult := range output {
				if (outResult.Hash != "" && strings.EqualFold(outResult.Hash, result.Hash)) ||
					(outResult.Torrent != "" && outResult.Provider == result.Provider && outResult.Title == result.Title) {
					duplicate = true
					if outResult.Size == "0" && result.Size != "0" {
						outResult.Size = result.Size
						outResult.Title = result.Title
					}
				}
			}

			if !duplicate {
				output = append(output, result)
			}
		}
		count--
	}

	// Sort by seeds in descending order
	sort.Slice(output, func(i, j int) bool {
		si, _ := strconv.ParseInt(output[i].Seeds, 10, 64)
		sj, _ := strconv.ParseInt(output[j].Seeds, 10, 64)
		return si > sj
	})

	return output
}
