package decorator

type tesla interface {
	getPrice() int
}

type model3 struct {
	price int
}

func (m *model3) getPrice() int {
	return m.price
}

type longRange struct {
	tesla tesla
}

func (l *longRange) getPrice() int {
	return 20000 + l.tesla.getPrice()
}

type betterAI struct {
	tesla tesla
}

func (b *betterAI) getPrice() int {
	return 2000 + b.tesla.getPrice()
}
