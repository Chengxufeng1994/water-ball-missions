package domain

import (
	"errors"
	"fmt"
	"time"
)

// Error messages for prescription validation
var (
	ErrPrescriptionNameLength = errors.New("處方名稱必須在4到30個字元之間")
	ErrPotentialDiseaseLength = errors.New("潛在疾病描述必須在3到100個字元之間")
	ErrNoMedicines            = errors.New("必須至少指定一個藥物")
	ErrUsageTooLong           = errors.New("使用方法不能超過1000個字元")
)

// Prescription represents a medical prescription
type Prescription struct {
	Name             string    `json:"name"`              // 4-30 characters
	PotentialDisease string    `json:"potential_disease"` // 3-100 characters
	Medicines        []string  `json:"medicines"`         // Each medicine: 3-30 characters
	Usage            string    `json:"usage"`             // 0-1000 characters
	CreatedAt        time.Time `json:"created_at"`
}

// NewPrescription creates a new prescription with the given information
func NewPrescription(
	name string,
	potentialDisease string,
	medicines []string,
	usage string,
) (Prescription, error) {
	// Validate prescription data
	if err := validatePrescription(name, potentialDisease, medicines, usage); err != nil {
		return Prescription{}, err
	}

	return Prescription{
		Name:             name,
		PotentialDisease: potentialDisease,
		Medicines:        medicines,
		Usage:            usage,
		CreatedAt:        time.Now(),
	}, nil
}

// validatePrescription validates the prescription data according to the rules
func validatePrescription(name, potentialDisease string, medicines []string, usage string) error {
	// Validate name (4-30 characters)
	if len(name) < 4 || len(name) > 30 {
		return ErrPrescriptionNameLength
	}

	// Validate potential disease (3-100 characters)
	if len(potentialDisease) < 3 || len(potentialDisease) > 100 {
		return ErrPotentialDiseaseLength
	}

	// Validate medicines
	if len(medicines) == 0 {
		return ErrNoMedicines
	}
	for i, medicine := range medicines {
		if len(medicine) < 3 || len(medicine) > 30 {
			return fmt.Errorf("第%d個藥物名稱必須在3到30個字元之間", i+1)
		}
	}

	// Validate usage (0-1000 characters)
	if len(usage) > 1000 {
		return ErrUsageTooLong
	}

	return nil
}

// String returns a string representation of the prescription
func (p Prescription) String() string {
	medicinesStr := ""
	for i, medicine := range p.Medicines {
		medicinesStr += fmt.Sprintf("%d. %s\n", i+1, medicine)
	}

	return fmt.Sprintf("處方名稱: %s\n潛在疾病: %s\n藥物:\n%s使用方法: %s\n開立時間: %s",
		p.Name,
		p.PotentialDisease,
		medicinesStr,
		p.Usage,
		p.CreatedAt.Format("2006-01-02 15:04:05"),
	)
}

// generatePrescriptionID generates a unique ID for a new prescription
func generatePrescriptionID() string {
	return "PRX" + time.Now().Format("20060102150405")
}

// a. 如果病患有打噴嚏、頭痛 (Headache) 和咳嗽 (Cough)等症狀的話，此時開出的治療處方是：
//
//	i. 		處方名字：清冠一號
//	ii. 	潛在疾病：新冠肺炎（專業學名：COVID-19）
//	iii.  食用藥物：清冠一號
//	iv. 	使用方法：將相關藥材裝入茶包裡，使用500 mL 溫、熱水沖泡悶煮1~3 分鐘後即可飲用。
func NewCOVID19Prescription() (Prescription, error) {
	return NewPrescription("清冠一號", "新冠肺炎", []string{"清冠一號"}, "將相關藥材裝入茶包裡，使用500 mL 溫、熱水沖泡悶煮1~3 分鐘後即可飲用。")
}

// b. 如果病患是正直 18 歲的女性，而且還打噴嚏 (sneeze) ****的話，此時開出的治療處方是：
//
//	i.		處方名字：青春抑制劑
//	ii.		潛在疾病：有人想你了 (專業學名：Attractive)
//	iii. 	食用藥物：假鬢角、臭味
//	iv.		使用方法：把假鬢角黏在臉的兩側，讓自己異性緣差一點，自然就不會有人想妳了。
func NewAttractivePrescription() (Prescription, error) {
	return NewPrescription("青春抑制劑", "有人想你了", []string{"假鬢角", "臭味"}, "把假鬢角黏在臉的兩側，讓自己異性緣差一點，自然就不會有人想妳了。")
}

// c. 如果病患的 BMI 大於 26，而且還打呼 (snore) 的話，此時開出的治療處方是：
//
//	i.		處方名字：打呼抑制劑
//	ii. 	潛在疾病：睡眠呼吸中止症（專業學名：SleepApneaSyndrome）
//	iii.	食用藥物：一捲膠帶
//	iv. 	使用方法：睡覺時，撕下兩塊膠帶，將兩塊膠帶交錯黏在關閉的嘴巴上，就不會打呼了。
func NewSleepApneaSyndromePrescription() (Prescription, error) {
	return NewPrescription("打呼抑制劑", "睡眠呼吸中止症", []string{"一捲膠帶"}, "睡覺時，撕下兩塊膠帶，將兩塊膠帶交錯黏在關閉的嘴巴上，就不會打呼了。")
}
