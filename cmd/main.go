package main

import (
	"log"
	"os"
	"sort"

	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/data"
	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/handlers"
	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/services"
	"github.com/Ekisa-Team/ekisa-chatbots-cli/pkg/utils"
	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func main() {
	// Setup CLI
	setupInfo()
	setupFlags()
	setupActions()
	setupCommands()

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	// Validate database connection
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	// Run CLI
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// Attempt a graceful shutdown
	data.Close()
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
		&cli.BoolFlag{
			Name:  "verbose",
			Usage: "Show process logs",
		},
	}
}

// setupActions configures allowed actions to be catched
func setupActions() {
	app.Action = func(ctx *cli.Context) error {
		// check config path argument
		if configPath := ctx.String("config"); configPath != "" {
			// load config from file
			var c utils.Config
			config, err := c.LoadConfig(configPath)
			if err != nil {
				log.Fatal(err)
			}

			// load environment variables
			os.Setenv("CLIENT_ID", config.ClientID)
			os.Setenv("CONN_STRING", config.ConnectionString)
			os.Setenv("API_ENDPOINT", config.ApiEndpoint)

			// Show info logs if verbose flag is passed
			if verbose := ctx.Bool("verbose"); verbose {
				// print variables status
				log.Println("Environment variables loaded")
				log.Printf("os.Getenv(\"CLIENT_ID\"): %v\n", os.Getenv("CLIENT_ID"))
				log.Printf("os.Getenv(\"CONN_STRING\"): %v\n", os.Getenv("CONN_STRING"))
				log.Printf("os.Getenv(\"API_ENDPOINT\"): %v\n", os.Getenv("API_ENDPOINT"))
			}

			return nil
		}

		return cli.Exit("You must specify a config file path", 1)
	}
}

// setupCommands configures allowed commands to be run in the CLI
func setupCommands() {
	// Setup proxy connection
	proxy := services.NewProxy()

	app.Commands = []*cli.Command{
		{
			Name:    "sync",
			Aliases: []string{"s"},
			Usage:   "Get local database appointments and upload them to the cloud",
			Action: func(ctx *cli.Context) error {
				return handlers.SyncAppointments(ctx, proxy)
			},
		},
		{
			Name:    "connectHub",
			Aliases: []string{"ch"},
			Usage:   "Connect to WebSocket and listen for appointments to be aupdated on local database",
			Action: func(ctx *cli.Context) error {
				return handlers.ConnectHub(ctx, proxy)
			},
		},
	}
}
