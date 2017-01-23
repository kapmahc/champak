package web

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	// Version version
	Version string
	// BuildTime build time
	BuildTime string
)

// Main main entry
func Main() error {
	app := cli.NewApp()

	app.Name = os.Args[0]
	app.Version = fmt.Sprintf("%s(%s)", Version, BuildTime)
	app.Usage = "CHAMPAK - A complete open source e-commerce solution by Go and Vue.js."
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{}

	for _, en := range engines {
		cmd := en.Shell()
		app.Commands = append(app.Commands, cmd...)
	}

	return app.Run(os.Args)
}
