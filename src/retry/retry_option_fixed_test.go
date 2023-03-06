package retry

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_FixedRetryOption(t *testing.T) {
	tests := []struct {
		description string
		retryCount  uint
		base        time.Duration
		jitter      time.Duration
		logError    bool
	}{
		{
			description: "1st attempt and no jitter",
			retryCount:  0,
			base:        time.Millisecond,
			jitter:      0,
			logError:    false,
		},
		{
			description: "1st retry and no jitter",
			retryCount:  1,
			base:        time.Millisecond,
			jitter:      0,
			logError:    false,
		},
		{
			description: "2nd retry and no jitter",
			retryCount:  2,
			base:        time.Millisecond,
			jitter:      0,
			logError:    false,
		},
		{
			description: "2nd retry and jitter 1 ms",
			retryCount:  2,
			base:        time.Second,
			jitter:      time.Second,
			logError:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var maxRetryAttempts uint = 5
			option := &FixedRetryOption{
				Base:             test.base,
				MaxRetryAttempts: maxRetryAttempts,
				Jitter:           test.jitter,
				LogError:         test.logError,
			}

			assert.Equal(t, maxRetryAttempts, option.maxRetryAttempts())
			waitTime := option.getWaitTime(test.retryCount)

			rangeStarts := option.Base
			rangeEnds := option.Base + test.jitter

			if test.jitter != 0 {
				// This will fail if no jitter
				assert.True(t, waitTime-test.jitter < rangeStarts, "Wait time is "+waitTime.String()+" and range starts with "+rangeStarts.String())
				assert.True(t, waitTime+test.jitter > rangeEnds, "Wait time is "+waitTime.String()+" and range ends with "+rangeEnds.String())
			}
			assert.True(t, waitTime >= rangeStarts, "Wait time is "+waitTime.String()+" and range starts with "+rangeStarts.String())
			assert.True(t, waitTime <= rangeEnds, "Wait time is "+waitTime.String()+" and range ends with "+rangeStarts.String())
		})
	}
}
