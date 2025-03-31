package domain

import (
	"encoding/json"
	"os"
	"sync"
)

type ErrPatientNotFound struct {
	id string
}

func (e ErrPatientNotFound) Error() string {
	return "patient not found: " + e.id
}

type PatientDatabase struct {
	path    string
	patents map[string]*Patient
	mu      sync.Mutex
}

func NewPatientDatabase(path string) *PatientDatabase {
	return &PatientDatabase{
		path:    path,
		patents: make(map[string]*Patient),
	}
}

func (db *PatientDatabase) LoadPatientData() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	// Load data from path
	data, err := os.ReadFile(db.path)
	if err != nil {
		return err
	}

	var patents map[string]*Patient
	if err := json.Unmarshal(data, &patents); err != nil {
		return err
	}

	db.patents = patents

	return nil
}

func (db *PatientDatabase) SavePatientData() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, err := json.Marshal(db.patents)
	if err != nil {
		return err
	}

	return os.WriteFile(db.path, data, 0644)
}

func (db *PatientDatabase) Save(patient *Patient) error {
	db.mu.Lock()
	db.patents[patient.ID] = patient
	db.mu.Unlock()

	return db.SavePatientData()
}

func (db *PatientDatabase) FindByID(id string) (*Patient, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	patient, ok := db.patents[id]
	if !ok {
		return nil, ErrPatientNotFound{id: id}
	}

	return patient, nil
}
