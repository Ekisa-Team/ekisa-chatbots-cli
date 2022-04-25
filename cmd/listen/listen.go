package listen

import (
	"fmt"

	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/services"
	"github.com/spf13/cobra"
)

var proxy = services.NewProxy()

func NewCmdListen() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listen",
		Short: "Listen for any appointment notifications from the remote server",
		Run:   listenForAppointments,
	}

	return cmd
}

// Prepare appointments
func listenForAppointments(cmd *cobra.Command, args []string) {
	fmt.Println("Listening for appointments")
}
