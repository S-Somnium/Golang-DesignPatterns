package builder

import "testing"

func TestCustomMonster(t *testing.T) {
	monster := NewMonster().withPower("fire").withSword("staff")
	if monster.power != "fire" && monster.sword != "staff" {
		t.Fatal("The creation of the monster failed, expected monster with fire staff, instead got", monster)
	}
}

func TestOrc(t *testing.T) {
	monster := NewOrc().withSword("Wood Sword")
	if monster.power != "strenght" && monster.sword != "Wood Sword" {
		t.Fatal("The creation of the monster failed, expected orc with strenght and Wood Sword, instead got", monster)
	}
}
