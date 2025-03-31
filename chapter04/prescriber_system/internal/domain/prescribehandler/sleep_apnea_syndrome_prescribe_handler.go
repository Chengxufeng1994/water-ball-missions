package prescribehandler

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/prescriber_system/internal/domain"

type SleepApneaSyndromePrescribeHandler struct {
	next domain.PrescribeHandler
}

var _ domain.PrescribeHandler = (*SleepApneaSyndromePrescribeHandler)(nil)

func NewSleepApneaSyndromePrescribeHandler() *SleepApneaSyndromePrescribeHandler {
	return &SleepApneaSyndromePrescribeHandler{}
}

func (handler *SleepApneaSyndromePrescribeHandler) SetNext(next domain.PrescribeHandler) {
	handler.next = next
}

func (handler *SleepApneaSyndromePrescribeHandler) HandlePrescribe(patient domain.Patient, symptoms []domain.Symptom) (domain.Prescription, error) {
	if handler.checkSleepApneaSyndromeSymptoms(patient, symptoms) {
		return domain.NewSleepApneaSyndromePrescription()
	} else if handler.next != nil {
		return handler.next.HandlePrescribe(patient, symptoms)
	}

	return domain.Prescription{}, ErrUnknownDisease
}

func (c *SleepApneaSyndromePrescribeHandler) checkSleepApneaSyndromeSymptoms(patient domain.Patient, symptoms []domain.Symptom) bool {
	symptomSet := map[domain.Symptom]struct{}{}
	for _, symptom := range symptoms {
		symptomSet[symptom] = struct{}{}
	}

	requiredSymptoms := []domain.Symptom{domain.Snore}
	for _, requiredSymptom := range requiredSymptoms {
		if _, ok := symptomSet[requiredSymptom]; !ok {
			return false
		}
	}

	return patient.IsObesity()
}
