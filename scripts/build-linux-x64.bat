set GOOS=linux
set GOARCH=amd64
go build -ldflags="-s -w" -o ..\build\api\linux-x64\raven ..\api