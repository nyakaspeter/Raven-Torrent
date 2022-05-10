set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=0
go build -ldflags="-s -w" -o ..\bin\windows-x64\raven.exe ..\cmd\raven-api