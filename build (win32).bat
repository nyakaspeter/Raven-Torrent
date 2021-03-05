set GOOS=windows
set GOARCH=386
go build -ldflags="-s -w" -o wrserver.exe 