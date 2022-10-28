package math

type Integer interface {
	int | int64 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8
}

type Float interface {
	Integer | float64 | float32
}

type Complex interface {
	Float | complex128 | complex64
}

// Max 返回多个数中最大的那个
func Max[T Float](nums ...T) T {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}

// Min 返回多个数中最小的那个
func Min[T Float](nums ...T) T {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < m {
			m = nums[i]
		}
	}
	return m
}
