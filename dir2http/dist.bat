rm -R .out
mkdir .out

set GOOS=linux
set GOARCH=amd64

go build -o .out/linux/dir2http

set GOOS=darwin
set GOARCH=amd64

go build -o .out/darwin/dir2http

set GOOS=windows
set GOARCH=amd64

go build -o .out/win/dir2http.exe