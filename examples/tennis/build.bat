@ECHO OFF

rem This script sets the environment variables to be able to build the app, and copies the CSFML DLLs over if there aren't any in the folder

rem Edit the CSFML_PATH variable to match the path of your CSFML installation
set CSFML_PATH=C:\CSFML_2.5.1
rem Edit the COMPILER_NAME variable if you're not using gcc
set COMPILER_NAME=gcc

set CGO_CFLAGS="-I%CSFML_PATH%\include"
set CGO_LDFLAGS="-L%CSFML_PATH%\lib\%COMPILER_NAME%"

go get
if %ERRORLEVEL% NEQ 0 (echo go get failed && exit /b %ERRORLEVEL%)

go build
if %ERRORLEVEL% NEQ 0 (echo go build failed && exit /b %ERRORLEVEL%)

echo Build complete

if not exist "%~dp0*.dll" (
	echo No DLLs in folder, getting them from CSFML folder
	xcopy /s "%CSFML_PATH%\bin" "%~dp0"
	if %ERRORLEVEL% NEQ 0 (echo failed to copy DLLs && exit /b %ERRORLEVEL%)
)