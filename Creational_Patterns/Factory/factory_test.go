package factory

import "testing"

func TestBigSword(t *testing.T) {
	sword, err := getSword("Big Sword")
	if err != nil {
		t.Fatal("Not expected to receive an error here:", err)
	}
	if sword.getName() != "Big Sword" {
		t.Fatal("Expected Big Sword but received:", sword.getName())
	}
	if sword.getPower() != 10 {
		t.Fatal("Expected 10 but received:", sword.getPower())
	}
	sword.uniquePower() // use his super power
	if sword.getPower() != 20 {
		t.Fatal("Expected 20 but received:", sword.getPower())
	}
}

func TestKnife(t *testing.T) {
	sword, err := getSword("Knife")
	if err != nil {
		t.Fatal("Not expected to receive an error here:", err)
	}
	if sword.getName() != "Knife" {
		t.Fatal("Expected Knife but received:", sword.getName())
	}
	if sword.getPower() != 1 {
		t.Fatal("Expected 10 but received:", sword.getPower())
	}
	sword.uniquePower() // use his super power
	if sword.getPower() != 0 {
		t.Fatal("Expected 20 but received:", sword.getPower())
	}
}

func TestError(t *testing.T) {
	if sword, err := getSword("Knife"); err == nil {
		t.Fatal("Expected to receive a error but instead received a sword:", sword.getName())
	}
}
