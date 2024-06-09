package nstd

import (
	"log"
)

// Different from log.Print, Log returns a true bool value,
// so that it can be used in
//
//    _ = Debug && nstd.Log(...)
func Log(v ...any) bool {
	log.Print(v...)
	return true
}

// Different from log.Printf, Logf returns a true bool value,
// so that it can be used in
//
//    _ = Debug && nstd.Logf(...)
func Logf(format string, v ...any) bool {
	log.Printf(format, v...)
	return true
}

// Different from log.Println, Logln returns a true bool value,
// so that it can be used in
//
//    _ = Debug && nstd.Logln(...)
func Logln(v ...any) bool {
	log.Println(v...)
	return true
}
