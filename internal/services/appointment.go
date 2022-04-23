package services

import (
	"github.com/Ekisa-Team/ekisa-chatbots-cli/pkg/appointment"
)

type AppointmentService struct {
	Repository appointment.Repository
}

func (s *AppointmentService) PrepareAppointments() (string, error) {
	return s.Repository.PrepareAppointments()
}

func (s *AppointmentService) FetchAppointments() ([]appointment.AppointmentChatbot, error) {
	return s.Repository.FetchAppointments()
}
