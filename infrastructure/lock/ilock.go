package lock

import "time"

type IMultiRoutineLock interface {
	Lock(key string, expired time.Duration) (locked bool)
	Unlock(key string)
}
