package main

import (
	"fmt"

	"github.com/pmarcais/transcode-sros/transsros"

	"github.com/urfave/cli/v2"
)

type appCfg struct {
	unflatConfFile string // file to transcode
	short          bool   // generate short line in service
	debug          bool   // for troubleshooting
}

// NewCLI defines the CLI flags and commands.
func NewCLI() *cli.App {
	appC := &appCfg{}
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "config-file",
			Aliases:     []string{"f"},
			Value:       "inventory.yml",
			Usage:       "Unflatten SROS configuration file",
			Destination: &appC.unflatConfFile,
		},
		&cli.BoolFlag{
			Name:        "shorten vprn line",
			Aliases:     []string{"s"},
			Value:       false,
			Usage:       "service vprn cli without service and customer name",
			Destination: &appC.short,
		},
		&cli.BoolFlag{
			Name:        "debug",
			Aliases:     []string{"d"},
			Value:       false,
			Usage:       "debug mode for troubleshooting",
			Destination: &appC.debug,
		},
	}

	app := &cli.App{
		Name:    "transcode-sros",
		Version: "dev",
		Usage:   "transcode a router CLI configuration file",
		Flags:   flags,
		Action: func(c *cli.Context) error {
			return appC.run()
		},
	}

	return app
}

func (app *appCfg) run() error {
	result := transsros.Transcode(app.unflatConfFile, app.short, app.debug)

	// Printing result to console
	for _, line := range result {
		fmt.Println(line)
	}

	return nil

}
