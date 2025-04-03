package domain

type Action interface {
	RequiredOfTargets() int
	MagicPointCost() int
	Damage() int
	Execute(target Unit)
}
