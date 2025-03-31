package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/prescriber_system/internal"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/prescriber_system/internal/domain"
)

const (
	PATIENT_DATA_PATH   = "data/patient.json"
	PRESCRIBE_DATA_PATH = "data/prescribe.txt"
)

func main() {
	ctx := context.Background()
	prescriber, err := internal.NewPrescribeFacade(ctx, PATIENT_DATA_PATH, PRESCRIBE_DATA_PATH)
	if err != nil {
		panic(err)
	}

	patientID := "A123456789"
	tests := []struct {
		prescriptionDemand domain.PrescriptionDemand
		exportType         internal.ExportType
		exportName         string
	}{
		{
			prescriptionDemand: domain.NewPrescriptionDemand(patientID, []domain.Symptom{domain.Headache, domain.Cough, domain.Sneeze}),
			exportType:         internal.ExportTypeNone,
			exportName:         "",
		},
		{
			prescriptionDemand: domain.NewPrescriptionDemand(patientID, []domain.Symptom{domain.Snore}),
			exportType:         internal.ExportTypeJSON,
			exportName:         "data/result_1",
		},
		{
			prescriptionDemand: domain.NewPrescriptionDemand(patientID, []domain.Symptom{domain.Sneeze}),
			exportType:         internal.ExportTypeCSV,
			exportName:         "data/result_2",
		},
	}

	var wg sync.WaitGroup
	wg.Add(len(tests))
	for _, test := range tests {
		go func() {
			defer wg.Done()

			prescription, err := prescriber.Prescribe(context.Background(), test.prescriptionDemand, test.exportType, test.exportName)
			if err != nil {
				fmt.Printf("%v\n", err)
				fmt.Println()
				return
			}

			fmt.Println(prescription)
			fmt.Println()
		}()
	}

	wg.Wait()
}
