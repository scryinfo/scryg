

@echo on


cd %~dp0
%~d0
set batPath=%cd%

cd %batPath%/dots
set dotPath=%cd%
cd %dotPath%/sutils & go build ./...

cd %batPath%
