package services

import (
	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/data"
	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/repositories"
)

type Proxy struct {
	AppointmentService *AppointmentService
}

// Returns facade containing the application services
func NewProxy() *Proxy {
	// Appointments
	appointmentService := &AppointmentService{
		Repository: &repositories.AppointmentRepository{
			Data: data.New(),
		},
	}

	return &Proxy{
		appointmentService,
	}
}
