set GOOS=linux
set GOARCH=amd64
go build -ldflags="-s -w" -o ..\bin\linux-x64\raven ..\cmd\raven-api