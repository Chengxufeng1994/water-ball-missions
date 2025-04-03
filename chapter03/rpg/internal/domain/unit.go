package domain

type Unit interface {
	TakeTurn(rpt *RPG)

	IsHero() bool
	IsAlive() bool
	IsDead() bool
	Detail() string
	LoseMagicPoint(amount int)
	GetID() int
	GetHP() int
	GetStrength() int
	AddStrength(amount int)
	LoseStrength(amount int)
	GetActions() []Action
	Actionable() bool
	SetActionable(enabled bool)
	Summon()
	Suicide()
	GetTroop() *Troop
	SetTroop(troop *Troop)
	RetrieveState(state State)
	GetCurrentState() State

	OnDamage(amount int)
	OnHeal(amount int)
	OnPoisoned(amount int)
	OnCurse(curser Unit)
}
