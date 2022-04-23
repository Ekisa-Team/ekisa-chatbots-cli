package appointment

type Repository interface {
	PrepareAppointments() (string, error)
}
