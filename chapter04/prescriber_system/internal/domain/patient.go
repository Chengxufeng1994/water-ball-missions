package domain

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// Error messages for patient validation
var (
	ErrInvalidIDFormat    = errors.New("身分證字號格式錯誤：必須以大寫英文字母開頭，後接9位數字")
	ErrInvalidNameLength  = errors.New("姓名長度必須在1到30個字元之間")
	ErrInvalidNameFormat  = errors.New("姓名只能包含英文字母")
	ErrInvalidGender      = errors.New("性別必須是 male 或 female")
	ErrInvalidAge         = errors.New("年紀必須在1到180歲之間")
	ErrInvalidHeight      = errors.New("身高必須在1到500公分之間")
	ErrInvalidWeight      = errors.New("體重必須在1到500公斤之間")
	ErrInvalidCaseID      = errors.New("病例ID格式錯誤：必須以CASE開頭，後接9位數字")
	ErrInvalidCaseContent = errors.New("病例內容不能為空")
	ErrCaseNotFound       = errors.New("找不到指定的病例")
	ErrCaseAlreadyExists  = errors.New("病例ID已存在")
	ErrInvalidSymptom     = errors.New("症狀描述不能為空")
)

// Patient represents a patient in the system
type Patient struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Age          int            `json:"age"`
	Gender       Gender         `json:"gender"`
	Height       float64        `json:"height"` // in centimeters
	Weight       float64        `json:"weight"` // in kilograms
	PatientCases []*PatientCase `json:"patient_cases"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

// validatePatient validates the patient data according to the rules
func validatePatient(id, name string, gender Gender, age int, height, weight float64) error {
	// Validate ID format
	if !regexp.MustCompile(`^[A-Z][0-9]{9}$`).MatchString(id) {
		return ErrInvalidIDFormat
	}

	// Validate name
	if len(name) < 1 || len(name) > 30 {
		return ErrInvalidNameLength
	}
	if !regexp.MustCompile(`^[A-Za-z\s]+$`).MatchString(name) {
		return ErrInvalidNameFormat
	}

	// Validate gender
	if gender != GenderMale && gender != GenderFemale {
		return ErrInvalidGender
	}

	// Validate age
	if age < 1 || age > 180 {
		return ErrInvalidAge
	}

	// Validate height
	if height < 1 || height > 500 {
		return ErrInvalidHeight
	}

	// Validate weight
	if weight < 1 || weight > 500 {
		return ErrInvalidWeight
	}

	return nil
}

// NewPatient creates a new patient with the given information
func NewPatient(
	id string,
	name string,
	gender Gender,
	age int,
	height float64,
	weight float64,
) (*Patient, error) {
	// Validate patient data
	if err := validatePatient(id, name, gender, age, height, weight); err != nil {
		return nil, err
	}

	now := time.Now()
	return &Patient{
		ID:           id,
		Name:         name,
		Gender:       gender,
		Age:          age,
		Height:       height,
		Weight:       weight,
		PatientCases: make([]*PatientCase, 0),
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}

// AddPatientCase adds a new patient case to the patient's record
func (p *Patient) AddPatientCase(patientCase *PatientCase) {
	p.PatientCases = append(p.PatientCases, patientCase)
	p.UpdatedAt = time.Now()
}

// IsAdult returns true if the patient is an adult
func (p *Patient) IsAdult() bool {
	return p.Age >= 18
}

// IsTooFat returns true if the patient is too fat
func (p *Patient) IsObesity() bool {
	return p.CalculateBMI() > 26
}

// CalculateBMI returns the patient's body mass index (BMI)
func (p *Patient) CalculateBMI() float64 {
	heightInMeters := p.Height / 100
	return p.Weight / (heightInMeters * heightInMeters)
}

// String returns a string representation of the patient
func (p *Patient) String() string {
	casesStr := ""
	for _, c := range p.PatientCases {
		casesStr += c.String() + "\n"
	}

	return fmt.Sprintf("病患資料:\nID: %s\n姓名: %s\n性別: %s\n年紀: %d\n身高: %.1f cm\n體重: %.1f kg\n病例:\n%s",
		p.ID,
		p.Name,
		p.Gender,
		p.Age,
		p.Height,
		p.Weight,
		casesStr,
	)
}

// GeneratePatientID generates a patient ID with 1 capital letter followed by 9 digits
func GeneratePatientID() string {
	letter := fmt.Sprintf("%c", 'A'+rand.Intn(26))       // Generate a random capital letter
	digits := fmt.Sprintf("%09d", rand.Intn(1000000000)) // Generate 9 random digits
	return letter + digits
}
