package builder

// Pre build object from the builder
func NewOrc() *monster {
	return NewMonster().withHp(100).withPower("strenght").withWeakness("fire").withArmor("leather")
}
