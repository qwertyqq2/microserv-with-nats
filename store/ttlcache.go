package store

import (
	"sync"
	"time"
)

type Item[T any] struct {
	value        T
	expirationAt int64
}

type TTLCache[K comparable, T any] struct {
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	items             map[K]Item[T]
	mu                sync.RWMutex
	stop              chan struct{}
}

func NewTTLCache[K comparable, T any](defaultExpiration, cleanupInterval time.Duration) *TTLCache[K, T] {
	cache := &TTLCache[K, T]{
		stop:              make(chan struct{}),
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
		items:             make(map[K]Item[T]),
	}

	go cache.gc()

	return cache
}

func (c *TTLCache[K, T]) Add(uid K, value T, expiryTime time.Duration) {
	if expiryTime == 0 {
		expiryTime = c.defaultExpiration
	}
	expireAt := time.Now().Add(expiryTime).UnixNano()

	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[uid] = Item[T]{
		value:        value,
		expirationAt: expireAt,
	}

}

func (c *TTLCache[K, T]) Get(uid K) (value T, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[uid]
	if !found {
		return
	}

	if item.expirationAt < time.Now().UnixNano() {
		return
	}

	return item.value, true
}

func (c *TTLCache[K, T]) Stop() {
	select {
	case <-c.stop:
		return
	default:
		select {
		case c.stop <- struct{}{}:
		}
	}

}

func (c *TTLCache[K, V]) gc() {
	ticker := time.NewTicker(c.cleanupInterval)
	for {
		select {
		case <-ticker.C:
			c.purge()

		case <-c.stop:
			ticker.Stop()
			return
		}
	}
}

func (c *TTLCache[K, V]) purge() {
	now := time.Now().UnixNano()
	for uid, val := range c.items {
		if val.expirationAt > now {
			delete(c.items, uid)
		}
	}
}
