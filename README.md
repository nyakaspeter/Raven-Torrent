# White Raven Server

**White Raven Server is a REST-like API controlled torrent client application to find movies and tv shows from various sources and stream them over http connection. Mainly created for [White Raven](https://github.com/silentmurdock/whiteraven), which is a torrent player application for Samsung Smart TV E, F, H series.**

<br/>

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

- [Get movie magnet by IMDB id](documents/api/moviebyimdb.md)
- [Get movie magnet by query text](documents/api/moviebytext.md)
- [Get movie magnet by IMDB id and query text at once](documents/api/moviebyboth.md)
- [Get show magnet by IMDB id](documents/api/showbyimdb.md)
- [Get show magnet by query text](documents/api/showbytext.md)
- [Get show magnet by IMDB id and query text at once](documents/api/showbyboth.md)
- [Discover movies or tv shows](documents/api/tmdbdiscover.md)
- [Search movies or tv shows by query text](documents/api/tmdbsearch.md)
- [Get more info about movie or tv show by TMDB id](documents/api/tmdbinfo.md)

### Subtitle related

- [Search subtitles by IMDB id](documents/api/subtitlesbyimdb.md)
- [Search subtitles by query text](documents/api/subtitlesbytext.md)
- [Search subtitles by inner file hash](documents/api/subtitlesbyhash.md)
- [Download subtitle file](documents/api/getsubtitle.md)
  <br/>

## Command-Line Arguments

- **-background** run the server in the background
- **-cors** enable CORS
- **-dir** `string` specify the directory where files will be downloaded to if storagetype is set to "piecefile" or "file"
- **-downrate** `int` download speed rate in Kbps (`default 32768`)
- **-help** print this help message
- **-host** `string` listening server ip
- **-log** enable log messages
- **-maxconn** `int` max connections per torrent (`default 40`)
- **-memorysize** `int` specify the storage memory size in MB if storagetype is set to "memory" (minimum 64) (`default 128`)
- **-nodht** disable dht
- **-osuseragent**`string` set external OpenSubtitles user agent
- **-port** `int` listening port (`default 9000`)
- **-storagetype**`string` select storage type (must be set to "memory" or "piecefile" or "file") (`default "memory"`)
- **-tmdbkey**`string` set external TMDB API key
- **-uprate** `int` upload speed rate in Kbps (`default 256`)
  <br/>

## Build Instructions

### Build On Windows

**Download:**

```
$ go get -v -u github.com/silentmurdock/wrserver
```

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

<br/>

## Run The Server

**Simply run the executable file without parameters to serve torrent data from memory.**

```
$ wrserver
```

**Run the executable file with the following parameters to serve torrent data from local disk.**

```
$ wrserver -storagetype="file" -dir="downloads"
```

<br/>

## Note For Releases

The releases always compressed with the latest version of [UPX](https://upx.github.io), an advanced executable file packer to decrease the size of the application. This is important for embedded devices such as Samsung Smart TVs because they have a very limited amount of resources!
<br/>

## License

[GNU GENERAL PUBLIC LICENSE Version 3](LICENSE)
