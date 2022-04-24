package prepare

import (
	"log"

	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/services"
	"github.com/spf13/cobra"
)

var proxy = services.NewProxy()

func NewCmdPrepare() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prepare",
		Short: "Add appointments to ChatBotCitas table",
		Run:   prepareAppointments,
	}

	return cmd
}

// Prepare appointments
func prepareAppointments(cmd *cobra.Command, args []string) {
	msg, err := proxy.AppointmentService.PrepareAppointments()
	if err != nil {
		log.Fatal(err)
	}
	// TODO: change color to green
	log.Println(msg)
}
