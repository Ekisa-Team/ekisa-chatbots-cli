package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Ekisa-Team/kibot-cli/pkg/appointment"
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

func (s *AppointmentService) UploadAppointments(uri string, appointments appointment.AppointmentChatbotList) (string, error) {
	// Encode data
	encodedBody, _ := json.Marshal(appointments)
	requestBody := bytes.NewBuffer(encodedBody)

	// Perform request
	resp, err := http.Post(uri, "application/json", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// TODO: change color
	log.Printf("Server response: [%v] %v\n", resp.Status, uri)

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	sb := string(responseBody)
	return sb, nil
}

func (s *AppointmentService) UpdateAppointmentDeliveryInfo(appointmentID int, pacientID string, sent bool, sentDate time.Time) error {
	return s.Repository.UpdateAppointmentDeliveryInfo(appointmentID, pacientID, sent, sentDate)
}

func (s *AppointmentService) UpdateUserAnswer(appointmentID int, pacientID string, answer bool, answerDate time.Time) error {
	return s.Repository.UpdateUserAnswer(appointmentID, pacientID, answer, answerDate)
}
