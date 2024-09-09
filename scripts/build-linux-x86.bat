set GOOS=linux
set GOARCH=386
go build -ldflags="-s -w -checklinkname=0" -o ..\build\api\linux-x86\raven ..\api