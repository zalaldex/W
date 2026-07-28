package main

import "sync"

// store holds each user's selected mode in memory, keyed by Telegram user
// ID. It is safe for concurrent use.
type store struct {
	mu    sync.RWMutex
	modes map[int64]Mode
}

// newStore creates an empty store.
func newStore() *store {
	return &store{modes: make(map[int64]Mode)}
}

// get returns the mode for id, or DefaultMode if none has been set.
func (s *store) get(id int64) Mode {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if m, ok := s.modes[id]; ok {
		return m
	}
	return DefaultMode
}

// set stores the mode for id.
func (s *store) set(id int64, m Mode) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.modes[id] = m
}
