package domain

type IndividualFactory interface {
	CreateIndividual(genes []Gene) Individual
}
