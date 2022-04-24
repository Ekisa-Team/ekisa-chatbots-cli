package upload

import (
	"log"
	"os"

	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/services"
	"github.com/Ekisa-Team/ekisa-chatbots-cli/pkg/appointment"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var proxy = services.NewProxy()

func NewCmdUpload() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload prepared appointments to ekisapp",
		Run:   uploadAppointments,
	}

	return cmd
}

func uploadAppointments(cmd *cobra.Command, args []string) {
	// Fetch appointments
	appointments, err := proxy.AppointmentService.FetchAppointments()
	if err != nil {
		log.Fatal(err)
	}

	// Validates if there are any appointments
	if len(appointments) == 0 {
		log.Println("No appointments were found")
		os.Exit(0)
	}

	// Upload appointments
	payload := appointment.AppointmentChatbotList{
		Appointments: appointments,
		ClientID:     uint16(viper.GetInt("client")),
	}

	msg, err := proxy.AppointmentService.UploadAppointments(viper.GetString("upload_webhook_uri"), payload)
	if err != nil {
		log.Fatal(err)
	}

	if msg != "" {
		log.Println(msg)
	}
}
