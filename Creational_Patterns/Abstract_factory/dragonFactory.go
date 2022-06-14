package abstract_factory

type dragonHelmet struct {
	Helmet
}

func newDragonHelmet() *dragonHelmet {
	return &dragonHelmet{
		Helmet: Helmet{
			name: "Dragon Helmet",
		},
	}
}

type dragonPlate struct {
	Plate
}

func newDragonPlate() *dragonPlate {
	return &dragonPlate{
		Plate: Plate{
			name: "Dragon Plate",
		},
	}
}

type dragonBoots struct {
	Boots
}

func newDragonBoots() *dragonBoots {
	return &dragonBoots{
		Boots: Boots{
			name: "Dragon Boots",
		},
	}
}

type DragonFactory struct {
}

func (d *DragonFactory) makeHelmet() iHelmet {
	return newDragonHelmet()
}

func (d *DragonFactory) makePlate() iPlate {
	return newDragonPlate()
}

func (d *DragonFactory) makeBoots() iBoots {
	return newDragonBoots()
}
