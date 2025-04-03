package domain

type Troop struct {
	Name  string
	Units []Unit
}

func NewTroop(id int, name string, units ...Unit) *Troop {
	for _, unit := range units {
		unit.SetTroopID(id)
	}

	return &Troop{
		Name:  name,
		Units: units,
	}
}

func (t *Troop) AddUnit(unit Unit) {
	t.Units = append(t.Units, unit)
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

func (t *Troop) IsHeroAlive() bool {
	return t.Units[0].IsAlive()
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
