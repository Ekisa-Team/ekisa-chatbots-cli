package handlers

import (
	"log"

	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/services"
	"github.com/urfave/cli/v2"
)

func SyncAppointments(ctx *cli.Context, proxy *services.Proxy) error {
	log.Println("Sync database appointments")

	// Prepare appointments
	msg, err := proxy.AppointmentService.PrepareAppointments()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(msg)

	// Fetch appointments
	appointments, err := proxy.AppointmentService.FetchAppointments()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(appointments)

	return nil
}
