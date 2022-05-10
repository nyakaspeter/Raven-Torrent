# White Raven Server

**This repo is a fork of [White Raven Server](https://github.com/silentmurdock/wrserver) by Murdock. White Raven Server is a REST-like API controlled torrent client application to find movies and tv shows from various sources and stream them over http connection. Mainly created for [White Raven](https://github.com/nyakaspeter/White-Raven), which is a torrent player application for Samsung Smart TV E, F, H series.**

## HTTP API Functions

### Server and Client related

- [Get server information](docs/api/about.md)
- [Stop server](docs/api/stop.md)
- [Restart torrent client](docs/api/restart.md)

### Torrent related

- [Add torrent by magnet link](docs/api/add.md)
- [Delete torrent by hash](docs/api/delete.md)
- [Delete all running torrents](docs/api/deleteall.md)
- [Get all running torrents](docs/api/torrents.md)
- [Get running torrent statistics by hash](docs/api/stats.md)
- [Stream or download the selected file](docs/api/get.md)

### Movie or TV Show related

- [Discover movies or tv shows](docs/api/tmdbdiscover.md)
- [Search movies or tv shows by query text](docs/api/tmdbsearch.md)
- [Get more info about movie or tv show by TMDB id](docs/api/tmdbinfo.md)
- [Get tv show episodes by IMDB id or TVDB id or both](docs/api/tvmazeepisodes.md)
- [Get movie torrents by IMDB id or query text or both](docs/api/getmoviemagnet.md)
- [Get tv show torrents by IMDB id or query text or both](docs/api/getshowmagnet.md)

### Subtitle related

- [Search subtitles by IMDB id](docs/api/subtitlesbyimdb.md)
- [Search subtitles by query text](docs/api/subtitlesbytext.md)
- [Search subtitles by inner file hash](docs/api/subtitlesbyhash.md)
- [Download subtitle file](docs/api/getsubtitle.md)

## Command-Line Arguments

- **-storagetype** `string` select storage type (must be set to "memory" or "file") (`default "memory"`)
- **-memorysize** `int` specify the storage memory size in MB if storagetype is set to "memory" (minimum 64) (`default 128`)
- **-dir** `string` specify the directory where files will be downloaded to if storagetype is set to "file"

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

## How to use

White Raven Server is a command line application. Some examples for usage:

<details>
<summary>Running the executable file without parameters to serve torrent data from memory</summary>

```
raven
```

</details>

<details>
<summary>Running the executable file with parameters to serve torrent data from local disk</summary>

```
raven -storagetype="file" -dir="downloads"
```

</details>

<details>
<summary>Running the executable file with parameters to use Jackett for torrent search</summary>

```
raven -jackettaddress="http://192.168.0.2:9117" -jackettkey="1n53rty0urj4ck3tt4p1k3yh3r3"
```

</details>

## Build Instructions

You can build the application by running the following commands from the project directory. [Go](https://golang.org/) must be installed for these to work.

<details>
<summary>Build for Samsung Smart TV E, F, H ARM series</summary>

```
set GOOS=linux
set GOARCH=arm
set GOARM=7
go build -ldflags="-s -w" -o raven
```

</details>

<details>
<summary>Build for Windows (x64)</summary>

```
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=0
go build -ldflags="-s -w" -o raven.exe
```

</details>

<details>
<summary>Build for Windows (x86)</summary>

```
set GOOS=windows
set GOARCH=386
go build -ldflags="-s -w" -o raven.exe
```

</details>

<details>
<summary>Build for Linux (x64)</summary>

```
set GOOS=linux
set GOARCH=amd64
go build -ldflags="-s -w" -o raven
```

</details>

<details>
<summary>Build for Linux (x86)</summary>

```
set GOOS=linux
set GOARCH=386
go build -ldflags="-s -w" -o raven
```

</details>

## Donation
Support the original creator using these addresses:
```
BTC: 3NHNGcGEyD3m88nWZpNCSjnGq95rgFq4P5
LTC: MNyTkncUNkLgv8TCuPn69NVvBXnEsrtWcW
```
```
PATREON: https://www.patreon.com/murdock
```

## License

[GNU GENERAL PUBLIC LICENSE Version 3](LICENSE)
