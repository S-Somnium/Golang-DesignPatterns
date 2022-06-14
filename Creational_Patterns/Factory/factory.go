package factory

import (
	"fmt"
)

// the methods that all sword must have to be a sword
type iSword interface {
	cut()
	uniquePower()
	getPower() int
	setPower(power int)
	getName() string
	setName(name string)
}

// the basic implementation of all sword, notice that the method uniquePower() wasnt implemented here.
type sword struct {
	name  string
	power int
}

func (s *sword) setName(name string) {
	s.name = name
}

func (s *sword) getName() string {
	return s.name
}

func (s *sword) setPower(power int) {
	s.power = power
}

func (s *sword) getPower() int {
	return s.power
}

func (s *sword) cut() {
	fmt.Println("You are being cuted by:", s.name)
}

// bigSword with a unique power
type bigSword struct {
	sword
}

func newBigSword() *bigSword {
	return &bigSword{
		sword: sword{
			name:  "Big Sword",
			power: 10,
		},
	}
}

func (s *bigSword) uniquePower() {
	s.setPower(20)
	fmt.Println("Berserker mode ON")
}

//knife with a unique power

type knife struct {
	sword
}

func newKnife() *knife {
	return &knife{
		sword: sword{
			name:  "Knife",
			power: 1,
		},
	}
}

func (s *knife) uniquePower() {
	s.setPower(100)
	fmt.Println("Shadow cut!!!!")
	s.setPower(0)
	fmt.Println("The knife broke")
}

// sword factory

func getSword(sword string) (iSword, error) {
	switch sword {
	case "Big Sword":
		return newBigSword(), nil
	case "Knife":
		return newKnife(), nil
	}
	return nil, fmt.Errorf("We dont have this type of sword here!!")
}
