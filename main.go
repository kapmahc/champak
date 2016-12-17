package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/kapmahc/champak/engines/auth"
	_ "github.com/kapmahc/champak/engines/cms"
	_ "github.com/kapmahc/champak/engines/ops"
	_ "github.com/kapmahc/champak/engines/ops/mail"
	_ "github.com/kapmahc/champak/engines/ops/vpn"
	_ "github.com/kapmahc/champak/engines/reading"
	_ "github.com/kapmahc/champak/engines/shop"
	"github.com/kapmahc/champak/web"
	_ "github.com/lib/pq"
)

var version string

func main() {
	if err := web.Run(version); err != nil {
		log.Fatal(err)
	}
}