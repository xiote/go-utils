package redisutil

import (
	"context"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"time"
)

type Mutex struct {
	redisClient *redis.Client
	lockClient  *redislock.Client
	lock        *redislock.Lock
	key         string
}

func NewMutex(Addr string, key string) (mutex *Mutex) {
	mutex = &Mutex{}
	mutex.key = key
	mutex.redisClient = redis.NewClient(&redis.Options{
		Network: "tcp",
		// Addr:    "169.254.178.182:6379",
		Addr: Addr,
	})
	mutex.lockClient = redislock.New(mutex.redisClient)
	return
}

func (m *Mutex) Close() {
	m.redisClient.Close()
}

func (m *Mutex) Lock(ctx context.Context) (err error) {

	// Retry every 100ms, for up-to 100x
	backoff := redislock.LimitRetry(redislock.LinearBackoff(100*time.Millisecond), 300)

	// Obtain lock with retry
	m.lock, err = m.lockClient.Obtain(ctx, m.key, 10*time.Second, &redislock.Options{
		RetryStrategy: backoff,
	})

	return
}
func (m *Mutex) Unlock(ctx context.Context) (err error) {
	m.lock.Release(ctx)
	return
}
