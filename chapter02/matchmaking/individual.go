package main

import (
	"errors"
	"math"
	"strings"
)

var (
	ErrInvalidAge   = errors.New("invalid age")
	ErrInvalidIntro = errors.New("invalid intro")
)

type Individual struct {
	ID     int
	Name   string
	Gender Gender
	Age    int
	Intro  string
	Habits string
	Coord  Coord
}

func NewIndividual(id int, name string, gender Gender, age int, intro string, habits string, coord Coord) (Individual, error) {
	if age < 18 {
		return Individual{}, ErrInvalidAge
	}

	if len(intro) > 200 {
		return Individual{}, ErrInvalidIntro
	}

	return Individual{
		ID:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Intro:  intro,
		Habits: habits,
		Coord:  coord,
	}, nil
}

func (self Individual) GetHabits() []string {
	habits := strings.Split(self.Habits, ",")
	return habits
}

func (self Individual) CalculateHabitsScore(other Individual) int {
	habits := self.GetHabits()
	habitMap := make(map[string]struct{})
	for _, habit := range habits {
		habitMap[habit] = struct{}{}
	}

	intersection := 0
	for _, otherHabit := range other.GetHabits() {
		if _, ok := habitMap[otherHabit]; ok {
			intersection++
		}
	}

	return intersection
}

func (self Individual) CalculateDistance(other Individual) float64 {
	distance := math.Sqrt(math.Pow(self.Coord.Y-other.Coord.Y, 2) + math.Pow(self.Coord.X-other.Coord.X, 2))
	return distance
}
