set GOOS=linux
set GOARCH=amd64
go build -ldflags="-s -w -checklinkname=0" -o ..\build\api\linux-x64\raven ..\api