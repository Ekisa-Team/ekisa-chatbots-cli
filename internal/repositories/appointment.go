package repositories

import (
	"fmt"

	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/data"
	"github.com/Ekisa-Team/ekisa-chatbots-cli/pkg/appointment"
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

// Fetch chatbot appointments to be uploaded to the cloud
func (r *AppointmentRepository) FetchAppointments() ([]appointment.AppointmentChatbot, error) {
	q := "set nocount on; exec Quiron.dbo.ObtenerChatBotCitas"

	rows, err := r.Data.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []appointment.AppointmentChatbot
	for rows.Next() {
		var a appointment.AppointmentChatbot
		rows.Scan(
			&a.ID,
			&a.ClientID,
			&a.AppointmentNumber,
			&a.AppointmentDate,
			&a.AppointmentTime,
			&a.PacientID,
			&a.PacientName,
			&a.PacientPhoneNumber,
			&a.Sent,
			&a.Professional,
			&a.MedicalInstitution,
			&a.SentDate,
			&a.Answer,
			&a.AnswerDate,
		)
		appointments = append(appointments, a)
	}

	return appointments, nil
}
