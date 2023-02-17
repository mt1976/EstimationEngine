@echo off
rem Starts the server
rem
set STARTER_APP=./...
set STARTER_PATH=C:\Dev\GitHub\ebEstimates
rem set APP_PORT=5050
set DB_SERVER=localhost
rem set DB_PORT=1433
set DB_USER=ebEstimates
set DB_PASSWORD=6LmYqp&*CBzo@s
set DB_DATABASE=EST_ALPHA
set STARTER_NAME=EB Estimates Server
title %STARTER_NAME% - %STARTER_APP% - %STARTER_PATH% - %APP_PORT%
cls
echo %STARTER_NAME% - %STARTER_APP% - %STARTER_PATH% - %APP_PORT%
cd %STARTER_PATH%
go version
go run %STARTER_APP%