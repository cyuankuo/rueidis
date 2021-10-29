package cache

import (
	"time"

	"github.com/rueian/rueidis/internal/proto"
)

type Cache interface {
	GetOrPrepare(key, cmd string, ttl time.Duration) (v proto.Message, ch chan struct{})
	Update(key, cmd string, value proto.Message)
	Delete(keys []proto.Message)
	DeleteAll()
}