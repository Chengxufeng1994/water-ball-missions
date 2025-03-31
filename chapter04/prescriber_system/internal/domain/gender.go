package domain

// Gender represents the gender of a patient
type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"
)

func (g Gender) Equal(other Gender) bool {
	return g == other
}
