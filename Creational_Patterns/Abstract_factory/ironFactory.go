package abstract_factory

type ironHelmet struct {
	Helmet
}

func newIronHelmet() *ironHelmet {
	return &ironHelmet{
		Helmet: Helmet{
			name: "Iron Helmet",
		},
	}
}

type ironPlate struct {
	Plate
}

func newIronPlate() *ironPlate {
	return &ironPlate{
		Plate: Plate{
			name: "Iron Plate",
		},
	}
}

type ironBoots struct {
	Boots
}

func newIronBoots() *ironBoots {
	return &ironBoots{
		Boots: Boots{
			name: "Iron Boots",
		},
	}
}

type IronFactory struct {
}

func (d *IronFactory) makeHelmet() iHelmet {
	return newIronHelmet()
}

func (d *IronFactory) makePlate() iPlate {
	return newIronPlate()
}

func (d *IronFactory) makeBoots() iBoots {
	return newIronBoots()
}
