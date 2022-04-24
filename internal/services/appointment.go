package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

func (s *AppointmentService) UploadAppointments(appointments appointment.AppointmentChatbotList) (string, error) {
	uri := os.Getenv("UPLOAD_APPOINTMENTS_URI")

	// Encode data
	encodedBody, _ := json.Marshal(appointments)
	requestBody := bytes.NewBuffer(encodedBody)

	// Perform request
	resp, err := http.Post(uri, "application/json", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	sb := string(responseBody)
	return sb, nil

}
