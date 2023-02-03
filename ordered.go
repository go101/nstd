package nstd

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
type Ordered interface {
	Integer | Float | ~string
}

func Compare[T Ordered](x, y T) int {
	if x == y {
		return 0
	}
	if x < y {
		return -1
	}
	return 1
}

func Min[T Ordered](x, y T) T {
	if x <= y {
		return x
	}
	return y
}

func MinN[T Ordered](vs ...T) T {
	if len(vs) == 0 {
		panic("no values")
	}
	v := vs[0]
	for _, x := range vs[1:] {
		if v > x {
			v = x
		}
	}
	return v
}

func Max[T Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func MaxN[T Ordered](vs ...T) T {
	if len(vs) == 0 {
		panic("no values")
	}
	v := vs[0]
	for _, x := range vs[1:] {
		if v < x {
			v = x
		}
	}
	return v
}