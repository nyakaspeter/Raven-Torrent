# Raven Torrent

Raven Torrent is a torrent client application to find movies and TV shows from various sources and stream them over HTTP connection.

## Raven Server

Raven Server is a command line application. It hosts a REST-like HTTP API that can help you discover new movies and TV shows, search torrents for them on various torrent sites, and directly play them on your TV or PC. I recommend setting up a [Jackett](https://github.com/Jackett/Jackett) server, to get the most out of the application.

Raven Server is a fork of [White Raven Server](https://github.com/silentmurdock/wrserver). It was originally created as a backend for [White Raven](https://github.com/nyakaspeter/White-Raven), which is a torrent player application for Samsung Smart TV E, F, H series.

### API documentation

API docs for the server are available [here](docs/swagger.md). Swagger UI is also available on the `/swagger` endpoint, if the `-swagger` CLI argument is supplied when launching the server.

### CLI arguments

**-storagetype** `string` select storage type (must be set to "memory" or "file") (`default "memory"`)

**-memorysize** `int` specify the storage memory size in MB if storagetype is set to "memory" (minimum 64) (`default 128`)

**-dir** `string` specify the directory where files will be downloaded to if storagetype is set to "file"

**-downrate** `int` download speed rate in Kbps (0 is unlimited speed) (`default 0`)

**-uprate** `int` upload speed rate in Kbps (0 is upload disabled) (`default 0`)

**-maxconn** `int` max connections per torrent (`default 50`)

**-nodht** disable dht

**-jackettaddress** `string` set external Jackett API address (enter only host address and port e.g. `http://192.168.0.2:9117`)

**-jackettkey** `string` set external Jackett API key

**-tmdbkey** `string` set external TMDB API key

**-osuseragent** `string` set external OpenSubtitles user agent

**-port** `int` listening port (`default 9000`)

**-host** `string` listening server ip

**-cors** enable CORS

**-background** run the server in the background

**-help** print this help message

**-log** enable log messages

**-swagger** enable Swagger UI

### CLI examples

#### Running the executable file with parameters to serve torrent data from a 512MB memory buffer

`./raven -memorysize=512`

#### Running the executable file with parameters to serve torrent data from local disk

`./raven -storagetype="file" -dir="downloads"`

#### Running the executable file with parameters to use Jackett for torrent search

`./raven -jackettaddress="http://192.168.0.2:9117" -jackettkey="1n53rty0urj4ck3tt4p1k3yh3r3"`

### Build Instructions

You can build the application by running the following commands from the project directory. [Go](https://golang.org/) must be installed for these to work.

#### Building for Samsung Smart TV E, F, H ARM series

`set GOOS=linux`
`set GOARCH=arm`
`set GOARM=7`
`go build -ldflags="-s -w" -o raven`

#### Building for Windows (x64)

`set GOOS=windows`
`set GOARCH=amd64`
`set CGO_ENABLED=0`
`go build -ldflags="-s -w" -o raven.exe`

#### Building for Windows (x86)

`set GOOS=windows`
`set GOARCH=386`
`go build -ldflags="-s -w" -o raven.exe`

#### Building for Linux (x64)

`set GOOS=linux`
`set GOARCH=amd64`
`go build -ldflags="-s -w" -o raven`

#### Building for Linux (x86)

`set GOOS=linux`
`set GOARCH=386`
`go build -ldflags="-s -w" -o raven`

### License

[GNU GENERAL PUBLIC LICENSE Version 3](LICENSE)
