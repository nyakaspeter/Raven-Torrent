set GOOS=linux
set GOARCH=386
go build -ldflags="-s -w" -o ..\build\api\linux-x86\raven ..\api