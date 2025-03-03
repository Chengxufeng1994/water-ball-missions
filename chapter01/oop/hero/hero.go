package hero

import (
	"errors"
	"slices"
)

type Hero struct {
	Level  int
	Exp    int
	HP     int
	Pet    *Pet
	Guilds []*Guild
}

func NewHero() *Hero {
	return &Hero{
		Level:  1,
		Exp:    0,
		HP:     100,
		Guilds: make([]*Guild, 0),
	}
}

func RebuildHero(level, exp, hp int) *Hero {
	hero := &Hero{}

	hero.SetLevel(level)
	hero.SetExp(exp)
	hero.SetHp(hp)

	return hero
}

func (h *Hero) SetLevel(level int) error {
	if level < 1 {
		return errors.New("invalid level")
	}

	h.Level = level
	return nil
}
func (h *Hero) SetExp(exp int) error {
	if exp < 0 {
		return errors.New("invalid exp")
	}

	h.Exp += exp
	return nil
}
func (h *Hero) SetHp(hp int) error {
	if hp < 0 {
		return errors.New("invalid hp")
	}

	h.HP = hp
	return nil
}

func (h *Hero) SetPet(pet *Pet) {
	if pet != nil {
		pet.SetOwner(nil)
	}

	h.Pet = pet
	pet.SetOwner(h)
}

func (h *Hero) RemovePet() {
	if h.Pet != nil {
		h.Pet.SetOwner(nil)
	}
	h.Pet = nil
}

func (h *Hero) GainExp(exp int, levelSheet LevelSheet) error {
	if exp < 0 {
		return errors.New("invalid exp")
	}
	h.SetExp(exp)
	h.SetLevel(levelSheet.queryLevel(h.Exp))
	return nil
}

func (h *Hero) AddGuild(guild *Guild) {
	h.Guilds = append(h.Guilds, guild)
}

func (h *Hero) RemoveGuild(guild *Guild) {
	for i, g := range h.Guilds {
		if g == guild {
			h.Guilds = slices.Delete(h.Guilds, i, i+1)
		}
	}
}
