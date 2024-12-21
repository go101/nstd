//go:build go1.23

package nstd

import (
	"os"
)

// CopyDir copies a directory.
func CopyDir(dest, src string) error {
	return os.CopyFS(dest, os.DirFS(src))
}
