//go:build go1.23

package nstd

import (
	"os"
)

func CopyDir(dest, src string) error {
	return os.CopyFS(dest, os.DirFS(src))
}
