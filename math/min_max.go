package math

import (
	"cmp"
)

// Max 返回多个数中最大的那个
//
// deprecated
//
// 已废弃，请直接使用 max(1, 2)
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
//
// deprecated
//
// 已废弃，请直接使用 min(1, 2)
func Min[T cmp.Ordered](nums ...T) T {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < m {
			m = nums[i]
		}
	}
	return m
}
