package timeutil

import (
	"math/rand"
	"time"
)

// RandomDuration returns a random duration between [0, d).
func RandomDuration(d time.Duration) time.Duration {
	if d == 0 {
		return 0
	}

	return time.Duration(rand.Int63n(int64(d)))
}
