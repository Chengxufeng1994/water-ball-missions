package domain

type PrescribeHandler interface {
	HandlePrescribe(patient Patient, symptoms []Symptom) (Prescription, error)
	SetNext(handler PrescribeHandler)
}
