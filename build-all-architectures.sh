VERSION=1.1.3-hotfix

#GOOS=darwin;        GOARCH=amd64; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-amd64 .
#GOOS=darwin;        GOARCH=arm64; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-arm64 .

#GOOS=dragonfly;     GOARCH=amd64; 	    go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-amd64 .

#GOOS=freebsd; 		GOARCH=386; 	 	go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-386 .
#GOOS=freebsd; 		GOARCH=amd64; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-amd64 .
#GOOS=freebsd; 		GOARCH=arm; 	 	go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-arm .

#GOOS=linux;	 		GOARCH=386; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-386 .
GOOS=linux;	 		GOARCH=amd64; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-amd64 .
GOOS=linux;	 		GOARCH=arm;	 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-arm .
GOOS=linux;	 		GOARCH=arm64; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-arm64 .
#GOOS=linux;	 		GOARCH=ppc64; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-ppc64 .
#GOOS=linux;	 		GOARCH=ppc64le; 	go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-ppc64le .
#GOOS=linux;	 		GOARCH=mips; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-mips .
#GOOS=linux;	 		GOARCH=mipsle; 	    go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-mipsle .
#GOOS=linux;	 		GOARCH=mips64; 	    go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-mips64 .
#GOOS=linux;	 		GOARCH=mips64le;    go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-mips64le .

#GOOS=netbsd;		GOARCH=386; 	    go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-386 .
#GOOS=netbsd;		GOARCH=amd64; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-amd64 .
#GOOS=netbsd;		GOARCH=arm; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-arm .

#GOOS=openbsd;		GOARCH=386; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-386 .
#GOOS=openbsd;		GOARCH=amd64; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-amd64 .
#GOOS=openbsd;		GOARCH=arm; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-arm .

#GOOS=plan9;			GOARCH=386; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-386 .
#GOOS=plan9;			GOARCH=amd64; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-amd64 .

#GOOS=solaris;		GOARCH=amd64; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-amd64 .

GOOS=windows; 		GOARCH=386; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-386.exe .
GOOS=windows; 		GOARCH=amd64; 		go build -o build/digital-overdose-bot-v$VERSION-$GOOS-$GOARCH-amd64.exe .