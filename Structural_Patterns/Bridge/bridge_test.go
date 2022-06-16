package bridge

import "testing"

func TestBridge(t *testing.T) {
	ev := eletrical{speed: 0}
	gs := gasoline{speed: 0}

	tesla := tesla{engine: &ev}
	tesla.speed(10)
	tesla.changeEngine(&gs)
	if tesla.engine.getSpeed() != 0 {
		t.Fatal("Expected to have speed 0, but instead got", tesla.engine.getSpeed())
	}
}
