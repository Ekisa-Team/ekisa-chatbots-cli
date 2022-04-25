package appointment

import "time"

type Repository interface {
	PrepareAppointments() (string, error)
	FetchAppointments() ([]AppointmentChatbot, error)
	UpdateAppointmentDeliveryInfo(appointmentID int, pacientID string, sent bool, sentDate time.Time) error
	UpdateUserAnswer(appointmentID int, pacientID string, answer bool, answerDate time.Time) error
}
