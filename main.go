package main

import (
	"log"

	_ "github.com/kapmahc/champak/engines/auth"
	_ "github.com/kapmahc/champak/engines/forum"
	_ "github.com/kapmahc/champak/engines/ops/mail"
	_ "github.com/kapmahc/champak/engines/ops/vpn"
	_ "github.com/kapmahc/champak/engines/reading"
	_ "github.com/kapmahc/champak/engines/shop"
	_ "github.com/kapmahc/champak/engines/site"
	"github.com/kapmahc/champak/web"
)

func main() {
	if err := web.Run(); err != nil {
		log.Fatal(err)
	}
}
