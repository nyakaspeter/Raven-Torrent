# White Raven Server

**White Raven Server is a REST-like API controlled torrent client application to find movies and tv shows from various sources and stream them over http connection. Mainly created for [White Raven](https://github.com/silentmurdock/whiteraven), which is a torrent player application for Samsung Smart TV E, F, H series.**

## HTTP API Functions

### Server and Client related

- [Get server information](documents/api/about.md)
- [Stop server](documents/api/stop.md)
- [Restart torrent client](documents/api/restart.md)

### Torrent related

- [Add torrent by hash](documents/api/add.md)
- [Delete torrent by hash](documents/api/delete.md)
- [Delete all running torrents](documents/api/deleteall.md)
- [Get all running torrents](documents/api/torrents.md)
- [Get running torrent statistics by hash](documents/api/stats.md)
- [Stream or download the selected file](documents/api/get.md)

### Movie or TV Show related

- [Get movie torrents by IMDB id or query text or both](documents/api/getmoviemagnet.md)
- [Get tv show torrents by IMDB id or query text or both](documents/api/getshowmagnet.md)
- [Discover movies or tv shows](documents/api/tmdbdiscover.md)
- [Search movies or tv shows by query text](documents/api/tmdbsearch.md)
- [Get more info about movie or tv show by TMDB id](documents/api/tmdbinfo.md)
- [Get tv show episodes by IMDB id or TVDB id](documents/api/tvmazeepisodes.md)

### Subtitle related

- [Search subtitles by IMDB id](documents/api/subtitlesbyimdb.md)
- [Search subtitles by query text](documents/api/subtitlesbytext.md)
- [Search subtitles by inner file hash](documents/api/subtitlesbyhash.md)
- [Download subtitle file](documents/api/getsubtitle.md)

## Command-Line Arguments

- **-storagetype** `string` select storage type (must be set to "memory" or "piecefile" or "file") (`default "memory"`)
- **-memorysize** `int` specify the storage memory size in MB if storagetype is set to "memory" (minimum 64) (`default 128`)
- **-dir** `string` specify the directory where files will be downloaded to if storagetype is set to "piecefile" or "file"

- **-downrate** `int` download speed rate in Kbps (0 is unlimited speed) (`default 0`)
- **-uprate** `int` upload speed rate in Kbps (0 is upload disabled) (`default 0`)
- **-maxconn** `int` max connections per torrent (`default 50`)
- **-nodht** disable dht

- **-jackettaddress** `string` set external Jackett API address (enter only host address and port e.g. `http://192.168.0.2:9117`)
- **-jackettkey** `string` set external Jackett API key
- **-tmdbkey** `string` set external TMDB API key
- **-osuseragent** `string` set external OpenSubtitles user agent

- **-port** `int` listening port (`default 9000`)
- **-host** `string` listening server ip
- **-cors** enable CORS

- **-background** run the server in the background
- **-help** print this help message
- **-log** enable log messages

## Build Instructions

### Build On Windows

**Build in vendor mode for Samsung Smart TV E, F, H ARM series:**

```
$ set GOOS=linux
$ set GOARCH=arm
$ set GOARM=7
$ go build -ldflags="-s -w" -mod=vendor -o wrserver
```

**Build in vendor mode for Windows x32:**

```
$ set GOOS=windows
$ set GOARCH=386
$ go build -ldflags="-s -w" -mod=vendor -o wrserver.exe
```

**Build in vendor mode for Windows x64:**

```
$ set GOOS=windows
$ set GOARCH=amd64
$ set CGO_ENABLED=0
$ go build -ldflags="-s -w" -mod=vendor -o wrserver.exe
```

**Build in vendor mode for Linux x32:**

```
$ set GOOS=linux
$ set GOARCH=386
$ go build -ldflags="-s -w" -mod=vendor -o wrserver
```

**Build in vendor mode for Linux x64:**

```
$ set GOOS=linux
$ set GOARCH=amd64
$ go build -ldflags="-s -w" -mod=vendor -o wrserver
```

## Run The Server

**Simply run the executable file without parameters to serve torrent data from memory.**

```
$ wrserver
```

**Run the executable file with the following parameters to serve torrent data from local disk.**

```
$ wrserver -storagetype="file" -dir="downloads"
```

## Note For Releases

The releases always compressed with the latest version of [UPX](https://upx.github.io), an advanced executable file packer to decrease the size of the application. This is important for embedded devices such as Samsung Smart TVs because they have a very limited amount of resources!

## License

[GNU GENERAL PUBLIC LICENSE Version 3](LICENSE)
