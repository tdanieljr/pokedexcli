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
	m        map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = cacheEntry{time.Now(), val}
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.m[key]
	if !ok {
		return nil, ok
	}
	return val.val, ok
}
func (c *Cache) reapLoop() {
	t := time.NewTicker(c.interval)
	for range t.C {
		c.mu.Lock()
		for k, v := range c.m {
			if time.Since(v.createdAt) > c.interval {
				delete(c.m, k)
			}
		}
		c.mu.Unlock()

	}
}

func NewCache(interval time.Duration) *Cache {
	m := make(map[string]cacheEntry)
	c := &Cache{m: m, interval: interval}
	go c.reapLoop()
	return c
}
