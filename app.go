package main

import (
	"context"

	"github.com/nyakaspeter/raven-torrent/internal/httpserver"
	"github.com/nyakaspeter/raven-torrent/internal/settings"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
	torrentclienttypes "github.com/nyakaspeter/raven-torrent/internal/torrentclient/types"
	"github.com/nyakaspeter/raven-torrent/pkg/dlnacast"
	dlnacasttypes "github.com/nyakaspeter/raven-torrent/pkg/dlnacast/types"
	"github.com/nyakaspeter/raven-torrent/pkg/mediainfo"
	mediainfotypes "github.com/nyakaspeter/raven-torrent/pkg/mediainfo/types"
	"github.com/nyakaspeter/raven-torrent/pkg/mediaplayer"
	mediaplayertypes "github.com/nyakaspeter/raven-torrent/pkg/mediaplayer/types"
	"github.com/nyakaspeter/raven-torrent/pkg/subtitles"
	subtitlestypes "github.com/nyakaspeter/raven-torrent/pkg/subtitles/types"
	"github.com/nyakaspeter/raven-torrent/pkg/torrents"
	torrentstypes "github.com/nyakaspeter/raven-torrent/pkg/torrents/types"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	settings.Init()
	torrentclient.StartTorrentClient()
	httpserver.StartHttpServer(nil)
}

func (a *App) DiscoverMovies(params mediainfotypes.MovieDiscoverParams, language string, page int) mediainfotypes.MovieResults {
	return mediainfo.DiscoverMovies(params, language, page)
}

func (a *App) DiscoverShows(params mediainfotypes.ShowDiscoverParams, language string, page int) mediainfotypes.ShowResults {
	return mediainfo.DiscoverShows(params, language, page)
}

func (a *App) SearchMovies(title string, language string, page int) mediainfotypes.MovieResults {
	return mediainfo.SearchMovies(title, language, page)
}

func (a *App) SearchShows(title string, language string, page int) mediainfotypes.ShowResults {
	return mediainfo.SearchShows(title, language, page)
}

func (a *App) GetMovieInfo(tmdbId int, language string) mediainfotypes.MovieInfo {
	return mediainfo.GetMovieInfo(tmdbId, language)
}

func (a *App) GetShowInfo(tmdbId int, language string) mediainfotypes.ShowInfo {
	return mediainfo.GetShowInfo(tmdbId, language)
}

func (a *App) GetShowSeason(tmdbId int, seasonNumber int, language string) mediainfotypes.SeasonInfo {
	return mediainfo.GetShowSeason(tmdbId, seasonNumber, language)
}

func (a *App) GetMovieTorrents(movie torrentstypes.MovieParams, sources torrentstypes.SourceParams) []torrentstypes.MovieTorrent {
	return torrents.GetMovieTorrents(movie, sources)
}

func (a *App) GetShowTorrents(show torrentstypes.ShowParams, sources torrentstypes.SourceParams) []torrentstypes.ShowTorrent {
	return torrents.GetShowTorrents(show, sources)
}

func (a *App) AddTorrent(uri string) torrentclienttypes.TorrentInfo {
	return torrentclient.AddTorrent(uri)
}

func (a *App) RemoveTorrent(hash string) error {
	return torrentclient.RemoveTorrent(hash)
}

func (a *App) GetActiveTorrents() []torrentclienttypes.TorrentInfo {
	return torrentclient.GetActiveTorrents()
}

func (a *App) GetSubtitles(movie subtitlestypes.MediaParams, languages []string) []subtitlestypes.SubtitleFile {
	return subtitles.GetSubtitles(movie, languages)
}

func (a *App) GetSubtitlesForEpisode(movie subtitlestypes.MediaParams, episode subtitlestypes.EpisodeParams, languages []string) []subtitlestypes.SubtitleFile {
	return subtitles.GetSubtitlesForEpisode(movie, episode, languages)
}

func (a *App) GetMediaDevices() []dlnacasttypes.MediaDevice {
	return dlnacast.GetMediaDevices()
}

func (a *App) CastMediaToDevice(media dlnacasttypes.MediaParams, deviceLocation string) error {
	return dlnacast.CastMediaToDevice(media, deviceLocation)
}

func (a *App) StartMediaPlayer(params mediaplayertypes.MediaPlayerParams) error {
	return mediaplayer.StartMediaPlayer(params)
}
