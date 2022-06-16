package bridge

type car interface {
	changeEngine(engine)
	speed(int)
}

type engine interface {
	acellerate(int)
	getSpeed() int
}

type ferrari struct {
	engine engine
}

func (f *ferrari) changeEngine(e engine) {
	f.engine = e
}

func (f *ferrari) speed(v int) {
	f.engine.acellerate(v)
}

type tesla struct {
	engine engine
}

func (t *tesla) changeEngine(e engine) {
	t.engine = e
}

func (t *tesla) speed(v int) {
	t.engine.acellerate(v)
}

type gasoline struct {
	speed int
}

func (g *gasoline) acellerate(a int) {
	g.speed = a
}

func (g *gasoline) getSpeed() int {
	return g.speed
}

type eletrical struct {
	speed int
}

func (e *eletrical) acellerate(a int) {
	e.speed = a
}

func (e *eletrical) getSpeed() int {
	return e.speed
}
