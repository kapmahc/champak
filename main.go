package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/kapmahc/champak/engines/auth"
	_ "github.com/kapmahc/champak/engines/forum"
	_ "github.com/kapmahc/champak/engines/ops/mail"
	_ "github.com/kapmahc/champak/engines/ops/vpn"
	_ "github.com/kapmahc/champak/engines/reading"
	_ "github.com/kapmahc/champak/engines/shop"
	"github.com/kapmahc/champak/web"
)

func main() {
	if err := web.Main(); err != nil {
		log.Fatal(err)
	}
}
