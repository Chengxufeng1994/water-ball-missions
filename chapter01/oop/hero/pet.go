package hero

type Pet struct {
	Name  string
	Owner *Hero
}

func NewPet(name string) *Pet {
	return &Pet{
		Name: name,
	}
}

func (p *Pet) Eat(fruit string) {
	if p.Owner == nil {
		return
	}
	p.Owner.SetHp(p.Owner.HP + 10)
}

func (p *Pet) SetOwner(owner *Hero) {
	p.Owner = owner
}
