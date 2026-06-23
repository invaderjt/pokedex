package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	vals map[string]cacheEntry
	mu   sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		vals: make(map[string]cacheEntry),
	}
	go newCache.reapLoop(interval)
	return &newCache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.vals[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.vals[key]; !ok {
		return nil, false
	}
	return c.vals[key].val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.vals {
			if time.Since(entry.createdAt) > interval {
				delete(c.vals, key)
			}
		}
		c.mu.Unlock()
	}
}
