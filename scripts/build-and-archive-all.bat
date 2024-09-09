@echo off

set VERSION=0.7.0
set PLATFORMS=(windows-x64,windows-x86, linux-x64, linux-x86, linux-armv7)

for %%P in %PLATFORMS% do (
    set PLATFORM=%%P

    call build-%%P
    echo Binary built: %%P

    call archive-build
    echo Binary archived: %%P
)