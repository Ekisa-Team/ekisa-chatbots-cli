package appointment

import "time"

type AppointmentChatbot struct {
	ID                 uint      `json:"Id,omitempty"`
	ClientID           uint16    `json:"IdCliente,omitempty"`
	PacientID          string    `json:"IdPaciente,omitempty"`
	PacientName        string    `json:"NombresPaciente,omitempty"`
	PacientPhoneNumber string    `json:"Celular,omitempty"`
	AppointmentNumber  float32   `json:"NumeroCita,omitempty"`
	AppointmentDate    time.Time `json:"FechaCita,omitempty"`
	AppointmentTime    time.Time `json:"HoraCita,omitempty"`
	Professional       string    `json:"Profesional,omitempty"`
	MedicalInstitution string    `json:"Institucion,omitempty"`
	Sent               bool      `json:"Enviado,omitempty"`
	SentDate           time.Time `json:"FechaHoraEnvio,omitempty"`
	Answer             string    `json:"Respuesta,omitempty"`
	AnswerDate         bool      `json:"FechaHoraRespuesta,omitempty"`
}

type AppointmentChatbotList struct {
	Appointments []AppointmentChatbot `json:"Citas,omitempty"`
	ClientID     uint16               `json:"IdCliente,omitempty"`
}
