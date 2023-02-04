


## ByteSeq

```go
package nstd

type byteSeq[T any] struct{
	_ [0]*byteSeq[T]
}

var ByteSeq byteSeq

func (ByteSeq[T]) TrimLeadingSpaces(s T) T {...}
func (ByteSeq[T[) TrimTrailingSpaces(s T) T {...}
// or
func (ByteSeqq[T]) TrimStartingSpaces(s T) T {...}
func (ByteSeqq[T]) TrimEndingSpaces(s T) T {...}


```



