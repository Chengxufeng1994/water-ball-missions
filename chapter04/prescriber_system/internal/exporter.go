package internal

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/prescriber_system/internal/domain"
)

type ExportType string

const (
	ExportTypeJSON ExportType = "json"
	ExportTypeCSV  ExportType = "csv"
	ExportTypeNone ExportType = "none"
)

type Exporter interface {
	Export(patientCase *domain.PatientCase, filename string) error
}

type JSONExporter struct {
	filename string
}

func NewJSONExporter(filename string) Exporter {
	return &JSONExporter{filename: filename}
}

func (e *JSONExporter) Export(patientCase *domain.PatientCase, filename string) error {
	exportedPrescription, err := json.MarshalIndent(patientCase, "", "  ")
	if err != nil {
		return err
	}

	filename = fmt.Sprintf("%s.json", filename)
	if err := os.WriteFile(filename, exportedPrescription, 0600); err != nil {
		return err
	}

	return nil
}

type CSVExporter struct {
	filename string
}

func NewCSVExporter(filename string) Exporter {
	return &CSVExporter{filename: filename}
}

func (e *CSVExporter) Export(pc *domain.PatientCase, filename string) error {
	// 1. 建立 CSV 檔案
	file, err := os.Create(fmt.Sprintf("%s.csv", filename))
	if err != nil {
		return err
	}
	defer file.Close()

	// 2. 建立 CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush() // 確保寫入內容

	// 3. 寫入標題列
	header := []string{"ID", "PatientID", "Symptoms", "PrescriptionName", "PotentialDisease", "Medicines", "Usage", "CreatedAt"}
	if err := writer.Write(header); err != nil {
		return err
	}

	p := pc.Prescription

	// 4. 寫入數據行
	var medicines string
	for _, med := range p.Medicines {
		medicines += med
		if med != p.Medicines[len(p.Medicines)-1] {
			medicines += ", "
		}
	}

	var symptoms string
	for _, symptom := range pc.Symptoms {
		symptoms += string(symptom)
		if symptom != pc.Symptoms[len(pc.Symptoms)-1] {
			symptoms += ", "
		}
	}

	record := []string{pc.ID, pc.PatientID, symptoms, p.Name, p.PotentialDisease, medicines, p.Usage, p.CreatedAt.Format("2006-01-02")}
	if err := writer.Write(record); err != nil {
		return err
	}

	return nil
}
