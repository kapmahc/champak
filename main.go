package main

import (
	"log"

	_ "github.com/kapmahc/champak/engines/auth"
	_ "github.com/kapmahc/champak/engines/cms"
	_ "github.com/kapmahc/champak/engines/ops/mail"
	_ "github.com/kapmahc/champak/engines/ops/vpn"
	_ "github.com/kapmahc/champak/engines/reading"
	_ "github.com/kapmahc/champak/engines/shop"
	"github.com/kapmahc/champak/web"
)

var version string

func main() {
	if err := web.Run(version); err != nil {
		log.Fatal(err)
	}
}
