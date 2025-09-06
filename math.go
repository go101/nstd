package nstd

// Scale scales a Real value to (x * s).
//
// An example: Scale(time.Duration, 0.1)
//
// See: https://github.com/golang/go/issues/75265
func Scale[R, S Real](x R, s S) R {
	return R(float64(x) * float64(s))
}

// Scale scales a Real value to (x * s1 / s2).
func Scale2[R, S Real](x R, s1, s2 S) R {
	return R(float64(x) * float64(s1) / float64(s2))
}
