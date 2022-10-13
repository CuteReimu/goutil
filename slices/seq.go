package slices

import "fmt"

// SeqInt 生成连续多个 int 组成的 slice
func SeqInt(from, to int, span ...int) []int {
	if from == to {
		return nil
	}
	s := 1
	if len(span) >= 1 {
		s = span[s]
		if s == 0 {
			panic("span shouldn't be 0")
		}
	}
	if (s > 0) == (from > to) {
		panic(fmt.Sprint("illegal param: ", from, to, s))
	}
	var arr []int
	if to > from {
		arr = make([]int, 0, (to-from-1)/s+1)
	} else {
		arr = make([]int, 0, (from-to-1)/s+1)
	}
	for i := from; i < to; i += s {
		arr = append(arr, i)
	}
	return arr
}

// SeqInt64 生成连续多个 int64 组成的 slice
func SeqInt64(from, to int64, span ...int64) []int64 {
	if from == to {
		return nil
	}
	s := int64(1)
	if len(span) >= 1 {
		s = span[s]
		if s == 0 {
			panic("span shouldn't be 0")
		}
	}
	if (s > 0) == (from > to) {
		panic(fmt.Sprint("illegal param: ", from, to, s))
	}
	var arr []int64
	if to > from {
		arr = make([]int64, 0, (to-from-1)/s+1)
	} else {
		arr = make([]int64, 0, (from-to-1)/s+1)
	}
	for i := from; i < to; i += s {
		arr = append(arr, i)
	}
	return arr
}

// SeqInt32 生成连续多个 int32 组成的 slice
func SeqInt32(from, to int32, span ...int32) []int32 {
	if from == to {
		return nil
	}
	s := int32(1)
	if len(span) >= 1 {
		s = span[s]
		if s == 0 {
			panic("span shouldn't be 0")
		}
	}
	if (s > 0) == (from > to) {
		panic(fmt.Sprint("illegal param: ", from, to, s))
	}
	var arr []int32
	if to > from {
		arr = make([]int32, 0, (to-from-1)/s+1)
	} else {
		arr = make([]int32, 0, (from-to-1)/s+1)
	}
	for i := from; i < to; i += s {
		arr = append(arr, i)
	}
	return arr
}

// SeqInt16 生成连续多个 int16 组成的 slice
func SeqInt16(from, to int16, span ...int16) []int16 {
	if from == to {
		return nil
	}
	s := int16(1)
	if len(span) >= 1 {
		s = span[s]
		if s == 0 {
			panic("span shouldn't be 0")
		}
	}
	if (s > 0) == (from > to) {
		panic(fmt.Sprint("illegal param: ", from, to, s))
	}
	var arr []int16
	if to > from {
		arr = make([]int16, 0, (to-from-1)/s+1)
	} else {
		arr = make([]int16, 0, (from-to-1)/s+1)
	}
	for i := from; i < to; i += s {
		arr = append(arr, i)
	}
	return arr
}

// SeqInt8 生成连续多个 int8 组成的 slice
func SeqInt8(from, to int8, span ...int8) []int8 {
	if from == to {
		return nil
	}
	s := int8(1)
	if len(span) >= 1 {
		s = span[s]
		if s == 0 {
			panic("span shouldn't be 0")
		}
	}
	if (s > 0) == (from > to) {
		panic(fmt.Sprint("illegal param: ", from, to, s))
	}
	var arr []int8
	if to > from {
		arr = make([]int8, 0, (to-from-1)/s+1)
	} else {
		arr = make([]int8, 0, (from-to-1)/s+1)
	}
	for i := from; i < to; i += s {
		arr = append(arr, i)
	}
	return arr
}

// SeqUint 生成连续多个 uint 组成的 slice
func SeqUint(from, to uint, span ...uint) []uint {
	if from == to {
		return nil
	}
	s := uint(1)
	if len(span) >= 1 {
		s = span[s]
		if s == 0 {
			panic("span shouldn't be 0")
		}
	}
	arr := make([]uint, 0, (to-from-1)/s+1)
	for i := from; i < to; i += s {
		arr = append(arr, i)
	}
	return arr
}

// SeqUint64 生成连续多个 uint64 组成的 slice
func SeqUint64(from, to uint64, span ...uint64) []uint64 {
	if from == to {
		return nil
	}
	s := uint64(1)
	if len(span) >= 1 {
		s = span[s]
		if s == 0 {
			panic("span shouldn't be 0")
		}
	}
	arr := make([]uint64, 0, (to-from-1)/s+1)
	for i := from; i < to; i += s {
		arr = append(arr, i)
	}
	return arr
}

// SeqUint32 生成连续多个 uint32 组成的 slice
func SeqUint32(from, to uint32, span ...uint32) []uint32 {
	if from == to {
		return nil
	}
	s := uint32(1)
	if len(span) >= 1 {
		s = span[s]
		if s == 0 {
			panic("span shouldn't be 0")
		}
	}
	arr := make([]uint32, 0, (to-from-1)/s+1)
	for i := from; i < to; i += s {
		arr = append(arr, i)
	}
	return arr
}

// SeqUint16 生成连续多个 uint16 组成的 slice
func SeqUint16(from, to uint16, span ...uint16) []uint16 {
	if from == to {
		return nil
	}
	s := uint16(1)
	if len(span) >= 1 {
		s = span[s]
		if s == 0 {
			panic("span shouldn't be 0")
		}
	}
	arr := make([]uint16, 0, (to-from-1)/s+1)
	for i := from; i < to; i += s {
		arr = append(arr, i)
	}
	return arr
}

// SeqUint8 生成连续多个 uint8 组成的 slice
func SeqUint8(from, to uint8, span ...uint8) []uint8 {
	if from == to {
		return nil
	}
	s := uint8(1)
	if len(span) >= 1 {
		s = span[s]
		if s == 0 {
			panic("span shouldn't be 0")
		}
	}
	arr := make([]uint8, 0, (to-from-1)/s+1)
	for i := from; i < to; i += s {
		arr = append(arr, i)
	}
	return arr
}
