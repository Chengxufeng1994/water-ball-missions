package domain

type Troop struct {
	ID       int
	Name     string
	Units    []Unit
	newUnits []Unit
}

func NewTroop(id int, name string, units ...Unit) *Troop {
	troop := &Troop{
		ID:       id,
		Name:     name,
		newUnits: make([]Unit, 0),
	}

	for _, unit := range units {
		unit.SetTroop(troop)
	}

	troop.Units = units

	return troop
}

func (t *Troop) AddUnit(unit Unit) {
	unit.SetTroop(t)
	t.Units = append(t.Units, unit)
	t.newUnits = append(t.newUnits, unit)
}

func (t *Troop) GetAliveUnits() []Unit {
	var aliveUnits []Unit
	for _, unit := range t.Units {
		if unit.IsAlive() {
			aliveUnits = append(aliveUnits, unit)
		}
	}
	return aliveUnits
}

func (t *Troop) GetNewUnitsAndClear() []Unit {
	newUnits := t.newUnits
	t.newUnits = []Unit{}
	return newUnits
}

func (t *Troop) IsHeroAlive() bool {
	hero := t.Units[0]
	if hero.IsHero() {
		return hero.IsAlive()
	}

	return true
}

func (t *Troop) IsAnnihilated() bool {
	units := t.Units
	if len(units) == 0 {
		return false
	}

	for _, unit := range units {
		if unit.IsAlive() {
			return false
		}
	}

	return true
}
