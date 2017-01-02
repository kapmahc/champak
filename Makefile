dst=dist

all:
	go build -ldflags "-s -X github.com/kapmahc/champak/web.Version=`git rev-parse --short HEAD` -X github.com/kapmahc/champak/web.BuildTime=`date +%FT%T%z`" -o $(dst)/champak main.go
	-cp -rv locales db themes $(dst)/

clean:
	-rm -rv $(dst)
	-rm -rv front-react/build
