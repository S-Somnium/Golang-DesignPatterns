package adapter

type rounderHole struct {
	objects []iRoundObject
}

func NewHole() *rounderHole {
	return &rounderHole{objects: []iRoundObject{}}
}

func (r *rounderHole) insertObject(object iRoundObject) *rounderHole {
	r.objects = append(r.objects, object)
	return r
}

type iRoundObject interface {
	getRadius() int
}

type squareAdapterToRoundObject struct {
	object *squareObject
}

func NewAdapter(s *squareObject) *squareAdapterToRoundObject {
	return &squareAdapterToRoundObject{object: s}
}

func (s *squareAdapterToRoundObject) getRadius() int {
	return s.object.getSide()
}

type squareObject struct {
	side int
}

func NewSquare(side int) *squareObject {
	return &squareObject{side: side}
}

func (s *squareObject) getSide() int {
	return s.side
}
