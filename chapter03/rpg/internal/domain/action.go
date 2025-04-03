package domain

type Action interface {
	RequiredTargetType() int
	RequiredOfTargets() int
	MagicPointCost() int
	Amount() int
	Description(attacker Unit, targets []Unit) string
	Execute(rpg *RPG, attacker Unit, targets []Unit)
}
