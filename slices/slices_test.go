package slices

import (
	"math/rand"
	"testing"
	"time"
)

func TestContains(t *testing.T) {
	arr := []int{1, 3, 4, 2, 5}
	if !Contains(arr, 5) {
		t.Log("incorrect Contains: ", arr, 5)
		t.Fail()
	}
	if Contains(arr, 9) {
		t.Log("incorrect Contains: ", arr, 9)
		t.Fail()
	}
}

func TestEquals(t *testing.T) {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{1, 2, 3, 4, 5}
	if Equals(arr1, arr2) {
		t.Log("incorrect Equals: ", arr1, arr2)
		t.Fail()
	}
	if !Equals(arr1, arr1) {
		t.Log("incorrect Equals: ", arr1, arr1)
		t.Fail()
	}
	if !Equals(arr1, arr2[:len(arr1)]) {
		t.Log("incorrect Equals: ", arr1, arr2[:len(arr1)])
		t.Fail()
	}
}

func TestCopyOf(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !Equals(CopyOf(arr, 3), []int{1, 2, 3}) {
		t.Log("incorrect CopyOf: ", CopyOf(arr, 3), []int{1, 2, 3})
		t.Fail()
	}
	if !Equals(CopyOf(arr, 7), []int{1, 2, 3, 4, 5, 0, 0}) {
		t.Log("incorrect CopyOf: ", CopyOf(arr, 3), []int{1, 2, 3, 4, 5, 0, 0})
		t.Fail()
	}
}

func TestCopyOfRange(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !Equals(CopyOfRange(arr, 1, 3), []int{2, 3}) {
		t.Log("incorrect CopyOfRange: ", CopyOfRange(arr, 1, 3), []int{2, 3})
		t.Fail()
	}
	if !Equals(CopyOfRange(arr, 1, 7), []int{2, 3, 4, 5, 0, 0}) {
		t.Log("incorrect CopyOfRange: ", CopyOfRange(arr, 1, 7), []int{2, 3, 4, 5, 0, 0})
		t.Fail()
	}
}

func TestShuffleN(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	for i := 0; i < 100; i++ {
		arr := []int{1, 2, 3}
		ShuffleN(r, arr, 1)
		if !Equals(arr, []int{1, 2, 3}) && !Equals(arr, []int{2, 1, 3}) && !Equals(arr, []int{3, 2, 1}) {
			t.Log("incorrect ShuffleN: ", []int{1, 2, 3}, arr)
			t.FailNow()
		}
	}
}

func TestAny(t *testing.T) {
	arr := []int{1, 2, 3}
	if !Any(len(arr), func(i int) bool { return arr[i] <= 1 }) {
		t.Log("incorrect Any: ", arr)
		t.Fail()
	}
	if Any(len(arr), func(i int) bool { return arr[i] < 1 }) {
		t.Log("incorrect Any: ", arr)
		t.Fail()
	}
}

func TestAll(t *testing.T) {
	arr := []int{1, 2, 3}
	if !All(len(arr), func(i int) bool { return arr[i] <= 3 }) {
		t.Log("incorrect All: ", arr)
		t.Fail()
	}
	if All(len(arr), func(i int) bool { return arr[i] < 3 }) {
		t.Log("incorrect All: ", arr)
		t.Fail()
	}
}

func TestSort(t *testing.T) {
	arr := []int{1, 2, 3, 2, 1, 2, 3}
	Sort(arr, func(a, b int) bool { return a > b })
	if !Equals(arr, []int{3, 3, 2, 2, 2, 1, 1}) {
		t.Log("incorrect Sort: ", arr, []int{3, 3, 2, 2, 2, 1, 1})
		t.Fail()
	}
}

func TestUsort(t *testing.T) {
	arr := []int{1, 2, 3, 2, 1, 2, 3}
	arr = Usort(arr, func(a, b int) bool { return a > b })
	if !Equals(arr, []int{3, 2, 1}) {
		t.Log("incorrect Sort: ", arr, []int{3, 2, 1})
		t.Fail()
	}
	arr = []int{1, 2, 3, 2, 1, 2, 3}
	arr = Usort(arr, func(a, b int) bool { return a < b })
	if !Equals(arr, []int{1, 2, 3}) {
		t.Log("incorrect Sort: ", arr, []int{1, 2, 3})
		t.Fail()
	}
}

func TestUniq(t *testing.T) {
	arr := Uniq([]int{1, 2, 3, 2, 1, 2, 3})
	Sort(arr, func(a, b int) bool { return a < b })
	if !Equals(arr, []int{1, 2, 3}) {
		t.Log("incorrect Sort: ", arr, []int{3, 2, 1})
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}
	arr = Map(len(arr), func(i int) (int, bool) { return arr[i] * 2, arr[i]%2 == 1 })
	if !Equals(arr, []int{2, 6}) {
		t.Log("incorrect Filter: ", arr, []int{2, 6})
		t.Fail()
	}
}

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3}
	Reverse(arr)
	if !Equals(arr, []int{3, 2, 1}) {
		t.Log("incorrect Map: ", arr, []int{3, 2, 1})
		t.Fail()
	}
	arr = []int{1, 2, 3, 4}
	Reverse(arr)
	if !Equals(arr, []int{4, 3, 2, 1}) {
		t.Log("incorrect Map: ", arr, []int{4, 3, 2, 1})
		t.Fail()
	}
}

func TestFold(t *testing.T) {
	sum := Fold(100, func(i int, acc int) int { return acc + i + 1 }, 0)
	if sum != SumInt(SeqInt(1, 101)) {
		t.Log("incorrect Fold: ", sum, SumInt(SeqInt(1, 101)))
		t.Fail()
	}
}

func TestFoldReverse(t *testing.T) {
	sum := FoldReverse(100, func(e int, acc int) int { return acc + e + 1 }, 0)
	if sum != SumInt(SeqInt(1, 101)) {
		t.Log("incorrect Fold: ", sum, SumInt(SeqInt(1, 101)))
		t.Fail()
	}
}

func TestDuplicate(t *testing.T) {
	arr := Duplicate(100, 100)
	if SumInt(arr) != 10000 {
		t.Log("incorrect Fold: ", SumInt(arr), 10000)
		t.Fail()
	}
}
