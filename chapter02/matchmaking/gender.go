package main

type Gender int

const (
	Male Gender = iota + 1
	Female
)

func NewGender(gender int) Gender {
	switch gender {
	case int(Male):
		return Male
	case int(Female):
		return Female
	default:
		return Male
	}
}
