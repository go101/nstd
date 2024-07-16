package nstd

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
type Ordered interface {
	Real | ~string
}

func minOfTwo[T Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Clamp clamps an ordered value within a range.
// Both min and max are inclusive.
// If v is NaN, then NaN is returned.
//
// See: https://github.com/golang/go/issues/58146
func Clamp[T Ordered](v, min, max T) T {
	if min > max {
		Panicf("min (%v) > max (%v)!", min, max)
	}

	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
