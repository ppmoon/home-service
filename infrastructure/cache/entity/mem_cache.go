package entity

import (
	"sync"
	"time"
)

var (
	memCacheOnce sync.Once
	memCache     *MemCache
)

type MemCache struct {
	mu    sync.Mutex
	store map[string]*Resource
}

type Resource struct {
	expiredTime time.Time
	item        []byte
}

func NewMemCache() *MemCache {
	memCacheOnce.Do(func() {
		memCache = &MemCache{
			mu:    sync.Mutex{},
			store: make(map[string]*Resource),
		}
	})
	return memCache
}

func (m *MemCache) Set(key string, value []byte, expired time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.store[key] = &Resource{
		expiredTime: time.Now().Add(expired),
		item:        value,
	}
}

func (m *MemCache) Get(key string) []byte {
	m.mu.Lock()
	defer m.mu.Unlock()
	value := m.store[key]
	if value.expiredTime.Before(time.Now()) {
		return nil
	}
	return value.item
}
