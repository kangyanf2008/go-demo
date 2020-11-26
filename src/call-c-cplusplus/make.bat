::cd src
::编译linux版本
set GOPATH=%~dp0
set GOBIN=%~dp0/../../bin
set GOOS=linux
set GOARCH=amd64
set GOHOSTOS=linux
set CGO_ENABLED=0

::go.exe build -o  ../../bin/call-c call-c-cplusplus/main/main.go
go.exe build -o  ../../bin/call-c ./main/main.go
::cd ..
