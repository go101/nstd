package nstd

import (
	"fmt"
)

// Printfln(format, a...) is a combintion of fmt.Printf(format, a...)
// and fmt.Println().
func Printfln(format string, a ...any) (n int, err error) {
	n, err = fmt.Printf(format, a...)
	if err != nil {
		return
	}

	m, err := fmt.Println()
	n += m
	return
}
