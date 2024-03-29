package slices

import (
	"fmt"
)

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Seq 生成连续多个数组成的 slice
func Seq[T integer](from, to T, span ...T) []T {
	if from == to {
		return nil
	}
	var s T = 1
	if len(span) >= 1 {
		s = span[0]
		if s == 0 {
			panic("span shouldn't be 0")
		}
	}
	if (s > 0) == (from > to) {
		panic(fmt.Sprint("illegal param: ", from, to, s))
	}
	var arr []T
	if to > from {
		arr = make([]T, 0, (to-from-1)/s+1)
	} else {
		arr = make([]T, 0, (from-to-1)/s+1)
	}
	for i := from; i < to; i += s {
		arr = append(arr, i)
	}
	return arr
}
