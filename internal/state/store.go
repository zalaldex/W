package state

import (
	"sync"
	"telegram-monospace-bot/internal/formatter"
)

type Store struct {
	mu    sync.RWMutex
	modes map[int64]formatter.Mode
}

func NewStore() *Store {
	return &Store{modes: make(map[int64]formatter.Mode)}
}

func (s *Store) Get(id int64) formatter.Mode {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if m, ok := s.modes[id]; ok {
		return m
	}
	return formatter.DefaultMode
}

func (s *Store) Set(id int64, m formatter.Mode) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.modes[id] = m
}