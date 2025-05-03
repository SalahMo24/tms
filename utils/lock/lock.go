package locking

import (
	"sync"
	"time"
)

type InMemoryLock struct {
	locks    sync.Map
	timeouts sync.Map
}

func NewInMemoryLock() *InMemoryLock {
	return &InMemoryLock{}
}

func (l *InMemoryLock) AcquireLock(key string, ttl time.Duration) (bool, error) {
	// Check if lock exists and isn't expired
	if expiry, ok := l.timeouts.Load(key); ok {
		if time.Now().Before(expiry.(time.Time)) {
			return false, nil // Lock still active
		}
	}

	// Try to acquire lock
	_, loaded := l.locks.LoadOrStore(key, struct{}{})
	if !loaded {
		// Set expiration time
		l.timeouts.Store(key, time.Now().Add(ttl))
		return true, nil
	}
	return false, nil
}

func (l *InMemoryLock) ReleaseLock(key string) error {
	l.locks.Delete(key)
	l.timeouts.Delete(key)
	return nil
}
