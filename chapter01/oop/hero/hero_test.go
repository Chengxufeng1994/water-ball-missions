package hero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHero(t *testing.T) {
	hero := NewHero()
	levelSheet := LevelSheet{}
	hero.GainExp(1000, levelSheet)
	assert.Equal(t, 2, hero.Level)
	hero.GainExp(1000, levelSheet)
	assert.Equal(t, 3, hero.Level)
	hero.GainExp(1000, levelSheet)
	assert.Equal(t, 4, hero.Level)
}
