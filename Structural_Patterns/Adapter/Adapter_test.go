package adapter

import "testing"

func TestAdapter(t *testing.T) {
	square := NewSquare(10)
	adapter := NewAdapter(square)
	roundHole := NewHole().insertObject(adapter)
	if roundHole.objects[0].getRadius() != 10 {
		t.Fatal("Expected to receive 10, but instead got", roundHole.objects[0].getRadius())
	}
}
