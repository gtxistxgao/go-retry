package go_retry

import "time"

// retryOption provides an interface so we can provide retryHandler different retry strategies.
type retryOption interface {
	// maxRetryAttempts returns how many times we can retry. This excludes the 1st attempt
	// So the task can be at maximum maxRetryAttempts() + 1 times.
	maxRetryAttempts() uint
	// getWaitTime returns how long should the retryHandler wait before the next retry
	getWaitTime(retryCount uint) time.Duration
	// logError returns true if we need to log the errors during retrying
	logError() bool
}
