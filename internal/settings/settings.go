package settings

import "flag"

var Host *string
var Port *int
var DlnaPort *int
var DownloadDir *string
var DownloadRate *int
var UploadRate *int
var MaxConnections *int
var NoDHT *bool
var EnableLog *bool
var EnableReceiver *bool
var EnableSwaggerUi *bool
var StorageType *string
var MemorySize *int64
var Background *bool
var CORS *bool
var TMDBKey *string
var OpenSubtitlesUser *string
var OpenSubtitlesPassword *string
var OpenSubtitlesKey *string
var JackettAddress *string
var JackettKey *string
var NcoreUser *string
var NcorePassword *string

func Init() {
	Host = flag.String("host", "", "listening server ip")
	Port = flag.Int("port", 9000, "listening port")
	DlnaPort = flag.Int("dlnaport", 3500, "DLNA server port")
	DownloadDir = flag.String("dir", "data", "specify the directory where files will be downloaded to if storagetype is set to \"file\"")
	DownloadRate = flag.Int("downrate", 0, "download speed rate in Kbps")
	UploadRate = flag.Int("uprate", 0, "upload speed rate in Kbps")
	MaxConnections = flag.Int("maxconn", 50, "max connections per torrent")
	NoDHT = flag.Bool("nodht", false, "disable dht")
	EnableLog = flag.Bool("log", false, "enable log messages")
	EnableReceiver = flag.Bool("receiver", true, "enable torrent receiver page")
	EnableSwaggerUi = flag.Bool("swagger", false, "enable swagger ui")
	StorageType = flag.String("storagetype", "memory", "select storage type (must be set to \"memory\" or \"file\")")
	Background = flag.Bool("background", false, "run the server in the background")
	CORS = flag.Bool("cors", true, "enable CORS")
	MemorySize = flag.Int64("memorysize", 128, "specify the storage memory size in MB if storagetype is set to \"memory\" (minimum 64)") // 64MB is optimal for TVs
	TMDBKey = flag.String("tmdbkey", "a4d9ad8d2d072c50dc998cc0d1a508fa", "set external TMDB API key")
	JackettAddress = flag.String("jackettaddress", "", "set external Jackett API address")
	JackettKey = flag.String("jackettkey", "", "set external Jackett API key")
	OpenSubtitlesUser = flag.String("osuser", "", "set OpenSubtitles.com username")
	OpenSubtitlesPassword = flag.String("ospassword", "", "set external OpenSubtitles.com password")
	OpenSubtitlesKey = flag.String("osapikey", "", "set OpenSubtitles.com api key")
	NcoreUser = flag.String("ncoreuser", "", "set nCore username")
	NcorePassword = flag.String("ncorepassword", "", "set nCore password")
	flag.Parse()
}
