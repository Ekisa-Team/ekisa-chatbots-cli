package appointment

import "time"

type AppointmentChatbot struct {
	ID                 uint      `json:"-"`
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
	AnswerDate         time.Time `json:"FechaHoraRespuesta,omitempty"`
}

type AppointmentChatbotList struct {
	Appointments []AppointmentChatbot `json:"Citas,omitempty"`
	ClientID     uint16               `json:"IdCliente,omitempty"`
}

type AppointmentChatbotResponse struct {
	IsSuccess bool        `json:"isSuccess"`
	Message   interface{} `json:"message"`
	Result    struct {
		ResponseCreate struct {
			IsSuccess bool   `json:"isSuccess"`
			Message   string `json:"message"`
			Result    []struct {
				IDCliente          int       `json:"idCliente"`
				NumeroCita         int       `json:"numeroCita"`
				FechaCita          time.Time `json:"fechaCita"`
				HoraCita           time.Time `json:"horaCita"`
				IDPaciente         string    `json:"idPaciente"`
				NombresPaciente    string    `json:"nombresPaciente"`
				Celular            string    `json:"celular"`
				Enviado            bool      `json:"enviado"`
				Profesional        string    `json:"profesional"`
				Institucion        string    `json:"institucion"`
				FechaHoraEnvio     time.Time `json:"fechaHoraEnvio"`
				Respuesta          string    `json:"respuesta"`
				FechaHoraRespuesta time.Time `json:"fechaHoraRespuesta"`
			} `json:"result"`
			MetaData interface{} `json:"metaData"`
		} `json:"responseCreate"`
		ResponseSend struct {
			IsSuccess bool   `json:"isSuccess"`
			Message   string `json:"message"`
			Result    []struct {
				IDCliente          int       `json:"idCliente"`
				NumeroCita         int       `json:"numeroCita"`
				FechaCita          time.Time `json:"fechaCita"`
				HoraCita           time.Time `json:"horaCita"`
				IDPaciente         string    `json:"idPaciente"`
				NombresPaciente    string    `json:"nombresPaciente"`
				Celular            string    `json:"celular"`
				Enviado            bool      `json:"enviado"`
				Profesional        string    `json:"profesional"`
				Institucion        string    `json:"institucion"`
				FechaHoraEnvio     time.Time `json:"fechaHoraEnvio"`
				Respuesta          string    `json:"respuesta"`
				FechaHoraRespuesta time.Time `json:"fechaHoraRespuesta"`
			} `json:"result"`
			MetaData interface{} `json:"metaData"`
		} `json:"responseSend"`
	} `json:"result"`
	MetaData interface{} `json:"metaData"`
}
