package abstract_factory

import (
	"fmt"
)

// The methods that all armor factorys must have
type iArmorFactory interface {
	makeHelmet() iHelmet
	makePlate() iPlate
	makeBoots() iBoots
}

// The products of all factorys

type iHelmet interface {
	getName() string
}

type Helmet struct {
	name string
}

func (h *Helmet) getName() string {
	return h.name
}

type iPlate interface {
	getName() string
}

type Plate struct {
	name string
}

func (p *Plate) getName() string {
	return p.name
}

type iBoots interface {
	getName() string
}

type Boots struct {
	name string
}

func (b *Boots) getName() string {
	return b.name
}

// the abstract factory of factorys

func ArmorFactory(material string) (iArmorFactory, error) {
	switch material {
	case "dragon":
		return &DragonFactory{}, nil
	case "iron":
		return &IronFactory{}, nil
	default:
		return nil, fmt.Errorf("We dont have this type of armor here!!")
	}
}
