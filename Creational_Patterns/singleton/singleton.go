package singleton

type musicBox struct {
	queue []string
}

var singleton *musicBox

func init() {
	singleton = &musicBox{
		queue: []string{},
	}
}

func GetMusicBox() *musicBox {
	return singleton
}

func (s *musicBox) queueMusic(music string) *musicBox {
	s.queue = append(s.queue, music)
	return s
}

func (s *musicBox) getQueue() []string {
	return s.queue
}
