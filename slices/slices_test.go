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
	for i := 0; i < 100; i++ {
		arr := []int{1, 2, 3}
		ShuffleN(rand.New(rand.NewSource(time.Now().UnixMilli())), arr, 1)
		if !Equals(arr, []int{1, 2, 3}) && !Equals(arr, []int{2, 1, 3}) && !Equals(arr, []int{3, 2, 1}) {
			t.Log("incorrect ShuffleN: ", []int{1, 2, 3}, arr)
			t.FailNow()
		}
	}
}

func TestAny(t *testing.T) {
	arr := []int{1, 2, 3}
	if !Any(arr, func(i int) bool { return i <= 1 }) {
		t.Log("incorrect Any: ", arr)
		t.Fail()
	}
	if Any(arr, func(i int) bool { return i < 1 }) {
		t.Log("incorrect Any: ", arr)
		t.Fail()
	}
}

func TestAll(t *testing.T) {
	arr := []int{1, 2, 3}
	if !All(arr, func(i int) bool { return i <= 3 }) {
		t.Log("incorrect All: ", arr)
		t.Fail()
	}
	if All(arr, func(i int) bool { return i < 3 }) {
		t.Log("incorrect All: ", arr)
		t.Fail()
	}
}
