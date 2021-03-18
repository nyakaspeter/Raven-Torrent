set VERSION=0.5.1

call build-linux-armv7 & tar -a -c -f wrserver-%VERSION%-linux-armv7.zip wrserver & del wrserver
call build-linux-x64 & tar -a -c -f wrserver-%VERSION%-linux-x64.zip wrserver & del wrserver
call build-linux-x86 & tar -a -c -f wrserver-%VERSION%-linux-x86.zip wrserver & del wrserver
call build-windows-x64 & tar -a -c -f wrserver-%VERSION%-windows-x64.zip wrserver.exe & del wrserver.exe
call build-windows-x86 & tar -a -c -f wrserver-%VERSION%-windows-x86.zip wrserver.exe & del wrserver.exe