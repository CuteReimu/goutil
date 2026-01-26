package goutil

import (
	"cmp"
	"container/heap"
	"slices"
)

// PriorityQueue 优先队列
type PriorityQueue[T any] interface {
	Add(e T)

	Peek() T

	Remove(o T) bool

	Contains(o T) bool

	// ToSlice 返回队列中的所有元素。
	//
	// 注意，ToSlice并不一定按照大小顺序返回
	ToSlice(in []T) []T

	// Foreach 循环对队列中的每个元素调用函数f。对于f的返回值，若返回true表示继续循环，返回false表示跳出循环。
	//
	// 注意，Foreach并不一定按照大小顺序执行
	Foreach(f func(e T) bool)

	Len() int

	Clear()

	Poll() T
}

// heapWithComparator is an internal heap implementation using a custom comparator
type heapWithComparator[T any] struct {
	elements   []T
	comparator func(o1, o2 T) int
}

func (h *heapWithComparator[T]) Len() int {
	return len(h.elements)
}

func (h *heapWithComparator[T]) Less(i, j int) bool {
	return h.comparator(h.elements[i], h.elements[j]) < 0
}

func (h *heapWithComparator[T]) Swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *heapWithComparator[T]) Push(x any) {
	h.elements = append(h.elements, x.(T))
}

func (h *heapWithComparator[T]) Pop() any {
	old := h.elements
	n := len(old)
	x := old[n-1]
	var zero T
	old[n-1] = zero
	h.elements = old[:n-1]
	return x
}

type priorityQueueWithComparator[T any] struct {
	heap *heapWithComparator[T]
}

// NewPriorityQueue 用给定的初始值和比较函数新建优先队列
func NewPriorityQueue[T any](values []T, comparator func(o1, o2 T) int) PriorityQueue[T] {
	h := &heapWithComparator[T]{
		elements:   values,
		comparator: comparator,
	}
	heap.Init(h)
	q := &priorityQueueWithComparator[T]{heap: h}
	return q
}

func (q *priorityQueueWithComparator[T]) Add(e T) {
	heap.Push(q.heap, e)
}

func (q *priorityQueueWithComparator[T]) Peek() T {
	return q.heap.elements[0]
}

func (q *priorityQueueWithComparator[T]) indexOf(o T) int {
	for i := 0; i < len(q.heap.elements); i++ {
		if q.heap.comparator(o, q.heap.elements[i]) == 0 {
			return i
		}
	}
	return -1
}

func (q *priorityQueueWithComparator[T]) Remove(o T) bool {
	i := q.indexOf(o)
	if i == -1 {
		return false
	}
	heap.Remove(q.heap, i)
	return true
}

func (q *priorityQueueWithComparator[T]) Contains(o T) bool {
	return q.indexOf(o) >= 0
}

func (q *priorityQueueWithComparator[T]) ToSlice(in []T) []T {
	l := len(q.heap.elements)
	if len(in) < l {
		in = slices.Clone(q.heap.elements)
	} else {
		in = in[:l]
		copy(in, q.heap.elements)
	}
	return in
}

func (q *priorityQueueWithComparator[T]) Foreach(f func(e T) bool) {
	es := slices.Clone(q.heap.elements)
	n := len(es)
	for i := 0; i < n; i++ {
		if !f(es[i]) {
			break
		}
	}
}

func (q *priorityQueueWithComparator[T]) Len() int {
	return q.heap.Len()
}

func (q *priorityQueueWithComparator[T]) Clear() {
	q.heap.elements = nil
}

func (q *priorityQueueWithComparator[T]) Poll() T {
	// Access the first element directly to preserve the previous
	// slice bounds-check panic behavior when the queue is empty.
	result := q.heap.elements[0]
	heap.Remove(q.heap, 0)
	return result
}

// defaultHeap is an internal heap implementation for ordered types
type defaultHeap[T cmp.Ordered] struct {
	elements []T
}

func (h *defaultHeap[T]) Len() int {
	return len(h.elements)
}

func (h *defaultHeap[T]) Less(i, j int) bool {
	return h.elements[i] < h.elements[j]
}

func (h *defaultHeap[T]) Swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *defaultHeap[T]) Push(x any) {
	h.elements = append(h.elements, x.(T))
}

func (h *defaultHeap[T]) Pop() any {
	old := h.elements
	n := len(old)
	x := old[n-1]
	var zero T
	old[n-1] = zero
	h.elements = old[:n-1]
	return x
}

type defaultPriorityQueue[T cmp.Ordered] struct {
	heap *defaultHeap[T]
}

// NewDefaultPriorityQueue 用给定的初始值新建优先队列
func NewDefaultPriorityQueue[T cmp.Ordered](values []T) PriorityQueue[T] {
	h := &defaultHeap[T]{elements: values}
	heap.Init(h)
	q := &defaultPriorityQueue[T]{heap: h}
	return q
}

func (q *defaultPriorityQueue[T]) Add(e T) {
	heap.Push(q.heap, e)
}

func (q *defaultPriorityQueue[T]) Peek() T {
	return q.heap.elements[0]
}

func (q *defaultPriorityQueue[T]) indexOf(o T) int {
	for i := 0; i < len(q.heap.elements); i++ {
		if o == q.heap.elements[i] {
			return i
		}
	}
	return -1
}

func (q *defaultPriorityQueue[T]) Remove(o T) bool {
	i := q.indexOf(o)
	if i == -1 {
		return false
	}
	heap.Remove(q.heap, i)
	return true
}

func (q *defaultPriorityQueue[T]) Contains(o T) bool {
	return q.indexOf(o) >= 0
}

func (q *defaultPriorityQueue[T]) ToSlice(in []T) []T {
	l := len(q.heap.elements)
	if len(in) < l {
		in = slices.Clone(q.heap.elements)
	} else {
		in = in[:l]
		copy(in, q.heap.elements)
	}
	return in
}

func (q *defaultPriorityQueue[T]) Foreach(f func(e T) bool) {
	es := slices.Clone(q.heap.elements)
	n := len(es)
	for i := 0; i < n; i++ {
		if !f(es[i]) {
			break
		}
	}
}

func (q *defaultPriorityQueue[T]) Len() int {
	return q.heap.Len()
}

func (q *defaultPriorityQueue[T]) Clear() {
	q.heap.elements = nil
}

func (q *defaultPriorityQueue[T]) Poll() T {
	// Access the first element directly to preserve the previous
	// slice bounds-check panic behavior when the queue is empty.
	result := q.heap.elements[0]
	heap.Remove(q.heap, 0)
	return result
}
