package pokecache

import (
	"time"
	"sync"
)


type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	Cache_entries map[string]cacheEntry
	mu *sync.RWMutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Cache_entries[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.Cache_entries[key]
	if !ok {
		return nil, false
	} else {
		return entry.val, true
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	time_now := time.Now()
	for key, entry := range c.Cache_entries {
		if entry.createdAt.Before(time_now.Add(-interval)) {
			delete(c.Cache_entries, key)
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Cache_entries: map[string]cacheEntry{},
		mu: &sync.RWMutex{},
	}

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			c.reapLoop(interval)
		}
	} ()

	return c
}