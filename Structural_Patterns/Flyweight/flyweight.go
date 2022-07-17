package flyweight

import "fmt"

type dress interface {
	getColor() string
	setColor(string)
}

type terroristDress struct {
	color string
}

func (t *terroristDress) getColor() string {
	return t.color
}

func (t *terroristDress) setColor(color string) {
	t.color = color
}

func newTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}

type counterTerroristDress struct {
	color string
}

func (c *counterTerroristDress) getColor() string {
	return c.color
}

func (c *counterTerroristDress) setColor(color string) {
	c.color = color
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "green"}
}

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]dress),
	}
)

type dressFactory struct {
	dressMap map[string]dress
}

func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == "tDress" {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == "ctDress" {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}

type player struct {
	dress      dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

type game struct {
	terrorists        []*player
	counterTerrorists []*player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*player, 1),
		counterTerrorists: make([]*player, 1),
	}
}

func (c *game) addTerrorist() {
	player := newPlayer("T", "tDress")
	c.terrorists = append(c.terrorists, player)
	return
}

func (c *game) addCounterTerrorist() {
	player := newPlayer("CT", "ctDress")
	c.counterTerrorists = append(c.counterTerrorists, player)
	return
}
