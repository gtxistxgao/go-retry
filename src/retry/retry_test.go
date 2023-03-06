package retry

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func Test_Run(t *testing.T) {
	type attemptResult struct {
		result      string
		err         error
		shouldRetry bool
	}

	nameTag := "test"

	tests := []struct {
		description           string
		attemptResults        []*attemptResult
		cancelledAfterAttempt int
		maxRetryAttempts      uint

		expectedResult string
		expectedError  error
	}{
		{
			description: "1st attempt succeeded",
			attemptResults: []*attemptResult{
				{
					result:      "Good",
					err:         nil,
					shouldRetry: false,
				},
				// If call it again it will fail. This case case should only get the result once
				// So put a failure case after 1st attempt
				{
					result:      "",
					err:         errors.New("try again"),
					shouldRetry: true,
				},
			},
			maxRetryAttempts: 0,
			expectedResult:   "Good",
			expectedError:    nil,
		},
		{
			description: "2nd attempt can succeed but we do not allow retry maxRetryAttempts = 0",
			attemptResults: []*attemptResult{
				{
					result:      "",
					err:         errors.New("try again"),
					shouldRetry: true,
				},
				{
					result:      "Good",
					err:         nil,
					shouldRetry: false,
				},
			},
			maxRetryAttempts: 0,
			expectedResult:   "",
			expectedError:    fmt.Errorf("retry job %s error at %d attempt: %w", nameTag, 1, errors.New("try again")),
		},
		{
			description: "1st succeeded and we do not allow retry maxRetryAttempts = 0",
			attemptResults: []*attemptResult{
				{
					result:      "Good",
					err:         nil,
					shouldRetry: false,
				},
				{
					result:      "",
					err:         errors.New("try again"),
					shouldRetry: true,
				},
			},
			maxRetryAttempts: 0,
			expectedResult:   "Good",
			expectedError:    nil,
		},
		{
			description: "2nd attempt succeeded and maxRetryAttempts is 1",
			attemptResults: []*attemptResult{
				{
					result:      "",
					err:         errors.New("try again"),
					shouldRetry: true,
				},
				{
					result:      "Good",
					err:         nil,
					shouldRetry: false,
				},
			},
			maxRetryAttempts: 1,
			expectedResult:   "Good",
			expectedError:    nil,
		},
		{
			description: "2nd attempt can succeed but we cancelled the task",
			attemptResults: []*attemptResult{
				{
					result:      "bad",
					err:         errors.New("try again"),
					shouldRetry: true,
				},
				{
					result:      "Good",
					err:         nil,
					shouldRetry: false,
				},
			},
			cancelledAfterAttempt: 1,
			maxRetryAttempts:      1,
			expectedResult:        "bad",
			expectedError:         errors.New("retry job test contextCancelled after 1 attempt: retry job test error at 1 attempt: try again"),
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			attempt := 0

			retryConfig := &testRetryOption{
				maxRetryAttemptsValue: test.maxRetryAttempts,
				waitTimeValue:         0,
				logErrorValue:         false,
			}

			retryTask := func(innerCtx context.Context) (string, error, bool) {
				currentAttemptCount := attempt
				attempt++
				if test.cancelledAfterAttempt == attempt {
					cancel()
					// wait a long time so the context can be cancelled.
					retryConfig.waitTimeValue = time.Second
				}
				rst := test.attemptResults[currentAttemptCount]
				return rst.result, rst.err, rst.shouldRetry
			}

			handler := NewRetryHandler(
				nameTag,
				retryConfig,
				retryTask)

			rst, err := handler.Run(ctx)

			assert.Equal(t, test.expectedResult, rst)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

type testRetryOption struct {
	maxRetryAttemptsValue uint
	waitTimeValue         time.Duration
	logErrorValue         bool
}

func (t *testRetryOption) maxRetryAttempts() uint {
	return t.maxRetryAttemptsValue
}

func (t *testRetryOption) getWaitTime(counter uint) time.Duration {
	return t.waitTimeValue
}

func (t *testRetryOption) logError() bool {
	return t.logErrorValue
}
