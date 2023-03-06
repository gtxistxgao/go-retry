package retry

import (
	"github.com/gtxistxgao/go-retry/src/mathutil"
	"github.com/gtxistxgao/go-retry/src/timeutil"
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
	return e.Base + e.Step*time.Duration(mathutil.ExponentBase2(retryCount)) + timeutil.RandomDuration(e.Jitter)
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
