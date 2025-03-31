package prescribehandler

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/prescriber_system/internal/domain"
)

type COVID19PrescribeHandler struct {
	next domain.PrescribeHandler
}

var _ domain.PrescribeHandler = (*COVID19PrescribeHandler)(nil)

func NewCovid19PrescribeHandler() *COVID19PrescribeHandler {
	return &COVID19PrescribeHandler{}
}

func (handler *COVID19PrescribeHandler) SetNext(next domain.PrescribeHandler) {
	handler.next = next
}

func (handler *COVID19PrescribeHandler) HandlePrescribe(patient domain.Patient, symptoms []domain.Symptom) (domain.Prescription, error) {
	if handler.checkCovid19Symptoms(patient, symptoms) {
		return domain.NewCOVID19Prescription()
	} else if handler.next != nil {
		return handler.next.HandlePrescribe(patient, symptoms)
	}

	return domain.Prescription{}, ErrUnknownDisease
}

func (c *COVID19PrescribeHandler) checkCovid19Symptoms(patient domain.Patient, symptoms []domain.Symptom) bool {
	requiredSymptoms := map[domain.Symptom]struct{}{
		domain.Headache: {},
		domain.Cough:    {},
		domain.Sneeze:   {},
	}

	if len(symptoms) != len(requiredSymptoms) {
		return false
	}

	symptomSet := make(map[domain.Symptom]struct{}, len(symptoms))
	for _, symptom := range symptoms {
		symptomSet[symptom] = struct{}{}
	}

	for requiredSymptom := range requiredSymptoms {
		if _, ok := symptomSet[requiredSymptom]; !ok {
			return false
		}
	}

	return true
}
