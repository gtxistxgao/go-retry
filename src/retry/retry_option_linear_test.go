package retry

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_LinearRetryOption(t *testing.T) {
	tests := []struct {
		description string
		retryCount  uint
		base        time.Duration
		step        time.Duration
		jitter      time.Duration
		logError    bool

		expectedWaitTimeRange []time.Duration
	}{
		{
			description:           "1st attempt and no jitter",
			retryCount:            0,
			base:                  time.Millisecond,
			step:                  time.Millisecond,
			jitter:                0,
			logError:              false,
			expectedWaitTimeRange: []time.Duration{1 * time.Millisecond, 1 * time.Millisecond}, // 1 + 1*0 = 1
		},
		{
			description:           "1st retry and no jitter",
			retryCount:            1,
			base:                  time.Millisecond,
			step:                  time.Millisecond,
			jitter:                0,
			logError:              false,
			expectedWaitTimeRange: []time.Duration{2 * time.Millisecond, 2 * time.Millisecond}, // 1 + 1*1 = 2
		},
		{
			description:           "2nd retry and no jitter",
			retryCount:            2,
			base:                  time.Millisecond,
			step:                  time.Millisecond,
			jitter:                0,
			logError:              false,
			expectedWaitTimeRange: []time.Duration{3 * time.Millisecond, 3 * time.Millisecond}, // 1 + 2*1 = 3
		},
		{
			description:           "2nd retry and jitter 1 ms",
			retryCount:            2,
			base:                  time.Second,
			step:                  time.Second,
			jitter:                time.Second,
			logError:              false,
			expectedWaitTimeRange: []time.Duration{3 * time.Second, 4 * time.Second}, // 1 + 2*1 + jitter[0,1] = [3,4]
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var maxRetryAttempts uint = 5
			option := &LinearRetryOption{
				Base:             test.base,
				Step:             test.step,
				MaxRetryAttempts: maxRetryAttempts,
				Jitter:           test.jitter,
				LogError:         test.logError,
			}

			assert.Equal(t, maxRetryAttempts, option.maxRetryAttempts())
			waitTime := option.getWaitTime(test.retryCount)
			if test.jitter != 0 {
				// This will fail if no jitter
				assert.True(t, waitTime-test.jitter < test.expectedWaitTimeRange[0], "Wait time is "+waitTime.String()+" and range starts with "+test.expectedWaitTimeRange[0].String())
				assert.True(t, waitTime+test.jitter > test.expectedWaitTimeRange[1], "Wait time is "+waitTime.String()+" and range ends with "+test.expectedWaitTimeRange[1].String())
			}
			assert.True(t, waitTime >= test.expectedWaitTimeRange[0], "Wait time is "+waitTime.String()+" and range starts with "+test.expectedWaitTimeRange[0].String())
			assert.True(t, waitTime <= test.expectedWaitTimeRange[1], "Wait time is "+waitTime.String()+" and range ends with "+test.expectedWaitTimeRange[0].String())
		})
	}
}
