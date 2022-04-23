package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/Ekisa-Team/ekisa-chatbots-cli/utils"
	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func main() {
	setupInfo()
	setupFlags()
	setupActions()
	setupCommands()

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

// setupInfo configures CLI metadata
func setupInfo() {
	app.Name = "EkisaChatbots CLI"
	app.Version = "0.0.0"
}

// setupCommands configures allowed flags to setup environment
func setupFlags() {
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Load configuration from `file`",
		},
	}
}

// setupActions configures allowed actions to be catched
func setupActions() {
	app.Action = func(ctx *cli.Context) error {
		// check config path argument & load config
		if configPath := ctx.String("config"); configPath != "" {
			var c utils.Config
			config, err := c.LoadConfig(configPath)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("config.ClientID: %v\n", config.ClientID)
			fmt.Printf("config.ConnectionString: %v\n", config.ConnectionString)
			fmt.Printf("config.ApiEndpoint: %v\n", config.ApiEndpoint)

			return nil
		}

		return cli.Exit("You must specify a config file path", 1)
	}
}

// setupCommands configures allowed commands to be run in the CLI
func setupCommands() {
	app.Commands = []*cli.Command{
		{
			Name:    "sync",
			Aliases: []string{"s"},
			Usage:   "Get local database appointments and upload them to the cloud",
			Action: func(ctx *cli.Context) error {
				fmt.Println("Sync database appointments")
				return nil
			},
		},
		{
			Name:    "connectHub",
			Aliases: []string{"ch"},
			Usage:   "Connect to WebSocket and listen for appointments to be aupdated on local database",
			Action: func(ctx *cli.Context) error {
				fmt.Println("Connect to remote hub")
				return nil
			},
		},
	}
}
