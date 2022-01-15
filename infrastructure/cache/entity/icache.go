package entity

import "time"

type ICache interface {
	Set(key string, value []byte, expired time.Duration)
	Get(key string) []byte
}
