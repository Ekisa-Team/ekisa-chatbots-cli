package cmd

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/Ekisa-Team/kibot-cli/cmd/prepare"
	"github.com/Ekisa-Team/kibot-cli/cmd/upload"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config string

	RootCmd = &cobra.Command{
		Use:   "kibot",
		Short: "Kibot CLI",
		Run: func(cmd *cobra.Command, args []string) {
			// fallback on default help if no args/flags are passed
			cmd.HelpFunc()(cmd, args)
		},
	}
)

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// persistent flags
	RootCmd.PersistentFlags().StringVarP(&config, "config", "c", "", "Path to config file (with extension)")

	// commands
	RootCmd.AddCommand(prepare.NewCmdPrepare())
	RootCmd.AddCommand(upload.NewCmdUpload())
}

func initConfig() {
	// if --config is passed, attempts to parse the config file
	if config != "" {
		// get the filepath
		abs, err := filepath.Abs(config)
		if err != nil {
			log.Fatal("Error reading filepath: ", err.Error())
		}

		// get the config name
		base := filepath.Base(abs)

		// get the directory path
		path := filepath.Dir(abs)

		// setup viper configuration
		viper.SetConfigName(strings.Split(base, ".")[0])
		viper.SetConfigType("yaml")
		viper.AddConfigPath(path)

		// Find and read the config file
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Using config file: " + viper.ConfigFileUsed())
		}
	} else {
		// if --config is not passed, adds multiple locations to search for config file
		viper.SetConfigName("kibot-config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME\\.config\\kibot")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("\033[1;34m%s\033[0m", "Using config file: "+viper.ConfigFileUsed())
		}
	}
}
