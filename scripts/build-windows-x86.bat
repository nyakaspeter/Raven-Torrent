set GOOS=windows
set GOARCH=386
go build -ldflags="-s -w" -o ..\build\api\windows-x86\raven.exe ..\api