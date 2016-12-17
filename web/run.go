package web

import (
	"os"

	"github.com/urfave/cli"
)

// Run main entry
func Run(version string) error {

	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Version = version
	app.Usage = "CHAMPAK - A complete open source e-commerce solution by go-lang."
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{}

	for _, en := range engines {
		cmd := en.Shell()
		app.Commands = append(app.Commands, cmd...)
	}

	return app.Run(os.Args)
}
