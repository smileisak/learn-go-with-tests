package main

import "sync"

// InMemoryPlayerStore store score in memory
type InMemoryPlayerStore struct {
	store map[string]int
	// A mutex is used to synchronize read/write access to the map
	lock sync.RWMutex
}

// GetPlayerScore return
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.store[name]
}

// RecordWin records a win in the InMemoryStore
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.store[name]++
}

// GetLeague returns the list of players
func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}

// NewInMemoryPlayerStore return a reference of a new New InMemoryPlayerStore
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.RWMutex{},
	}
}
