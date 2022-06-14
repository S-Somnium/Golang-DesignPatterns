package builder

// the logic of creating a monster
type monster struct {
	power    string
	weakness string
	armor    string
	sword    string
	hp       int
}

func NewMonster() *monster {
	return &monster{}
}

func (m *monster) withPower(power string) *monster {
	m.power = power
	return m
}

func (m *monster) withWeakness(weakness string) *monster {
	m.weakness = weakness
	return m
}

func (m *monster) withHp(hp int) *monster {
	m.hp = hp
	return m
}

func (m *monster) withArmor(armor string) *monster {
	m.armor = armor
	return m
}

func (m *monster) withSword(sword string) *monster {
	m.sword = sword
	return m
}
