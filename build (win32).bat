set GOOS=windows
set GOARCH=386
go build -ldflags="-s -w" -mod=vendor -o wrserver.exe