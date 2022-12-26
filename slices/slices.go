package slices

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math/rand"
	"sort"
)

// Contains 判断一个 slice 中是否包含某个元素
func Contains[T comparable](arr []T, e T) bool {
	for _, e1 := range arr {
		if e1 == e {
			return true
		}
	}
	return false
}

// Equals 当且仅当两个 slice 长度相同且包含的元素完全相同时返回 ture ，否则返回 false
func Equals[T comparable](arr1, arr2 []T) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

// CopyOf 复制指定的 slice ，根据 newLength ，如果有必要，则截取或者在后面填默认的零值
func CopyOf[T any](original []T, newLength int) []T {
	newSlice := make([]T, newLength)
	copy(newSlice, original)
	return newSlice
}

// CopyOfRange 将指定 slice 从 from 到 to 下标复制出来，如果 to 比原 slice 的长度更长，则在后面补默认的零值
func CopyOfRange[T any](original []T, from, to int) []T {
	newLength := to - from
	if newLength < 0 {
		panic(fmt.Sprint(from, " > ", to))
	}
	newSlice := make([]T, newLength)
	copy(newSlice, original[from:])
	return newSlice
}

// ShuffleN 打乱一个 slice ，但后续只会用前n个值，因此做了一些优化
func ShuffleN[T any](rand *rand.Rand, arr []T, n int) {
	if n < 0 {
		panic("invalid argument to Shuffle")
	}
	for i := 0; i < n; i++ {
		j := i + rand.Intn(len(arr)-i)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Any 只要 0..n （含0，不含n）中的任一元素满足给定的 func ，就返回 true ，否则返回 false
func Any(n int, f func(i int) bool) bool {
	for i := 0; i < n; i++ {
		if f(i) {
			return true
		}
	}
	return false
}

// All 仅当 0..n （含0，不含n）中的所有元素都满足给定的 func ， 才返回 true ，否则返回 false
func All(n int, f func(i int) bool) bool {
	for i := 0; i < n; i++ {
		if !f(i) {
			return false
		}
	}
	return true
}

type anySlice[T any] struct {
	elems []T
	less  func(a, b T) bool
}

func (m *anySlice[T]) Len() int {
	return len(m.elems)
}

func (m *anySlice[T]) Less(i, j int) bool {
	return m.less(m.elems[i], m.elems[j])
}

func (m *anySlice[T]) Swap(i, j int) {
	m.elems[i], m.elems[j] = m.elems[j], m.elems[i]
}

// Sort 排序，改变原 slice
func Sort[T any](arr []T, less func(a, b T) bool) {
	sort.Sort(&anySlice[T]{
		elems: arr,
		less:  less,
	})
}

// Usort 排序并去重，返回新的 slice
func Usort[T any](arr []T, lessThan func(a, b T) bool) []T {
	if arr == nil {
		return nil
	}
	if len(arr) <= 1 {
		return CopyOf(arr, len(arr))
	}
	var newSlice []T
	for _, e := range arr {
		i := sort.Search(len(newSlice), func(i int) bool { return !lessThan(newSlice[i], e) })
		if i >= len(newSlice) || lessThan(e, newSlice[i]) || lessThan(newSlice[i], e) {
			if cap(newSlice) >= len(newSlice)+1 {
				newSlice = append(newSlice, newSlice[len(newSlice)-1])
				for j := len(newSlice) - 1; j > i; j-- {
					newSlice[j] = newSlice[j-1]
				}
				newSlice[i] = e
			}
			newSlice = append(append(newSlice[:i:i], e), newSlice[i:]...)
		}
	}
	return newSlice
}

// Uniq 去重，返回新的 slice
func Uniq[T comparable](arr []T) []T {
	if arr == nil {
		return nil
	}
	if len(arr) <= 1 {
		return CopyOf(arr, len(arr))
	}
	newSlice := make([]T, 0, len(arr))
	m := make(map[T]struct{}, len(arr))
	for _, e := range arr {
		if _, ok := m[e]; !ok {
			m[e] = struct{}{}
			newSlice = append(newSlice, e)
		}
	}
	return newSlice
}

// Map 对 n..0 （含0，不含n）根据给定函数将每个元素映射并筛选后得到一个新的 slice
func Map[T any](n int, f func(i int) (T, bool)) (ret []T) {
	for i := 0; i < n; i++ {
		if e1, ok := f(i); ok {
			ret = append(ret, e1)
		}
	}
	return
}

// Reverse 将给定 slice 反向
func Reverse[T any](arr []T) {
	half := len(arr) / 2
	for i := 0; i < half; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}

// Fold 对 0..n （含0，不含n）的每个元素依次调用给定函数，得到一个最终值
func Fold[T any](n int, f func(i int, acc T) T, initial T) T {
	for i := 0; i < n; i++ {
		initial = f(i, initial)
	}
	return initial
}

// FoldReverse 对 n..0 （含0，不含n） 的每个元素依次调用给定函数，得到一个最终值
func FoldReverse[T any](n int, f func(i int, acc T) T, initial T) T {
	for i := n - 1; i >= 0; i-- {
		initial = f(i, initial)
	}
	return initial
}

// Duplicate 生成一个元素全为 e 的长度为 count 的 slice
func Duplicate[T any](count int, e T) []T {
	ret := make([]T, count)
	for i := 0; i < count; i++ {
		ret[i] = e
	}
	return ret
}

// Sum 求和，小心溢出
func Sum[T interface {
	constraints.Integer | constraints.Float | constraints.Complex
}](arr []T) (sum T) {
	for _, e := range arr {
		sum += e
	}
	return
}
