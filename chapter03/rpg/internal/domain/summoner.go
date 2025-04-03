package domain

type Summoner interface {
	OnSummonedDead(amount int)
	AddSummoner(summoner Summoner)
	RemoveSummoner(summoner Summoner)
}
