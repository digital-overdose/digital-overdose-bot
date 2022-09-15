@echo off
set GOOS=darwin
set GOARCH=amd64
go build -o build/digital-overdose-bot-darwin-amd64 .

set GOOS=darwin
set GOARCH=arm64
go build -o build/digital-overdose-bot-darwin-arm64 .

set GOOS=dragonfly
set GOARCH=amd64
go build -o build/digital-overdose-bot-dragonfly-amd64 .

set GOOS=freebsd
set GOARCH=386
go build -o build/digital-overdose-bot-freebsd-386 .

set GOOS=freebsd
set GOARCH=amd64
go build -o build/digital-overdose-bot-freebsd-amd64 .

set GOOS=freebsd
set GOARCH=arm
go build -o build/digital-overdose-bot-freebsd-arm .

set GOOS=linux
set GOARCH=386
go build -o build/digital-overdose-bot-linux-386 .

set GOOS=linux
set GOARCH=amd64
go build -o build/digital-overdose-bot-linux-amd64 .

set GOOS=linux
set GOARCH=arm
go build -o build/digital-overdose-bot-linux-arm .

set GOOS=linux
set GOARCH=arm64
go build -o build/digital-overdose-bot-linux-arm64 .

set GOOS=linux
set GOARCH=ppc64
go build -o build/digital-overdose-bot-linux-ppc64 .

set GOOS=linux
set GOARCH=ppc64le
go build -o build/digital-overdose-bot-linux-ppc64le .

set GOOS=linux
set GOARCH=mips
go build -o build/digital-overdose-bot-linux-mips .

set GOOS=linux
set GOARCH=mipsle
go build -o build/digital-overdose-bot-linux-mipsle .

set GOOS=linux
set GOARCH=mips64
go build -o build/digital-overdose-bot-linux-mips64 .

set GOOS=linux
set GOARCH=mips64le
go build -o build/digital-overdose-bot-linux-mips64le .

set GOOS=netbsd
set GOARCH=386
go build -o build/digital-overdose-bot-netbsd-386 .

set GOOS=netbsd
set GOARCH=amd64
go build -o build/digital-overdose-bot-netbsd-amd64 .

set GOOS=netbsd
set GOARCH=arm
go build -o build/digital-overdose-bot-netbsd-arm .

set GOOS=openbsd
set GOARCH=386
go build -o build/digital-overdose-bot-openbsd-386 .

set GOOS=openbsd
set GOARCH=amd64
go build -o build/digital-overdose-bot-openbsd-amd64 .

set GOOS=openbsd
set GOARCH=arm
go build -o build/digital-overdose-bot-openbsd-arm .

set GOOS=plan9
set GOARCH=386
go build -o build/digital-overdose-bot-plan9-386 .

set GOOS=plan9
set GOARCH=amd64
go build -o build/digital-overdose-bot-plan9-amd64 .

set GOOS=solaris
set GOARCH=amd64
go build -o build/digital-overdose-bot-solaris-amd64 .

set GOOS=windows
set GOARCH=386
go build -o build/digital-overdose-bot-windows-386.exe .

set GOOS=windows
set GOARCH=amd64
go build -o build/digital-overdose-bot-windows-amd64.exe .