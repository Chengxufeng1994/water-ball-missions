package domain

import "time"

// PrescriptionDemand represents a prescription request from a patient
type PrescriptionDemand struct {
	PatientID string    `json:"patient_id"`
	Symptoms  []Symptom `json:"symptoms"`
}

// NewPrescriptionDemand creates a new prescription demand
func NewPrescriptionDemand(patientID string, symptoms []Symptom) PrescriptionDemand {
	return PrescriptionDemand{PatientID: patientID, Symptoms: symptoms}
}

// GetSymptoms returns all symptoms in the prescription demand
func (pd *PrescriptionDemand) GetSymptoms() []Symptom {
	return pd.Symptoms
}

// AddSymptom adds a new symptom to the prescription demand
func (pd *PrescriptionDemand) AddSymptom(symptom Symptom) {
	pd.Symptoms = append(pd.Symptoms, symptom)
}

// generatePrescriptionDemandID generates a unique ID for a new prescription demand
func generatePrescriptionDemandID() string {
	return "PD" + time.Now().Format("20060102150405")
}
