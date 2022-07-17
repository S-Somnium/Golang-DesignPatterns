package flyweight

import "testing"

func TestGame(t *testing.T) {
	game := newGame()
	game.addTerrorist()
	game.addTerrorist()
	game.addTerrorist()

	game.addCounterTerrorist()
	game.addCounterTerrorist()
	game.addCounterTerrorist()

}
