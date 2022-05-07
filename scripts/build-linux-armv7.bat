set GOOS=linux
set GOARCH=arm
set GOARM=7
go build -ldflags="-s -w" -o ..\bin\linux-armv7\wrserver ..\cmd\wrserver