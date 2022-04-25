package repositories

import (
	"fmt"
	"time"

	"github.com/Ekisa-Team/ekisa-chatbots-cli/internal/data"
	"github.com/Ekisa-Team/ekisa-chatbots-cli/pkg/appointment"
)

type AppointmentRepository struct {
	Data *data.Database
}

// (Local) Insert appointments in "ChatBotCitas"
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

// (Local) Fetch chatbot appointments from "ChatBotCitas"
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

// (Local) Update chatbot appointment delivery info in "ChatBotCitas"
func (r *AppointmentRepository) UpdateAppointmentDeliveryInfo(appointmentID int, pacientID string, sent bool, sentDate time.Time) error {
	q := "UPDATE Quiron.dbo.ChatBotCitas SET Enviado = $3, FechaHoraEnvio = $4 WHERE NumeroCita = $1 AND IdPaciente = $2"
	_, err := r.Data.DB.Exec(q, appointmentID, pacientID, sent, sentDate)
	return err
}

// (Local) Update user's answer in "ChatBotCitas"
func (r *AppointmentRepository) UpdateUserAnswer(appointmentID int, pacientID string, answer bool, answerDate time.Time) error {
	q := "UPDATE Quiron.dbo.ChatBotCitas SET Respuesta = $3, FechaHoraRespuesta = $4 WHERE NumeroCita = $1 AND IdPaciente = $2"
	_, err := r.Data.DB.Exec(q, appointmentID, pacientID, answer, answerDate)
	return err
}
