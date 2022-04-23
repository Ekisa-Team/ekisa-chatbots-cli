package handlers

import (
	"log"

	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/services"
	"github.com/urfave/cli/v2"
)

func ConnectHub(ctx *cli.Context, proxy *services.Proxy) error {
	log.Println("Connect to remote hub")
	return nil
}
