#!/bin/sh

GOOS=darwin			GOARCH=amd64 		go build -o build/digital-overdose-bot-darwin-amd64 .
GOOS=darwin			GOARCH=arm64 		go build -o build/digital-overdose-bot-darwin-arm64 .

GOOS=dragonfly 	GOARCH=amd64 		go build -o build/digital-overdose-bot-dragonfly-amd64 .

GOOS=freebsd 		GOARCH=386 	 		go build -o build/digital-overdose-bot-freebsd-386 .
GOOS=freebsd 		GOARCH=amd64 		go build -o build/digital-overdose-bot-freebsd-amd64 .
GOOS=freebsd 		GOARCH=arm 	 		go build -o build/digital-overdose-bot-freebsd-arm .

GOOS=linux	 		GOARCH=386 			go build -o build/digital-overdose-bot-linux-386 .
GOOS=linux	 		GOARCH=amd64 		go build -o build/digital-overdose-bot-linux-amd64 .
GOOS=linux	 		GOARCH=arm	 		go build -o build/digital-overdose-bot-linux-arm .
GOOS=linux	 		GOARCH=arm64 		go build -o build/digital-overdose-bot-linux-arm64 .
GOOS=linux	 		GOARCH=ppc64 		go build -o build/digital-overdose-bot-linux-ppc64 .
GOOS=linux	 		GOARCH=ppc64le 	go build -o build/digital-overdose-bot-linux-ppc64le .
GOOS=linux	 		GOARCH=mips 		go build -o build/digital-overdose-bot-linux-mips .
GOOS=linux	 		GOARCH=mipsle 	go build -o build/digital-overdose-bot-linux-mipsle .
GOOS=linux	 		GOARCH=mips64 	go build -o build/digital-overdose-bot-linux-mips64 .
GOOS=linux	 		GOARCH=mips64le go build -o build/digital-overdose-bot-linux-mips64le .

GOOS=netbsd			GOARCH=386 			go build -o build/digital-overdose-bot-netbsd-386 .
GOOS=netbsd			GOARCH=amd64 		go build -o build/digital-overdose-bot-netbsd-amd64 .
GOOS=netbsd			GOARCH=arm 			go build -o build/digital-overdose-bot-netbsd-arm .

GOOS=openbsd		GOARCH=386 			go build -o build/digital-overdose-bot-openbsd-386 .
GOOS=openbsd		GOARCH=amd64 		go build -o build/digital-overdose-bot-openbsd-amd64 .
GOOS=openbsd		GOARCH=arm 			go build -o build/digital-overdose-bot-openbsd-arm .

GOOS=plan9			GOARCH=386 			go build -o build/digital-overdose-bot-plan9-386 .
GOOS=plan9			GOARCH=amd64 		go build -o build/digital-overdose-bot-plan9-amd64 .

GOOS=solaris		GOARCH=amd64 		go build -o build/digital-overdose-bot-solaris-amd64 .

GOOS=windows 		GOARCH=386 			go build -o build/digital-overdose-bot-windows-386.exe .
GOOS=windows 		GOARCH=amd64 		go build -o build/digital-overdose-bot-windows-amd64.exe .