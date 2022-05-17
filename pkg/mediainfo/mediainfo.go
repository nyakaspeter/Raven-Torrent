package mediainfo

import (
	"github.com/nyakaspeter/raven-torrent/pkg/mediainfo/tmdb"
	"github.com/nyakaspeter/raven-torrent/pkg/mediainfo/tvmaze"
	"github.com/nyakaspeter/raven-torrent/pkg/mediainfo/types"
)

func DiscoverMovies(params types.MovieDiscoverParams, language string, page int) types.MovieResults {
	output, err := tmdb.DiscoverMovies(params, language, page)
	if err != nil {
		return types.MovieResults{}
	}

	return output
}

func DiscoverShows(params types.ShowDiscoverParams, language string, page int) types.ShowResults {
	output, err := tmdb.DiscoverShows(params, language, page)
	if err != nil {
		return types.ShowResults{}
	}

	return output
}

func SearchMovies(title string, language string, page int) types.MovieResults {
	output, err := tmdb.SearchMovies(title, language, page)
	if err != nil {
		return types.MovieResults{}
	}

	return output
}

func SearchShows(title string, language string, page int) types.ShowResults {
	output, err := tmdb.SearchShows(title, language, page)
	if err != nil {
		return types.ShowResults{}
	}

	return output
}

func GetMovieInfo(tmdbId int, language string) types.MovieInfo {
	output, err := tmdb.GetMovieInfo(tmdbId, language)
	if err != nil {
		return types.MovieInfo{}
	}

	return output
}

func GetShowInfo(tmdbId int, language string) types.ShowInfo {
	output, err := tmdb.GetShowInfo(tmdbId, language)
	if err != nil {
		return types.ShowInfo{}
	}

	return output
}

func GetShowSeason(tmdbId int, seasonNumber int, language string) types.SeasonInfo {
	output, err := tmdb.GetShowSeason(tmdbId, seasonNumber, language)
	if err != nil {
		return types.SeasonInfo{}
	}

	return output
}

func GetShowEpisodes(showId types.ShowIds) []types.TvMazeEpisode {
	output, err := tvmaze.GetEpisodes(showId)
	if err != nil {
		return []types.TvMazeEpisode{}
	}

	return output
}
