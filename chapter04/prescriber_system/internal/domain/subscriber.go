package domain

type Subscriber interface {
	Name() string
	HandlePrescription(prescription Prescription)
}
