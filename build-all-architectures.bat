@echo off
echo Building artifacts in /build
set VERSION=%1

:: call :build darwin				amd64
:: call :build darwin				arm64

:: call :build dragonfly		amd64

:: call :build freebsd			386
:: call :build freebsd			amd64
:: call :build freebsd			arm

:: call :build linux				386
call :build linux amd64
call :build linux arm
call :build linux arm64
:: call :build linux 		ppc64    
:: call :build linux 		ppc64le  
:: call :build linux 		mips     
:: call :build linux 		mipsle   
:: call :build linux 		mips64   
:: call :build linux 		mips64le 

:: call :build netbsd 	386      
:: call :build netbsd 	amd64    
:: call :build netbsd 	arm      

:: call :build openbsd 	386    
:: call :build openbsd 	amd64  
:: call :build openbsd 	arm    
 
:: call :build plan9 		386      
:: call :build plan9 		amd64    
 
:: call :build solaris 	amd64  

call :build windows 386
call :build windows amd64

echo Done!
exit /b

:build
set GOOS=%~1
set GOARCH=%~2
set EXT=
if "%GOOS%"=="windows" (
	set EXT=.exe
)
go build -o build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH%%EXT% .
echo Built build/digital-overdose-bot-v%VERSION%-%GOOS%-%GOARCH%%EXT%