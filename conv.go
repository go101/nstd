package nstd

// Btoi converts a bool value to int (true -> 1, false -> 0).
//
// See:
// * https://github.com/golang/go/issues/64825
func Btoi[B ~bool](x B) int {
	if x {
		return 1
	}
	return 0
}
