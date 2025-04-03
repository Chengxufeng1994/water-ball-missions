package domain

type Unit interface {
	TakeTurn(rpt *RPG)

	IsAlive() bool
	IsDead() bool
	Detail() string
	LoseMagicPoint(amount int)
	GetActions() []Action
	SetTroopID(troopID int)

	OnDamage(damage int)
}
