set GOOS=linux
set GOARCH=arm
set GOARM=7
go build -ldflags="-s -w -checklinkname=0" -o ..\build\api\linux-armv7\raven ..\api