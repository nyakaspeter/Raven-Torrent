set GOOS=windows
set GOARCH=386
go build -ldflags="-s -w" -o ..\bin\windows-x86\raven.exe ..\cmd\raven-api