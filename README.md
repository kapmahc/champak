# Champak

A complete open source e-commerce solution by go-lang.

## Installing

```bash
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
gvm install go1.7.4 -B
gvm use go1.7.4 --default

go get -u github.com/kardianos/govendor
go get -u github.com/kapmahc/champak
cd $GOPATH/src/github.com/kapmahc/champak
govendor sync
make
ls -l dist
```

## Editors

### Atom plugins

- git-plus
- go-plus
- atom-beautify
- autosave(remember to enable it)

## Documents

- [gorm](http://jinzhu.me/gorm/)
- [gin](https://github.com/gin-gonic/gin/)
- [cli](https://github.com/urfave/cli)
- [viper](https://github.com/spf13/viper)
