set GOOS=linux
set GOARCH=386
go build -ldflags="-s -w" -o wrserver