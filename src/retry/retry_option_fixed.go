package retry

import (
	"github.com/gtxistxgao/go-retry/src/timeutil"
	"time"
)

// FixedRetryOption is the setting for fixed retry
type FixedRetryOption struct {
	Base             time.Duration
	MaxRetryAttempts uint
	Jitter           time.Duration
	LogError         bool
}

func (e *FixedRetryOption) getWaitTime(retryCount uint) time.Duration {
	return e.Base + timeutil.RandomDuration(e.Jitter)
}

func (e *FixedRetryOption) maxRetryAttempts() uint {
	max := e.MaxRetryAttempts
	if max < 0 {
		return 0
	}

	return max
}

func (e *FixedRetryOption) logError() bool {
	return e.LogError
}
