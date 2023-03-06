package mathutil

// ExponentBase2 computes 2^a where a >= 0. If a is 0, the result is 0.
func ExponentBase2(a uint) uint {
	return 1 << a
}
