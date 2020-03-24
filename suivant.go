package suivant

import (
	"time"
)

type Suivant struct {
	names []string
	n     int
	delay time.Duration
	next  chan string
	stop  chan interface{}
}

func New(names []string, delay time.Duration) *Suivant {
	next, stop := make(chan string), make(chan interface{})
	return &Suivant{names: names, delay: delay, next: next, stop: stop}
}

func (s *Suivant) Next() <-chan string {
	return s.next
}

func (s *Suivant) Start() <-chan string {
	go s.start()
	return s.next
}

func (s *Suivant) Stop() {
	s.stop <- struct{}{}
}

func (s *Suivant) sendNext() {
	s.next <- s.names[s.n%len(s.names)]
	s.n++
}

func (s *Suivant) start() {
	s.sendNext()
	for {
		select {
		case <-time.After(s.delay):
			s.sendNext()
		case <-s.stop:
			return
		}
	}
}
