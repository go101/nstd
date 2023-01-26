package conv

// B2i converts a bool value to int (true -> 1, false -> 0).
//
// ToDo: move to the top "nstd" package?
func B2i(x bool) int {
	if x {
		return 1
	}
	return 0
}
