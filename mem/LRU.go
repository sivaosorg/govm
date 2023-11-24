package mem

import (
	"container/list"
	"time"

	"github.com/sivaosorg/govm/utils"
)

func NewCacheEntry() *cacheEntry {
	return &cacheEntry{}
}

func (c *cacheEntry) SetKey(value string) *cacheEntry {
	c.Key = value
	return c
}

func (c *cacheEntry) SetValue(value interface{}) *cacheEntry {
	c.Value = value
	return c
}

func (c *cacheEntry) Json() string {
	return utils.ToJson(c)
}

func NewLRUMetadata() *lruMetadata {
	return &lruMetadata{
		AccessTime: time.Now(),
	}
}

func (l *lruMetadata) SetKey(value string) *lruMetadata {
	l.Key = value
	return l
}

func (l *lruMetadata) SetValue(value interface{}) *lruMetadata {
	l.Value = value
	return l
}

func (l *lruMetadata) SetAccessTime(value time.Time) *lruMetadata {
	l.AccessTime = value
	return l
}

func (l *lruMetadata) SetExpiration(value time.Time) *lruMetadata {
	l.Expiration = value
	return l
}

func (l *lruMetadata) Json() string {
	return utils.ToJson(l)
}

// NewLRUCache creates a new LRUCache with the specified capacity.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

// NewLRUCache creates a new LRUCache with the specified capacity and an optional eviction callback.
func NewLRUCacheEvict(capacity int, onEvict OnEvictFunc) *LRUCache {
	c := NewLRUCache(capacity)
	c.onEvict = onEvict
	return c
}

// NewLRUCache creates a new LRUCache with the specified capacity, an optional eviction callback,
// and an optional time-to-live for cache entries.
func NewLRUCacheExpiration(capacity int, expiration time.Duration) *LRUCache {
	c := NewLRUCache(capacity)
	c.SetExpiration(expiration)
	c.stopCleanup = make(chan struct{})
	// Start a background goroutine for periodic cache cleanup
	go c.startCleanup()
	return c
}

// Get retrieves a value from the cache based on the key.
func (c *LRUCache) Get(key string) (value interface{}, ok bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if element, exists := c.cache[key]; exists {
		// Check if the entry has expired
		if c.expiration > 0 && time.Now().After(element.Value.(*cacheEntry).Expiration) {
			// If the entry has expired, evict it from the cache
			c.evict(element)
			return nil, false
		}
		// Move the accessed element to the front of the list (most recently used)
		c.list.MoveToFront(element)
		return element.Value.(*cacheEntry).Value, true
	}
	return nil, false
}

// Set adds a key-value pair to the cache. If the cache is full, it removes the least recently used item.
func (c *LRUCache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, exists := c.cache[key]; exists {
		// Update the value and move the element to the front (most recently used)
		entry := element.Value.(*cacheEntry)
		entry.Value = value
		entry.Expiration = c.calculateExpiration()
		c.list.MoveToFront(element)
	} else {
		// Add a new element to the cache
		entry := &cacheEntry{
			Key:        key,
			Value:      value,
			Expiration: c.calculateExpiration(),
		}
		element := c.list.PushFront(entry)
		c.cache[key] = element

		// If the cache is full, remove the least recently used item
		if len(c.cache) > c.capacity {
			oldest := c.list.Back()
			if oldest != nil {
				c.evict(oldest)
			}
		}
	}
}

// Len returns the number of items in the cache.
func (c *LRUCache) Len() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return len(c.cache)
}

// IsEmpty checks if the cache is empty.
func (c *LRUCache) IsEmpty() bool {
	return c.Len() == 0
}

// Clear removes all items from the cache.
func (c *LRUCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache = make(map[string]*list.Element)
	c.list.Init()
}

// SetCapacity updates the capacity of the cache.
// Allows you to dynamically update the capacity of the cache.
// If the new capacity is less than the current number of items, it removes the excess items from the cache.
func (c *LRUCache) SetCapacity(capacity int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.capacity = capacity
	// If the new capacity is less than the current number of items, remove the excess items
	for len(c.cache) > c.capacity {
		oldest := c.list.Back()
		if oldest != nil {
			c.evict(oldest)
		}
	}
}

// GetAll returns all key-value pairs in the cache.
func (c *LRUCache) GetAll() map[string]interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	allEntries := make(map[string]interface{})
	for _, element := range c.cache {
		entry := element.Value.(*cacheEntry)
		allEntries[entry.Key] = entry.Value
	}
	return allEntries
}

// StopCleanup stops the background goroutine for periodic cache cleanup.
func (c *LRUCache) StopCleanup() {
	close(c.stopCleanup)
}

// SetOnEvict sets the eviction callback function.
func (c *LRUCache) SetOnEvict(onEvict OnEvictFunc) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.onEvict = onEvict
}

// SetExpiration updates the expiration time for cache entries.
func (c *LRUCache) SetExpiration(expiration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.expiration = expiration
}

// Contains checks if a key exists in the cache without updating its access time.
func (c *LRUCache) Contains(key string) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	_, exists := c.cache[key]
	return exists
}

// Remove removes a specific key from the cache.
func (c *LRUCache) Remove(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if element, exists := c.cache[key]; exists {
		c.evict(element)
	}
}

// GetLRU returns the least recently used key-value pair without removing it from the cache.
func (c *LRUCache) GetLRU() (key string, value interface{}, ok bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	oldest := c.list.Back()
	if oldest != nil {
		entry := oldest.Value.(*cacheEntry)
		return entry.Key, entry.Value, true
	}
	return "", nil, false
}

// Update updates the value associated with a specific key in the cache.
func (c *LRUCache) Update(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, exists := c.cache[key]; exists {
		entry := element.Value.(*cacheEntry)
		entry.Value = value
		entry.Expiration = c.calculateExpiration()
		c.list.MoveToFront(element)
	}
}

// GetLRUMetadata returns the metadata of the least recently used item without removing it from the cache.
func (c *LRUCache) GetLRUMetadata() (metadata lruMetadata, ok bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	oldest := c.list.Back()
	if oldest != nil {
		entry := oldest.Value.(*cacheEntry)
		l := NewLRUMetadata().
			SetKey(entry.Key).
			SetValue(entry.Value).
			SetExpiration(entry.Expiration).
			SetAccessTime(time.Now())
		return *l, true
	}
	return *NewLRUMetadata(), false
}

// IsExpired checks if a specific key has expired without updating its access time.
func (c *LRUCache) IsExpired(key string) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if element, exists := c.cache[key]; exists {
		entry := element.Value.(*cacheEntry)
		return c.expiration > 0 && time.Now().After(entry.Expiration)
	}
	return false
}

// Snapshot returns a snapshot of the current cache state.
func (c *LRUCache) Snapshot() []lruMetadata {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	snapshot := make([]lruMetadata, 0, len(c.cache))
	now := time.Now()
	for _, element := range c.cache {
		entry := element.Value.(*cacheEntry)
		l := NewLRUMetadata().
			SetKey(entry.Key).
			SetValue(entry.Value).
			SetAccessTime(now).
			SetExpiration(entry.Expiration)
		snapshot = append(snapshot, *l)
	}
	return snapshot
}

// GetMostRecentlyUsed returns the most recently used key-value pair without removing it from the cache.
func (c *LRUCache) GetMostRecentlyUsed() (metadata lruMetadata, ok bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	newest := c.list.Front()
	if newest != nil {
		entry := newest.Value.(*cacheEntry)
		l := NewLRUMetadata().
			SetKey(entry.Key).
			SetValue(entry.Value).
			SetExpiration(entry.Expiration).
			SetAccessTime(time.Now())
		return *l, true
	}
	return *NewLRUMetadata(), false
}

// ExtendExpiration extends the expiration time of a specific key in the cache.
func (c *LRUCache) ExtendExpiration(key string, extension time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, exists := c.cache[key]; exists {
		entry := element.Value.(*cacheEntry)
		entry.Expiration = entry.Expiration.Add(extension)
		c.list.MoveToFront(element)
	}
}

// RemainingExpiration returns the remaining time until expiration for a specific key.
func (c *LRUCache) RemainingExpiration(key string) (remainingTime time.Duration, ok bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if element, exists := c.cache[key]; exists {
		entry := element.Value.(*cacheEntry)
		if c.expiration > 0 {
			remainingTime = entry.Expiration.Sub(time.Now())
			return remainingTime, true
		}
	}
	return 0, false
}

// IsMostRecentlyUsed checks if a specific key is the most recently used item in the cache.
func (c *LRUCache) IsMostRecentlyUsed(key string) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if e := c.list.Front(); e != nil {
		entry := e.Value.(*cacheEntry)
		return entry.Key == key
	}
	return false
}

// evict evicts an element from the cache.
func (c *LRUCache) evict(element *list.Element) {
	// Invoke the eviction callback before removing the item
	if c.onEvict != nil {
		evictedEntry := element.Value.(*cacheEntry)
		c.onEvict(evictedEntry.Key, evictedEntry.Value)
	}
	delete(c.cache, element.Value.(*cacheEntry).Key)
	c.list.Remove(element)
}

// cleanupExpired removes expired entries from the cache.
func (c *LRUCache) cleanupExpired() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	for _, element := range c.cache {
		entry := element.Value.(*cacheEntry)
		if entry.Expiration.After(now) {
			// Entry has expired, evict it from the cache
			c.evict(element)
		}
	}
}

// startCleanup starts a background goroutine for periodic cache cleanup.
func (c *LRUCache) startCleanup() {
	ticker := time.NewTicker(c.expiration / 2) // Run cleanup at half the expiration interval
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.cleanupExpired()
		case <-c.stopCleanup:
			return
		}
	}
}

// calculateExpiration calculates the expiration time for a cache entry.
func (c *LRUCache) calculateExpiration() time.Time {
	if c.expiration > 0 {
		return time.Now().Add(c.expiration)
	}
	return time.Time{}
}
