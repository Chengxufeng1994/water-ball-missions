package prescribehandler

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/prescriber_system/internal/domain"

type AttractivePrescribeHandler struct {
	next domain.PrescribeHandler
}

var _ domain.PrescribeHandler = (*AttractivePrescribeHandler)(nil)

func NewAttractivePrescribeHandler() *AttractivePrescribeHandler {
	return &AttractivePrescribeHandler{}
}
func (handler *AttractivePrescribeHandler) SetNext(next domain.PrescribeHandler) {
	handler.next = next
}

func (handler *AttractivePrescribeHandler) HandlePrescribe(patient domain.Patient, symptoms []domain.Symptom) (domain.Prescription, error) {
	if handler.checkAttractiveSymptoms(patient, symptoms) {
		return domain.NewAttractivePrescription()
	} else if handler.next != nil {
		return handler.next.HandlePrescribe(patient, symptoms)
	}

	return domain.Prescription{}, ErrUnknownDisease
}

func (c *AttractivePrescribeHandler) checkAttractiveSymptoms(patient domain.Patient, symptoms []domain.Symptom) bool {
	symptomSet := map[domain.Symptom]struct{}{}
	for _, symptom := range symptoms {
		symptomSet[symptom] = struct{}{}
	}

	requiredSymptoms := []domain.Symptom{domain.Sneeze}
	for _, requiredSymptom := range requiredSymptoms {
		if _, ok := symptomSet[requiredSymptom]; !ok {
			return false
		}
	}

	return patient.IsAdult() && patient.Gender.Equal(domain.GenderFemale)
}
