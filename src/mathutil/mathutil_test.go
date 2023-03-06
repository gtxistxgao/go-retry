package mathutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ExponentBase2(t *testing.T) {
	// 2^0 == 1
	assert.Equal(t, uint(1), ExponentBase2(0))
	// 2^2 == 4
	assert.Equal(t, uint(4), ExponentBase2(2))
	// 2^3 == 8
	assert.Equal(t, uint(8), ExponentBase2(3))
}
