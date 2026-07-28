package store

import "sync"

import "github.com/zalaldex/W/internal/mode"

// Store holds each user's selected mode in memory, keyed by Telegram user
// ID. It is safe for concurrent use.
type Store struct {
	mu    sync.RWMutex
	modes map[int64]mode.Mode
}

// New creates an empty store.
func New() *Store {
	return &Store{modes: make(map[int64]mode.Mode)}
}

// Get returns the mode for id, or DefaultMode if none has been set.
func (s *Store) Get(id int64) mode.Mode {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if m, ok := s.modes[id]; ok {
		return m
	}
	return mode.DefaultMode
}

// Set stores the mode for id.
func (s *Store) Set(id int64, m mode.Mode) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.modes[id] = m
}
