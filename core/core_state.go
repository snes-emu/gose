package core

import "sync"

const (
	running stateStatus = iota
	paused
	stopped
)

type stateStatus int

type state struct {
	sync.Mutex
	status stateStatus
}

func NewState() *state {
	return &state{status: paused}
}

func (s *state) Pause() {
	s.Lock()
	defer s.Unlock()
	s.status = paused
}

func (s *state) Start() {
	s.Lock()
	defer s.Unlock()
	s.status = running
}

func (s *state) Stop() {
	s.Lock()
	defer s.Unlock()
	s.status = running
}

func (s *state) SetStatus(status stateStatus) {
	s.Lock()
	defer s.Unlock()
	if status == paused || status == stopped || status == running {
		s.status = status
	}
}

func (s *state) Status() stateStatus {
	s.Lock()
	defer s.Unlock()
	return s.status
}
