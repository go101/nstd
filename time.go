package nstd

import (
	"log"
	"time"
)

// ElapsedTimeLogFunc returns a function which prints the duration elapsed
// since the time the function is invoked.
//
// Use example:
//
//	logElapsedTime := nstd.ElapsedTimeLogFunc("")
//	... // do task 1
//	logElapsedTime("task 1:")
//	... // do task 2
//	_ = showElapseLog && logElapsedTime("task 2:")
func ElapsedTimeLogFunc(commonPrefix string) func(prefix string) bool {
	var x string
	if commonPrefix != "" {
		x = " "
	}
	var start = time.Now()
	return func(prefix string) bool {
		var y string
		if commonPrefix != "" {
			y = " "
		}
		log.Printf("%s%s%s%s%s", commonPrefix, x, prefix, y, time.Since(start))
		return true
	}
}

// ScaleDuration scales a time.Duration.
//
// See: https://github.com/golang/go/issues/75265
func ScaleDuration(d time.Duration, s float64) time.Duration {
	return time.Duration(float64(d) * s)
}
