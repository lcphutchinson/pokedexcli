package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt	time.Time
	val		[]byte
}

type Cache struct {
	entries 	map[string]cacheEntry
	interval	time.Duration	
	mu		*sync.RWMutex
}

func (c Cache) Add(s string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[s] = cacheEntry{
		createdAt: time.Now(), 
		val: val,
	}
}

func (c Cache) Get(s string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.entries[s]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c Cache) reapLoop(t *time.Ticker) {
	for tick := range t.C {
		cutoff := tick.Add(-c.interval)
		c.mu.Lock()
		for key, entry := range c.entries {
			if entry.createdAt.Before(cutoff) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache (d time.Duration) Cache {
	c := Cache{
		entries: map[string]cacheEntry{},
		interval: d,
		mu: &sync.RWMutex{},
	}
	go c.reapLoop(time.NewTicker(d))
	return c
}


