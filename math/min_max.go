package math

import (
	"cmp"
)

// Max 返回多个数中最大的那个
func Max[T cmp.Ordered](nums ...T) T {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}

// Min 返回多个数中最小的那个
func Min[T cmp.Ordered](nums ...T) T {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < m {
			m = nums[i]
		}
	}
	return m
}
