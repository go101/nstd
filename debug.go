package nstd

import (
	"log"
	"os"
)

var debugLogger = log.New(os.Stderr, "[debug]: ", 0)

// Debug calls [log.Print] and returns true,
// so that it can be used in
//
//	_ = DebugMode && nstd.Debug(...)
func Debug(v ...any) bool {
	debugLogger.Print(v...)
	return true
}

// Debugf calls [log.Printf] and return true,
// so that it can be used in
//
//	_ = DebugMode && nstd.Debugf(...)
func Debugf(format string, v ...any) bool {
	debugLogger.Printf(format, v...)
	return true
}

// assert is used interanlly
func assert(condition bool, failMessage string, args ...any) bool {
	if !condition {
		if len(args) == 0 {
			debugLogger.Print(failMessage)
		} else {
			debugLogger.Printf(failMessage, args...)
		}
		panic("Assert fails")
	}

	return true
}
