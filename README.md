# Raven Torrent

Raven Torrent is a torrent client application to find movies and TV shows from various sources and stream them over HTTP connection.

## Raven Server

Raven Server is a command line application. It hosts a REST-like HTTP API that can help you discover new movies and TV shows, search torrents for them on various torrent sites, and directly play them on your TV or PC. I recommend setting up a [Jackett](https://github.com/Jackett/Jackett) server, to get the most out of the application.

Raven Server is a fork of [White Raven Server](https://github.com/silentmurdock/wrserver). It was originally created as a backend for [White Raven](https://github.com/nyakaspeter/White-Raven), which is a torrent player application for Samsung Smart TV E, F, H series.

### API documentation

By default, the API is available at http://localhost:9000/api/v0. API docs for the server are available [here](docs/swagger.md). Swagger UI is also available on the `/swagger` endpoint, if the `-swagger` CLI argument is supplied at launch. 

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

### How to use

Raven Server can run on all kinds of devices, using very little resources by today's standards. There are [builds](https://github.com/nyakaspeter/Raven-Torrent/releases) available for both Windows and Linux, and even an ARM version that can run on old TVs or smartphones. On Android the executable can be launched from a terminal emulator like [Termux](https://termux.dev). All you have to do is [download](https://github.com/nyakaspeter/Raven-Torrent/releases) or [build](#build-instructions) the correct binary for your device's operating system and architecture and run it from the command line. Some examples:

#### Running with arguments to serve torrent data from RAM

`./raven -storagetype="memory" -memorysize="128"`

#### Running with arguments to serve torrent data from local disk

`./raven -storagetype="file" -dir="downloads"`

#### Running with arguments to use Jackett for torrent search

`./raven -jackettaddress="http://192.168.0.2:9117" -jackettkey="1n53rty0urj4ck3tt4p1k3yh3r3"`

### Jackett integration

Raven only supports a handful of torrent sites with built-in API integration and webscrapers, but it can also connect to [Jackett](https://github.com/Jackett/Jackett) to enable torrent search on dozens of other sites. To be able to use this, you'll also have to run Jackett somewhere on your local network, and supply the `-jackettaddress` and `-jackettkey` arguments when launching the server.

Fortunately Jackett can run on most devices where Raven Server can run (except on Smart TVs where the resources are very limited). Check [their documentation](https://github.com/Jackett/Jackett#installation-on-windows) for info on how to run it on different systems. I also managed to run Jackett on an Android phone just like Raven Server with these steps:

1. Download & install Termux from [F-Droid](https://f-droid.org/packages/com.termux/) (the Google Play version is outdated), then open it and use the following commands
2. `pkg up` to update all Termux packages, use `termux-change-repo` to change download location if it throws an error
3. `pkg install proot-distro`
4. `proot-distro install ubuntu`
5. `proot-distro login ubuntu`
6. `apt-get update && apt-get upgrade -y` to update Ubuntu packages
7. `apt-get install libicu-dev`
8. `apt-get install wget` (or use any other means to get the Jackett binary into the Ubuntu filesystem)
9. `wget https://github.com/Jackett/Jackett/releases/latest/download/Jackett.Binaries.LinuxARM64.tar.gz` to download the latest ARM64 build of Jackett (ARM32 may work too)
10. `tar -xvzf Jackett.Binaries.LinuxARM64.tar.gz`
11. `./Jackett/jackett`

### Build instructions

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
