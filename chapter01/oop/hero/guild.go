package hero

import (
	"errors"
	"slices"
)

type Guild struct {
	Name   string
	Heroes []*Hero
}

func NewGuild(name string, heroes []*Hero) *Guild {
	if len(heroes) > 10 || len(heroes) < 1 {
		return nil
	}
	g := &Guild{
		Name:   name,
		Heroes: heroes,
	}
	for _, h := range heroes {
		h.AddGuild(g)
	}

	return g
}

func (g *Guild) Join(hero *Hero) error {
	if len(g.Heroes) >= 10 {
		return errors.New("guild is full")
	}
	if g.Contains(hero) {
		return errors.New("hero already in guild")
	}
	g.Heroes = append(g.Heroes, hero)
	hero.AddGuild(g)
	return nil
}

func (g *Guild) Leave(hero *Hero) error {
	if len(g.Heroes) == 1 {
		return errors.New("guild is empty")
	}
	if !g.Contains(hero) {
		return errors.New("hero not in guild")
	}

	for i, h := range g.Heroes {
		if h == hero {
			g.Heroes = slices.Delete(g.Heroes, i, i+1)
		}
	}

	hero.RemoveGuild(g)

	return nil
}

func (g Guild) Contains(hero *Hero) bool {
	return slices.Contains(g.Heroes, hero)
}
