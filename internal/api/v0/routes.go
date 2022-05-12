package v0

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	api "github.com/nyakaspeter/raven-torrent/internal/api/v0/handlers"
)

const version = "0.6.0"
const apiPrefix = "/api/"

func handleAPI(cors bool) http.Handler {
	router := mux.NewRouter()
	router.SkipClean(true)
	router.NotFoundHandler = api.NotFound()

	router.HandleFunc(apiPrefix+"about", api.About(version))
	router.HandleFunc(apiPrefix+"stop", api.StopApplication(procQuit))
	router.HandleFunc(apiPrefix+"restart/downrate/{downrate}/uprate/{uprate}", api.RestartTorrentClient(procRestart))
	router.HandleFunc(apiPrefix+"add/{base64uri}", api.AddTorrent())
	router.HandleFunc(apiPrefix+"getmediarenderers", api.GetMediaRenderers())
	router.HandleFunc(apiPrefix+"startplayer/{base64path}/{base64args}", api.StartMediaPlayer())
	router.HandleFunc(apiPrefix+"cast/{base64location}/{base64query}", api.CastTorrentFile())
	router.HandleFunc(apiPrefix+"delete/{hash}", api.DeleteTorrent())
	router.HandleFunc(apiPrefix+"deleteall", api.DeleteAllTorrents())
	router.HandleFunc(apiPrefix+"get/{hash}/{base64path}", api.GetTorrentFile())
	router.HandleFunc(apiPrefix+"stats/{hash}", api.GetTorrentStats())
	router.HandleFunc(apiPrefix+"subtitlesbyimdb/{imdb}/lang/{lang}/season/{season}/episode/{episode}", api.GetSubtitlesByImdb(OpenSubtitlesUserAgent))
	router.HandleFunc(apiPrefix+"subtitlesbytext/{text}/lang/{lang}/season/{season}/episode/{episode}", api.GetSubtitlesByText(OpenSubtitlesUserAgent))
	router.HandleFunc(apiPrefix+"subtitlesbyfile/{hash}/{base64path}/lang/{lang}", api.GetSubtitlesByFileHash(OpenSubtitlesUserAgent))
	router.HandleFunc(apiPrefix+"getsubtitle/{base64path}/encode/{encode}/subtitle.{subtype}", api.GetSubtitleFile(OpenSubtitlesUserAgent))
	router.HandleFunc(apiPrefix+"torrents", api.GetActiveTorrents())
	router.HandleFunc(apiPrefix+"getmoviemagnet/imdb/{imdb}/providers/{providers}", api.GetMovieTorrentsByImdb())
	router.HandleFunc(apiPrefix+"getmoviemagnet/query/{query}/providers/{providers}", api.GetMovieTorrentsByQuery())
	router.HandleFunc(apiPrefix+"getmoviemagnet/imdb/{imdb}/query/{query}/providers/{providers}", api.GetMovieTorrentsByImdbAndQuery())
	router.HandleFunc(apiPrefix+"getshowmagnet/imdb/{imdb}/season/{season}/episode/{episode}/providers/{providers}", api.GetShowTorrentsByImdb())
	router.HandleFunc(apiPrefix+"getshowmagnet/query/{query}/season/{season}/episode/{episode}/providers/{providers}", api.GetShowTorrentsByQuery())
	router.HandleFunc(apiPrefix+"getshowmagnet/imdb/{imdb}/query/{query}/season/{season}/episode/{episode}/providers/{providers}", api.GetShowTorrentsByImdbAndQuery())
	router.HandleFunc(apiPrefix+"tmdbdiscover/type/{type}/genretype/{genretype}/sort/{sort}/date/{date}/lang/{lang}/page/{page}", api.TmdbDiscover())
	router.HandleFunc(apiPrefix+"tmdbsearch/type/{type}/lang/{lang}/page/{page}/text/{text}", api.TmdbSearch())
	router.HandleFunc(apiPrefix+"tmdbinfo/type/{type}/tmdbid/{tmdbid}/lang/{lang}", api.TmdbInfo())
	router.HandleFunc(apiPrefix+"tvmazeepisodes/imdb/{imdb}", api.GetShowEpisodesByImdb())
	router.HandleFunc(apiPrefix+"tvmazeepisodes/tvdb/{tvdb}", api.GetShowEpisodesByTvdb())
	router.HandleFunc(apiPrefix+"tvmazeepisodes/tvdb/{tvdb}/imdb/{imdb}", api.GetShowEpisodesByImdbAndTvdb())
	router.HandleFunc(apiPrefix+"receivemagnet/{todo}", api.ReceiveTorrent())
	router.HandleFunc(apiPrefix+"websocket", api.Websocket(procQuit))

	// Create torrent magnet send page from main page
	router.HandleFunc("/", api.ReceiverPage(version, apiPrefix))

	// Enable CORS for api urls if required
	if !cors {
		return router
	} else {
		return handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)
	}
}
