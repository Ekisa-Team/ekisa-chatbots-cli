package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/config"
	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/data"
	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/handlers"
	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/services"
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

	fmt.Println("ci", os.Getenv(config.ENV_CLIENT_ID))

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
	os.Setenv(config.ENV_CLIENT_ID, "test")

	app.Action = func(ctx *cli.Context) error {
		// check config path argument
		if configPath := ctx.String("config"); configPath != "" {
			// load config from file
			var c config.Config
			cfg, err := c.ReadConfig(configPath)
			if err != nil {
				log.Fatal(err)
			}

			// load environment variables
			os.Setenv(config.ENV_CLIENT_ID, cfg.Application.ClientID)
			os.Setenv(config.ENV_CONN_STRING, cfg.Database.ConnectionString)
			os.Setenv(config.ENV_UPLOAD_APPOINTMENTS_URI, cfg.Webhooks.UploadAppointmentsUri)

			// show info logs if verbose flag is passed
			if verbose := ctx.Bool("verbose"); verbose {
				// print config status
				log.Println("Config  variables loaded")
				fmt.Printf("config.Application.ClientID: %v\n", os.Getenv(config.ENV_CLIENT_ID))
				fmt.Printf("config.Database.ConnectionString: %v\n", os.Getenv(config.ENV_CONN_STRING))
				fmt.Printf("config.Webhooks.UploadAppointmentsUri: %v\n", os.Getenv(config.ENV_UPLOAD_APPOINTMENTS_URI))
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
				fmt.Printf("config.Application.ClientID: %v\n", os.Getenv(config.ENV_CLIENT_ID))
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
