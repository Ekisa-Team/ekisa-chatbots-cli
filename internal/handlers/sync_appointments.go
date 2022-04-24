package handlers

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/config"
	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/services"
	"github.com/Ekisa-Team/ekisa-chatbots-cli/pkg/appointment"
	"github.com/urfave/cli/v2"
)

func SyncAppointments(ctx *cli.Context, proxy *services.Proxy) error {
	log.Println("Sync database appointments")

	// // Prepare appointments
	// prepareMsg, err := proxy.AppointmentService.PrepareAppointments()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(prepareMsg)

	// // Fetch appointments
	// appointments, err := proxy.AppointmentService.FetchAppointments()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(appointments)

	// // Upload appointments
	fmt.Println(config.ENV_CLIENT_ID, os.Getenv(config.ENV_CLIENT_ID))
	clientID, err := strconv.ParseInt(os.Getenv("CLIENT_ID"), 10, 16)
	if err != nil {
		log.Fatal(err)
	}

	l := appointment.AppointmentChatbotList{
		Appointments: []appointment.AppointmentChatbot{},
		ClientID:     uint16(clientID),
	}

	uploadMsg, err := proxy.AppointmentService.UploadAppointments(l)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(uploadMsg)

	return nil
}
