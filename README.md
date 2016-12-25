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
govender sync
make
ls -l dist
```

## Editors

### Atom plugins

- git-plus
- go-plus
```bash
go get -u github.com/zmb3/gogetdoc
go get -u github.com/golang/lint/golint
```
- atom-beautify
- autosave(remember to enable it)

## Notes

### RabbitMQ

- The web UI is located at: <http://server-name:15672/>, (user "guest" is created with password "guest")

  ```bash
  rabbitmq-plugins enable rabbitmq_management
  ```

## Documents
- [Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/preface.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [gorm](http://jinzhu.me/gorm/)
- [mux](https://github.com/gorilla/mux)
- [render](https://github.com/unrolled/render)
- [negroni](https://github.com/urfave/negroni)
- [cli](https://github.com/urfave/cli)
- [viper](https://github.com/spf13/viper)
- [machinery](https://github.com/RichardKnop/machinery)
- [validator](https://github.com/go-playground/validator)
- [RabbitMQ](https://www.rabbitmq.com/getstarted.html)
