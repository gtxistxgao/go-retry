package go_retry

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	// need to randomize the seed, else everytime the server restarts we will get the same numbers generated
	rand.Seed(time.Now().UnixNano())
}

// Task is the individual task we want to retry. It should return a result with the provided type T.
// We use generic here so that this task can be apply to any task that need one result.
type Task[T any] func(ctx context.Context) (result T, err error, shouldRetry bool)

// Handler is the struct holds the properties that we need to use for retrying
type Handler[T any] struct {
	nameTag     string // do not add space in the name tag
	run         Task[T]
	retryOption retryOption
}

// NewRetryHandler will create a retry handler object and also create a metrics operation by the retry nameTag
func NewRetryHandler[T any](
	nameTag string,
	retryOption retryOption,
	run Task[T],
) *Handler[T] {
	return &Handler[T]{
		nameTag:     nameTag,
		run:         run,
		retryOption: retryOption,
	}
}

// Run is the main function to run the job
func (r *Handler[T]) Run(ctx context.Context) (T, error) {
	var result T
	var err error

	// If the task tells us don't retry. Don't retry
	var shouldRetryByTask bool
	var attemptCount uint
	// runOnce will run the job once, then do the metrics and logging based on the response
	runOnce := func() {
		result, err, shouldRetryByTask = r.run(ctx)
		attemptCount++
		if err != nil {
			if r.retryOption.logError() {
				fmt.Errorf("retry job %s error at %d attempt: %w", r.nameTag, attemptCount, err)
			}
		}
	}

	// Run the task once
	runOnce()
	retryCount := attemptCount - 1 // retryCount does not count the first attempt
	for r.shouldRetry(shouldRetryByTask, retryCount) {
		select {
		case <-ctx.Done():
			// this means the parent context is cancelled. In this case we will return the last updated result and error
			if r.retryOption.logError() {
				fmt.Printf("retry job %s contextCancelled after %d attempt: %w", r.nameTag, attemptCount)
			}
			return result, fmt.Errorf("context cancelled. %w", err)
		case <-time.After(r.retryOption.getWaitTime(retryCount)):
			// Retry it!
			runOnce()
		}
	}

	// Only the last attempt's error will be returned. If the retry succeeded, the last operation may have a nil error.
	// If we want to know the errors before the last retry. Enable the LogError setting.
	return result, err
}

func (r *Handler[T]) shouldRetry(shouldRetryByTask bool, retryCount uint) bool {
	haveNotExceededRetryCountLimit := retryCount < r.retryOption.maxRetryAttempts()
	return shouldRetryByTask && haveNotExceededRetryCountLimit
}
