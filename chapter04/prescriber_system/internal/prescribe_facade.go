package internal

import (
	"context"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/prescriber_system/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/prescriber_system/internal/domain/registry"
)

type PrescribeFacade struct {
	prescriber *domain.Prescriber
	database   *domain.PatientDatabase
}

func NewPrescribeFacade(ctx context.Context, patientDataPath string, prescribesPath string) (*PrescribeFacade, error) {
	reg := registry.NewPrescribeHandlerRegistry()

	database := domain.NewPatientDatabase(patientDataPath)
	if err := database.LoadPatientData(); err != nil {
		return nil, err
	}

	prescriber := domain.NewPrescriber(ctx, reg, database)
	if err := prescriber.LoadPrescribeList(prescribesPath); err != nil {
		return nil, err
	}

	return &PrescribeFacade{
		prescriber: prescriber,
		database:   database,
	}, nil
}

func (pf *PrescribeFacade) Prescribe(
	ctx context.Context,
	prescribeDemand domain.PrescriptionDemand,
	exportType ExportType,
	exportName string,
) (domain.Prescription, error) {
	patient, err := pf.database.FindByID(prescribeDemand.PatientID)
	if err != nil {
		return domain.Prescription{}, err
	}

	prescription, err := pf.prescriber.Prescribe(ctx, prescribeDemand)
	if err != nil {
		return domain.Prescription{}, err
	}

	patientCase, err := domain.NewPatientCase(prescribeDemand.PatientID, prescribeDemand.GetSymptoms(), prescription)
	if err != nil {
		return domain.Prescription{}, err
	}

	patient.AddPatientCase(patientCase)
	pf.database.Save(patient)

	err = pf.Export(ctx, patientCase, exportType, exportName)
	if err != nil {
		return domain.Prescription{}, err
	}

	return prescription, nil
}

func (pf *PrescribeFacade) Export(ctx context.Context, patientCase *domain.PatientCase, exportType ExportType, exportName string) error {
	if exportType == ExportTypeNone {
		return nil
	}

	switch exportType {
	case ExportTypeJSON:
		return NewJSONExporter(exportName).Export(patientCase, exportName)
	case ExportTypeCSV:
		return NewCSVExporter(exportName).Export(patientCase, exportName)
	}

	return nil
}
