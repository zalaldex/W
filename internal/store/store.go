package store

import "sync"

// Store holds each user's selected mode in memory, keyed by Telegram user
// ID. It is safe for concurrent use.
type Store struct {
	mu    sync.RWMutex
	modes map[int64]string
}

// New creates an empty store.
func New() *Store {
	return &Store{modes: make(map[int64]string)}
}

// Get returns the mode for id, or "full" if none has been set.
func (s *Store) Get(id int64) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if m, ok := s.modes[id]; ok {
		return m
	}
	return "full"
}

// Set stores the mode for id.
func (s *Store) Set(id int64, m string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.modes[id] = m
}
