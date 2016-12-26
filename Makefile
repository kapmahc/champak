dst=dist

build:
	go build -ldflags "-s -X main.version=`git rev-parse --short HEAD`" -o $(dst)/champak main.go
	-cp -rv locales db $(dst)/

clean:
	-rm -rv $(dst)
