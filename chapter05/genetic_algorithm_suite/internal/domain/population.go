package domain

type Population struct {
	Individuals []Individual
}

func NewPopulation(individuals []Individual) Population {
	return Population{Individuals: individuals}
}
