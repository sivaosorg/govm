package mem

import (
	"container/list"
	"sync"
	"time"
)

// OnEvictFunc is a callback function type that gets called when an item is evicted from the cache.
type OnEvictFunc func(key string, value interface{})

// LRUCache represents the LRU cache.
type LRUCache struct {
	capacity    int
	cache       map[string]*list.Element
	list        *list.List
	mutex       sync.RWMutex
	onEvict     OnEvictFunc
	expiration  time.Duration
	stopCleanup chan struct{}
}

// lruMetadata represents the metadata of the least recently used item.
type lruMetadata struct {
	Key        string      `json:"key"`
	Value      interface{} `json:"value"`
	AccessTime time.Time   `json:"access_time"`
	Expiration time.Time   `json:"expiration"`
}

// cacheEntry represents a key-value pair in the cache.
type cacheEntry struct {
	Key        string      `json:"key"`
	Value      interface{} `json:"value"`
	Expiration time.Time   `json:"expiration"`
}
