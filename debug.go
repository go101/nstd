package nstd

import (
	"log"
)

// Debug calls [log.Print] and returns true,
// so that it can be used in
//
//	_ = DebugMode && nstd.Debug(...)
func Debug(v ...any) bool {
	log.Print(v...)
	return true
}

// Debugf calls [log.Printf] and return true,
// so that it can be used in
//
//	_ = DebugMode && nstd.Debugf(...)
func Debugf(format string, v ...any) bool {
	log.Printf(format, v...)
	return true
}
