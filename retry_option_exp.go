package go_retry

import (
	"math/rand"
	"time"
)

// ExponentialRetryOption is the setting for exponential retry
type ExponentialRetryOption struct {
	Base             time.Duration
	Step             time.Duration
	MaxRetryAttempts uint
	Jitter           time.Duration
	LogError         bool
}

func (e *ExponentialRetryOption) getWaitTime(retryCount uint) time.Duration {
	return e.Base + e.Step*time.Duration(ExponentBase2(retryCount)) + RandomDuration(e.Jitter)
}

func (e *ExponentialRetryOption) maxRetryAttempts() uint {
	max := e.MaxRetryAttempts
	if max < 0 {
		return 0
	}

	return max
}

func (e *ExponentialRetryOption) logError() bool {
	return e.LogError
}

// ExponentBase2 computes 2^a where a >= 0. If a is 0, the result is 0.
func ExponentBase2(a uint) uint {
	return 1 << a
}

// RandomDuration returns a random duration between [0, d).
func RandomDuration(d time.Duration) time.Duration {
	if d == 0 {
		return 0
	}

	return time.Duration(rand.Int63n(int64(d)))
}
