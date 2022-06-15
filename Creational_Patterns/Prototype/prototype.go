package prototype

import "fmt"

type iRobot interface {
	work()
	clone() iRobot
}

type robot struct {
	name     string
	function string
}

func NewRobot(name, function string) *robot {
	return &robot{name: name, function: function}
}

func (r *robot) work() {
	fmt.Println("working...")
}

func (r *robot) clone() *robot {
	return &robot{
		name:     r.name + "_clone",
		function: r.function,
	}
}
