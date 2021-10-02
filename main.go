package main

import (
	"log"
	"os"

	"gitlab.com/whype/gaming/builder/pkg/commands"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "builder",
		Usage: "Whype Rust Image Builder",
		Commands: []*cli.Command{
			{
				Name:   "check",
				Usage:  "Get the latest build details",
				Action: commands.Check,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "build-id",
						Usage: "Optionally specify a build id to pull manifest ids for.",
					},
					&cli.StringFlag{
						Name:  "app-id",
						Value: "258550",
					},
					&cli.StringFlag{
						Name:  "linux-depot-id",
						Value: "258552",
					},
					&cli.StringFlag{
						Name:  "common-depot-id",
						Value: "258554",
					},
					&cli.BoolFlag{
						Name: "write",
					},
					&cli.StringFlag{
						Name:  "branch",
						Value: "public",
					},
					&cli.BoolFlag{
						Name:  "export",
						Usage: "Export environment variables to feed into the builder.",
					},
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
