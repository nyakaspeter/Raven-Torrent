set GOOS=windows
set GOARCH=386
go build -ldflags="-s -w" -o ..\bin\windows-x86\wrserver.exe ..\cmd\wrserver