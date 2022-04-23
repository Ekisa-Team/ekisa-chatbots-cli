package appointment

type Repository interface {
	PrepareAppointments() (string, error)
	FetchAppointments() ([]AppointmentChatbot, error)
}
