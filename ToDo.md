


## try to merge all packages into the "nstd" package.

```go
package nstd

type strPkg struct{
	_ [0]*strPkg
}

var String strPkg

func (strPkg) TrimLeadingSpaces(s string) string {...}
func (strPkg) TrimTrailingSpaces(s string) string {...}
// or
func (strPkg) TrimStartingSpaces(s string) string {...}
func (strPkg) TrimEndingSpaces(s string) string {...}


```

```go
import "nstd"

void main() {
	s = nstd.String.TrimStartingSpaces(s)
}
```


