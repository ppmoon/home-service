package lock

import (
	"sync"
	"time"
)

var (
	multiRoutineLockOnce sync.Once
	multiRoutineLock     *MultiRoutineLock
)

type MultiRoutineLock struct {
	mu    sync.Mutex
	store map[string]int64 // key is lock name,value is second timestamp
}

func NewMultiRoutineLock() *MultiRoutineLock {
	multiRoutineLockOnce.Do(func() {
		multiRoutineLock = &MultiRoutineLock{
			mu:    sync.Mutex{},
			store: make(map[string]int64),
		}
	})
	return multiRoutineLock
}
func (m *MultiRoutineLock) Lock(key string, expired time.Duration) (locked bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	value, ok := m.store[key]
	now := time.Now()
	if ok && now.Unix() < value {
		return false
	}
	expireTime := now.Add(expired).Unix()
	m.store[key] = expireTime
	return true
}

func (m *MultiRoutineLock) Unlock(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.store, key)
}
