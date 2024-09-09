package httpserver

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	v0 "github.com/nyakaspeter/raven-torrent/internal/httpserver/api/v0"
	"github.com/nyakaspeter/raven-torrent/internal/settings"
)

const version = "0.7.0"

// @Title Raven Torrent API
// @Version 0.7.0
// @Host localhost:9000
// @BasePath /api/v0
func routesHandler() http.Handler {
	router := mux.NewRouter()
	router.SkipClean(true)

	router.HandleFunc("/file/{hash}/{base64path}", ServeTorrentFile())
	router.HandleFunc("/subtitle/{fileId}/{type}", ServeSubtitleFile())

	if *settings.EnableReceiver {
		router.HandleFunc("/", ReceiverPage())
		router.HandleFunc("/receiver/{todo}", ReceiveTorrent())
		router.HandleFunc("/websocket", Websocket(quitSignal))
	}

	if *settings.EnableSwaggerUi {
		enableSwagger(router)
	}

	api := router.PathPrefix("/api")
	apiV0 := api.PathPrefix("/v0").Subrouter()

	apiV0.HandleFunc("/tmdbdiscover/type/movie/genretype/{genretype}/sort/{sort}/date/{date}/lang/{lang}/page/{page}", v0.DiscoverMovies())
	apiV0.HandleFunc("/tmdbdiscover/type/tv/genretype/{genretype}/sort/{sort}/date/{date}/lang/{lang}/page/{page}", v0.DiscoverShows())
	apiV0.HandleFunc("/tmdbsearch/type/movie/lang/{lang}/page/{page}/text/{text}", v0.SearchMovies())
	apiV0.HandleFunc("/tmdbsearch/type/tv/lang/{lang}/page/{page}/text/{text}", v0.SearchShows())
	apiV0.HandleFunc("/tmdbinfo/type/movie/tmdbid/{tmdbid}/lang/{lang}", v0.GetMovieInfo())
	apiV0.HandleFunc("/tmdbinfo/type/tv/tmdbid/{tmdbid}/lang/{lang}", v0.GetShowInfo())
	apiV0.HandleFunc("/tvmazeepisodes/tvdb/{tvdb}/imdb/{imdb}", v0.GetShowEpisodesByImdbAndTvdb())
	apiV0.HandleFunc("/tvmazeepisodes/imdb/{imdb}", v0.GetShowEpisodesByImdb())
	apiV0.HandleFunc("/tvmazeepisodes/tvdb/{tvdb}", v0.GetShowEpisodesByTvdb())
	apiV0.HandleFunc("/getmoviemagnet/imdb/{imdb}/query/{query}/providers/{providers}", v0.GetMovieTorrentsByImdbAndQuery())
	apiV0.HandleFunc("/getmoviemagnet/imdb/{imdb}/providers/{providers}", v0.GetMovieTorrentsByImdb())
	apiV0.HandleFunc("/getmoviemagnet/query/{query}/providers/{providers}", v0.GetMovieTorrentsByQuery())
	apiV0.HandleFunc("/getshowmagnet/imdb/{imdb}/query/{query}/season/{season}/episode/{episode}/providers/{providers}", v0.GetShowTorrentsByImdbAndQuery())
	apiV0.HandleFunc("/getshowmagnet/imdb/{imdb}/season/{season}/episode/{episode}/providers/{providers}", v0.GetShowTorrentsByImdb())
	apiV0.HandleFunc("/getshowmagnet/query/{query}/season/{season}/episode/{episode}/providers/{providers}", v0.GetShowTorrentsByQuery())
	apiV0.HandleFunc("/subtitlesbyimdb/{imdb}/lang/{lang}/season/{season}/episode/{episode}", v0.GetSubtitlesByImdb())
	apiV0.HandleFunc("/subtitlesbytext/{text}/lang/{lang}/season/{season}/episode/{episode}", v0.GetSubtitlesByText())
	apiV0.HandleFunc("/subtitlesbyfile/{hash}/{base64path}/lang/{lang}", v0.GetSubtitlesByFileHash())
	apiV0.HandleFunc("/add/{base64uri}", v0.AddTorrent())
	apiV0.HandleFunc("/torrents", v0.GetActiveTorrents())
	apiV0.HandleFunc("/stats/{hash}", v0.GetTorrentStats())
	apiV0.HandleFunc("/delete/{hash}", v0.DeleteTorrent())
	apiV0.HandleFunc("/deleteall", v0.DeleteAllTorrents())
	apiV0.HandleFunc("/mediarenderers", v0.GetMediaRenderers())
	apiV0.HandleFunc("/cast/{base64location}/{base64query}", v0.CastTorrentFile())
	apiV0.HandleFunc("/startplayer/{base64path}/{base64args}", v0.StartMediaPlayer())
	apiV0.HandleFunc("/restart/downrate/{downrate}/uprate/{uprate}", v0.RestartTorrentClient(quitSignal))
	apiV0.HandleFunc("/stop", v0.StopApplication(quitSignal))
	apiV0.HandleFunc("/about", v0.About(version))
	apiV0.NotFoundHandler = v0.NotFound()

	// Enable CORS for api urls if required
	if !*settings.CORS {
		return router
	} else {
		return handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}),
		)(router)
	}
}
