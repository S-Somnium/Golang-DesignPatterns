package abstract_factory

import "testing"

func TestDragonFactory(t *testing.T) {
	factory, err := ArmorFactory("dragon")
	if err != nil {
		t.Fatal("Not expected to receive an error here:", err)
	}
	if helmet := factory.makeHelmet(); helmet.getName() != "Dragon Helmet" {
		t.Fatal("Wrong piece of armor here! Expected Dragon Helmet, but got instead:", helmet.getName())
	}
	if plate := factory.makePlate(); plate.getName() != "Dragon Plate" {
		t.Fatal("Wrong piece of armor here! Expected Dragon Plate, but got instead:", plate.getName())
	}
	if boots := factory.makeBoots(); boots.getName() != "Dragon Boots" {
		t.Fatal("Wrong piece of armor here! Expected Dragon helmet, but got instead:", boots.getName())
	}
}

func TestIronFactory(t *testing.T) {
	factory, err := ArmorFactory("iron")
	if err != nil {
		t.Fatal("Not expected to receive an error here:", err)
	}
	if helmet := factory.makeHelmet(); helmet.getName() != "Iron Helmet" {
		t.Fatal("Wrong piece of armor here! Expected Iron Helmet, but got instead:", helmet.getName())
	}
	if plate := factory.makePlate(); plate.getName() != "Iron Plate" {
		t.Fatal("Wrong piece of armor here! Expected Iron Plate, but got instead:", plate.getName())
	}
	if boots := factory.makeBoots(); boots.getName() != "Iron Boots" {
		t.Fatal("Wrong piece of armor here! Expected Iron helmet, but got instead:", boots.getName())
	}
}

func TestError(t *testing.T) {
	if _, err := ArmorFactory("silver"); err == nil {
		t.Fatal("Expected to receive a error but instead received a factory")
	}
}
