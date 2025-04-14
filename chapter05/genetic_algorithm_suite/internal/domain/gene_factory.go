package domain

type GeneFactory interface {
	CreateGene(num int) Gene
}
