#Before we can use the script, we have to make it executable with the chmod command:
#chmod +x ./go-executable-build.sh
#then we can use it  ./go-executable-build.sh yourpackage
#!/usr/bin/env bash

NOW=$(date +"%y%m")
figlet -f standard "Building ebEstimateEngine" 
echo Building Windows x86_64 
go-winres simply --icon ./assets/app.png
env GOOS=windows GOARCH=amd64 go build  -o "./ebEstimateEngine.exe" github.com/mt1976/ebEstimates
echo Building Crossplatform 
echo Building Windows x86_64 
go-winres simply --icon ./assets/app.png
env GOOS=windows GOARCH=amd64 go build  -o "./exec/windows/ebEstimateEngine_"$NOW".exe" github.com/mt1976/ebEstimates
echo Building MacOs x86_64 
go-winres simply --icon ./assets/app.png
env GOOS=darwin GOARCH=amd64 go build -o "./exec/mac/intel/ebEstimateEngine_"$NOW github.com/mt1976/ebEstimates 
echo Building MacOs arm64 
go-winres simply --icon ./assets/app.png
env GOOS=darwin GOARCH=arm64 go build -o "./exec/mac/apple/ebEstimateEngine_"$NOW github.com/mt1976/ebEstimates 
echo Done 