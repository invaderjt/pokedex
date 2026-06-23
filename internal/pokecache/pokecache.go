package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type cache struct {
	vals map[string]cacheEntry
	mu   sync.Mutex
}

func NewCache(interval time.Duration) *cache {
	newCache := cache{
		vals: make(map[string]cacheEntry),
	}
	go newCache.reapLoop(interval)
	return &newCache
}

func (c *cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.vals[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	return

}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.vals[key]; !ok {
		return nil, false
	}
	return c.vals[key].val, true
}

func (c *cache) reapLoop(interval time.Duration) {
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
