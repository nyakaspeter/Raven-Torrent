set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=0
go build -ldflags="-s -w" -o ..\build\api\windows-x64\raven.exe ..\api