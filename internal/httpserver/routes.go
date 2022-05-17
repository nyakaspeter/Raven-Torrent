package httpserver

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	v0 "github.com/nyakaspeter/raven-torrent/internal/httpserver/api/v0"
)

const version = "0.6.0"

func handleAPI(cors bool, receiver bool) http.Handler {
	router := mux.NewRouter()
	router.SkipClean(true)

	if receiver {
		// Create torrent magnet send page from main page
		router.HandleFunc("/", ReceiverPage())
	}

	var api = router.PathPrefix("/api")
	var apiV0 = api.PathPrefix("/v0").Subrouter()
	//var apiV1 = api.PathPrefix("/v1").Subrouter()

	apiV0.HandleFunc("/about", v0.About(version))
	apiV0.HandleFunc("/tmdbdiscover/type/{type}/genretype/{genretype}/sort/{sort}/date/{date}/lang/{lang}/page/{page}", v0.TmdbDiscover())
	apiV0.HandleFunc("/tmdbsearch/type/{type}/lang/{lang}/page/{page}/text/{text}", v0.TmdbSearch())
	apiV0.HandleFunc("/tmdbinfo/type/{type}/tmdbid/{tmdbid}/lang/{lang}", v0.TmdbInfo())
	apiV0.HandleFunc("/tvmazeepisodes/tvdb/{tvdb}/imdb/{imdb}", v0.GetShowEpisodesByImdbAndTvdb())
	apiV0.HandleFunc("/tvmazeepisodes/imdb/{imdb}", v0.GetShowEpisodesByImdb())
	apiV0.HandleFunc("/tvmazeepisodes/tvdb/{tvdb}", v0.GetShowEpisodesByTvdb())
	apiV0.HandleFunc("/getmoviemagnet/imdb/{imdb}/query/{query}/providers/{providers}", v0.GetMovieTorrentsByImdbAndQuery())
	apiV0.HandleFunc("/getmoviemagnet/imdb/{imdb}/providers/{providers}", v0.GetMovieTorrentsByImdb())
	apiV0.HandleFunc("/getmoviemagnet/query/{query}/providers/{providers}", v0.GetMovieTorrentsByQuery())
	apiV0.HandleFunc("/getshowmagnet/imdb/{imdb}/query/{query}/season/{season}/episode/{episode}/providers/{providers}", v0.GetShowTorrentsByImdbAndQuery())
	apiV0.HandleFunc("/getshowmagnet/imdb/{imdb}/season/{season}/episode/{episode}/providers/{providers}", v0.GetShowTorrentsByImdb())
	apiV0.HandleFunc("/getshowmagnet/query/{query}/season/{season}/episode/{episode}/providers/{providers}", v0.GetShowTorrentsByQuery())
	apiV0.HandleFunc("/get/{hash}/{base64path}", v0.GetTorrentFile())
	apiV0.HandleFunc("/add/{base64uri}", v0.AddTorrent())
	apiV0.HandleFunc("/receivemagnet/{todo}", v0.ReceiveTorrent())
	apiV0.HandleFunc("/torrents", v0.GetActiveTorrents())
	apiV0.HandleFunc("/stats/{hash}", v0.GetTorrentStats())
	apiV0.HandleFunc("/delete/{hash}", v0.DeleteTorrent())
	apiV0.HandleFunc("/deleteall", v0.DeleteAllTorrents())
	apiV0.HandleFunc("/subtitlesbyimdb/{imdb}/lang/{lang}/season/{season}/episode/{episode}", v0.GetSubtitlesByImdb())
	apiV0.HandleFunc("/subtitlesbytext/{text}/lang/{lang}/season/{season}/episode/{episode}", v0.GetSubtitlesByText())
	apiV0.HandleFunc("/subtitlesbyfile/{hash}/{base64path}/lang/{lang}", v0.GetSubtitlesByFileHash())
	apiV0.HandleFunc("/getsubtitle/{base64path}/encode/{encode}/subtitle.{subtype}", v0.GetSubtitleFile())
	apiV0.HandleFunc("/getmediarenderers", v0.GetMediaRenderers())
	apiV0.HandleFunc("/cast/{base64location}/{base64query}", v0.CastTorrentFile())
	apiV0.HandleFunc("/startplayer/{base64path}/{base64args}", v0.StartMediaPlayer())
	apiV0.HandleFunc("/restart/downrate/{downrate}/uprate/{uprate}", v0.RestartTorrentClient(quitSignal))
	apiV0.HandleFunc("/stop", v0.StopApplication(quitSignal))
	apiV0.HandleFunc("/websocket", v0.Websocket(quitSignal))
	apiV0.NotFoundHandler = v0.NotFound()

	// Enable CORS for api urls if required
	if !cors {
		return router
	} else {
		return handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}),
		)(router)
	}
}
