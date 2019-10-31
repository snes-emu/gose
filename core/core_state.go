package core

import "sync"

const (
	started stateStatus = iota
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
	s.status = started
}

func (s *state) Stop() {
	s.Lock()
	defer s.Unlock()
	s.status = started
}

func (s *state) SetStatus(status stateStatus) {
	s.Lock()
	defer s.Unlock()
	if status == paused || status == stopped || status == started {
		s.status = status
	}
}

func (s *state) Status() stateStatus {
	s.Lock()
	defer s.Unlock()
	return s.status
}
