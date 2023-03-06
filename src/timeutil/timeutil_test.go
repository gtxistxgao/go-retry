package timeutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_RandomDuration(t *testing.T) {
	d := RandomDuration(time.Second)
	assert.True(t, d >= 0)
	assert.True(t, d < time.Second)

	d = RandomDuration(time.Millisecond)
	assert.True(t, d >= 0)
	assert.True(t, d < time.Millisecond)
}
