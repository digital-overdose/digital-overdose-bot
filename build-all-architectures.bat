@echo off
set VERSION=1.1.3-hotfix

:: set GOOS=darwin     && set GOARCH=amd64     && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
:: set GOOS=darwin     && set GOARCH=arm64     && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .

:: set GOOS=dragonfly  && set GOARCH=amd64     && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .

:: set GOOS=freebsd    && set GOARCH=386       && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
:: set GOOS=freebsd    && set GOARCH=amd64     && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
:: set GOOS=freebsd    && set GOARCH=arm       && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .

:: set GOOS=linux      && set GOARCH=386       && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
set GOOS=linux      && set GOARCH=amd64     && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
set GOOS=linux      && set GOARCH=arm       && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
set GOOS=linux      && set GOARCH=arm64     && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
:: set GOOS=linux      && set GOARCH=ppc64     && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
:: set GOOS=linux      && set GOARCH=ppc64le   && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
:: set GOOS=linux      && set GOARCH=mips      && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
:: set GOOS=linux      && set GOARCH=mipsle    && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
:: set GOOS=linux      && set GOARCH=mips64    && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
:: set GOOS=linux      && set GOARCH=mips64le  && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .

:: set GOOS=netbsd     && set GOARCH=386       && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
:: set GOOS=netbsd     && set GOARCH=amd64     && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
:: set GOOS=netbsd     && set GOARCH=arm       && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .

:: set GOOS=openbsd    && set GOARCH=386       && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
:: set GOOS=openbsd    && set GOARCH=amd64     && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
:: set GOOS=openbsd    && set GOARCH=arm       && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .

:: set GOOS=plan9      && set GOARCH=386       && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .
:: set GOOS=plan9      && set GOARCH=amd64     && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .

:: set GOOS=solaris    && set GOARCH=amd64     && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH% .

set GOOS=windows    && set GOARCH=386       && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH%.exe .
set GOOS=windows    && set GOARCH=amd64     && go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH%.exe .