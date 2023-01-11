package slices

import "testing"

func TestSeqInt(t *testing.T) {
	arr := Seq[int](1, 10, 2)
	if !Equals(arr, []int{1, 3, 5, 7, 9}) {
		t.Log("incorrect Seq: ", arr, []int{1, 3, 5, 7, 9})
		t.Fail()
	}
}

func TestSeqInt32(t *testing.T) {
	arr := Seq[int32](1, 5)
	if !Equals(arr, []int32{1, 2, 3, 4}) {
		t.Log("incorrect Seq: ", arr, []int32{1, 2, 3, 4})
		t.Fail()
	}
}

func TestSeqUintptr(t *testing.T) {
	arr := Seq[uintptr](1, 10)
	if !Equals(arr, []uintptr{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Log("incorrect Seq: ", arr, []uintptr{1, 2, 3, 4, 5, 6, 7, 8, 9})
		t.Fail()
	}
}
