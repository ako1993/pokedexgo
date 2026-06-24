package pokecache

import (
	"errors"
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Cache_entries map[string]CacheEntry
	mu            sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	new_cache := &Cache{
		Cache_entries: make(map[string]CacheEntry),
	}

	ticker := time.NewTicker(interval)
	go func() {
		defer ticker.Stop()
		for range ticker.C {
			for _, entry := range new_cache.Cache_entries {
				now := time.Now()
				diff := entry.CreatedAt.Sub(now)
				if diff < interval {
					new_cache.RearLoop()
				}
			}
		}
	}()
	return new_cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Cache_entries[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	item, found := c.Cache_entries[key]
	defer c.mu.Unlock()

	if !found {
		print(errors.New("Cache miss: key not found"))
		return nil, false
	}
	return item.Val, true
}

func (c *Cache) RearLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Cache_entries = make(map[string]CacheEntry)
}
