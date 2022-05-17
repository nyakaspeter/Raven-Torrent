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
var StorageType *string
var MemorySize *int64
var Background *bool
var CORS *bool
var TMDBKey *string
var OpenSubtitlesUserAgent *string
var JackettAddress *string
var JackettKey *string

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
	StorageType = flag.String("storagetype", "memory", "select storage type (must be set to \"memory\" or \"file\")")
	Background = flag.Bool("background", false, "run the server in the background")
	CORS = flag.Bool("cors", true, "enable CORS")
	MemorySize = flag.Int64("memorysize", 128, "specify the storage memory size in MB if storagetype is set to \"memory\" (minimum 64)") // 64MB is optimal for TVs
	TMDBKey = flag.String("tmdbkey", "a4d9ad8d2d072c50dc998cc0d1a508fa", "set external TMDB API key")
	JackettAddress = flag.String("jackettaddress", "", "set external Jackett API address")
	JackettKey = flag.String("jackettkey", "", "set external Jackett API key")
	OpenSubtitlesUserAgent = flag.String("osuseragent", "White Raven v0.3", "set external OpenSubtitles user agent")
	flag.Parse()
}
