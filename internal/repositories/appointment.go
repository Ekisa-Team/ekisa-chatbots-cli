package repositories

import (
	"fmt"

	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/data"
)

type AppointmentRepository struct {
	Data *data.Database
}

// Insert appointments in "ChatBotCitas" table through "SpGrabarCitasChatBot" store procedure
func (r *AppointmentRepository) PrepareAppointments() (string, error) {
	q := "exec Quiron.dbo.SpGrabarCitasChatBot"

	stmt, err := r.Data.DB.Prepare(q)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	result, err := stmt.Exec()
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	message := fmt.Sprintf("Rows affected: %v", rowsAffected)
	return message, nil
}
