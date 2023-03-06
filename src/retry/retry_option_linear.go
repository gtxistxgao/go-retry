package retry

import (
	"github.com/gtxistxgao/go-retry/src/timeutil"
	"time"
)

// LinearRetryOption is the setting for linear retry
type LinearRetryOption struct {
	Base             time.Duration
	Step             time.Duration
	MaxRetryAttempts uint
	Jitter           time.Duration
	LogError         bool
}

func (e *LinearRetryOption) getWaitTime(retryCount uint) time.Duration {
	return e.Base + e.Step*time.Duration(retryCount) + timeutil.RandomDuration(e.Jitter)
}

func (e *LinearRetryOption) maxRetryAttempts() uint {
	max := e.MaxRetryAttempts
	if max < 0 {
		return 0
	}

	return max
}

func (e *LinearRetryOption) logError() bool {
	return e.LogError
}
