::编译linux版本
::cd src
set GOPATH=%~dp0
set GOBIN=%~dp0/../../bin
go.exe build -o  ../bin/call-c.exe ./main/main.go
::cd ..