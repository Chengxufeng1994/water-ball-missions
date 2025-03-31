package domain

import (
	"fmt"
	"time"
)

// PatientCase represents a medical case record
type PatientCase struct {
	ID           string       `json:"id"`           // Format: CASE[0-9]{9}
	PatientID    string       `json:"patient_id"`   // ID of the patient
	Prescription Prescription `json:"prescription"` // prescription
	Symptoms     []Symptom    `json:"symptoms"`     // List of symptoms
	CaseTime     time.Time    `json:"case_time"`    // When the case was diagnosed
}

// NewPatientCase creates a new patient case
func NewPatientCase(patientID string, symptoms []Symptom, prescription Prescription) (*PatientCase, error) {
	// Validate symptoms
	if len(symptoms) == 0 {
		return nil, ErrInvalidSymptom
	}

	return &PatientCase{
		ID:           generateCaseID(),
		PatientID:    patientID,
		Prescription: prescription,
		Symptoms:     symptoms,
		CaseTime:     prescription.CreatedAt,
	}, nil
}

func (p *PatientCase) String() string {
	return fmt.Sprintf("ID: %s, Symptoms: %v, Prescription: %v, CaseTime: %v",
		p.ID, p.Symptoms, p.Prescription, p.CaseTime)
}

// generateCaseID generates a unique ID for a new case
func generateCaseID() string {
	return "CASE" + time.Now().Format("20060102150405")
}
