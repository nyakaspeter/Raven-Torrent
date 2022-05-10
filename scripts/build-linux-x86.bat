set GOOS=linux
set GOARCH=386
go build -ldflags="-s -w" -o ..\bin\linux-x86\raven ..\cmd\raven-api