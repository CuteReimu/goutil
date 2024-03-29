package goutil

import (
	"cmp"
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

type priorityQueueWithComparator[T any] struct {
	queue      []T
	comparator func(o1, o2 T) int
}

// NewPriorityQueue 用给定的初始值和比较函数新建优先队列
func NewPriorityQueue[T any](values []T, comparator func(o1, o2 T) int) PriorityQueue[T] {
	q := &priorityQueueWithComparator[T]{queue: values, comparator: comparator}
	q.heapify()
	return q
}

func (q *priorityQueueWithComparator[T]) Add(e T) {
	q.queue = append(q.queue, e)
	q.siftUp(len(q.queue)-1, e)
}

func (q *priorityQueueWithComparator[T]) Peek() T {
	return q.queue[0]
}

func (q *priorityQueueWithComparator[T]) indexOf(o T) int {
	es := q.queue
	for i, n := 0, len(q.queue); i < n; i++ {
		if q.comparator(o, es[i]) == 0 {
			return i
		}
	}
	return -1
}

func (q *priorityQueueWithComparator[T]) Remove(o T) bool {
	i := q.indexOf(o)
	if i == -1 {
		return false
	} else {
		q.removeAt(i)
		return true
	}
}

func (q *priorityQueueWithComparator[T]) Contains(o T) bool {
	return q.indexOf(o) >= 0
}

func (q *priorityQueueWithComparator[T]) ToSlice(in []T) []T {
	l := len(q.queue)
	if len(in) < l {
		in = slices.Clone(q.queue)
	} else {
		copy(in, q.queue)
	}
	return in
}

func (q *priorityQueueWithComparator[T]) Foreach(f func(e T) bool) {
	es := q.queue
	for i, n := 0, len(q.queue); i < n && f(es[i]); i++ {
	}
}

func (q *priorityQueueWithComparator[T]) Len() int {
	return len(q.queue)
}

func (q *priorityQueueWithComparator[T]) Clear() {
	q.queue = nil
}

func (q *priorityQueueWithComparator[T]) Poll() T {
	es := q.queue
	result := es[0]
	n := uint(len(es)) - 1
	x := es[n]
	q.queue = q.queue[:n]
	if n > 0 {
		q.siftDown(0, x)
	}
	return result
}

func (q *priorityQueueWithComparator[T]) removeAt(i int) T {
	es := q.queue
	n := uint(len(es)) - 1
	moved := es[n]
	q.queue = q.queue[:n]
	s := len(q.queue)
	if s != i {
		q.siftDown(i, moved)
		if q.comparator(es[i], moved) == 0 {
			q.siftUp(i, moved)
			if q.comparator(es[i], moved) != 0 {
				return moved
			}
		}
	}
	return moved
}

func (q *priorityQueueWithComparator[T]) siftUp(k int, e T) {
	for k > 0 {
		parent := int((uint(k) - 1) >> 1)
		e1 := q.queue[parent]
		if q.comparator(e, e1) >= 0 {
			break
		}
		q.queue[k] = e1
		k = parent
	}
	q.queue[k] = e
}

func (q *priorityQueueWithComparator[T]) siftDown(k int, x T) {
	es := q.queue
	n := len(es)
	half := int(uint(n) >> 1)
	for k < half {
		child := (k << 1) + 1
		c := es[child]
		right := child + 1
		if right < n && q.comparator(c, es[right]) > 0 {
			child = right
			c = es[child]
		}
		if q.comparator(x, c) <= 0 {
			break
		}
		es[k] = c
		k = child
	}
	es[k] = x
}

func (q *priorityQueueWithComparator[T]) heapify() {
	es := q.queue
	n := len(es)
	if n > 1 {
		for i := int((uint(n) >> 1) - 1); i >= 0; i-- {
			q.siftDown(i, es[i])
		}
	}
}

type defaultPriorityQueue[T cmp.Ordered] struct {
	queue []T
}

// NewDefaultPriorityQueue 用给定的初始值新建优先队列
func NewDefaultPriorityQueue[T cmp.Ordered](values []T) PriorityQueue[T] {
	q := &defaultPriorityQueue[T]{queue: values}
	q.heapify()
	return q
}

func (q *defaultPriorityQueue[T]) Add(e T) {
	q.queue = append(q.queue, e)
	q.siftUp(len(q.queue)-1, e)
}

func (q *defaultPriorityQueue[T]) Peek() T {
	return q.queue[0]
}

func (q *defaultPriorityQueue[T]) indexOf(o T) int {
	es := q.queue
	for i, n := 0, len(q.queue); i < n; i++ {
		if o == es[i] {
			return i
		}
	}
	return -1
}

func (q *defaultPriorityQueue[T]) Remove(o T) bool {
	i := q.indexOf(o)
	if i == -1 {
		return false
	} else {
		q.removeAt(i)
		return true
	}
}

func (q *defaultPriorityQueue[T]) Contains(o T) bool {
	return q.indexOf(o) >= 0
}

func (q *defaultPriorityQueue[T]) ToSlice(in []T) []T {
	l := len(q.queue)
	if len(in) < l {
		in = slices.Clone(q.queue)
	} else {
		copy(in, q.queue)
	}
	return in
}

func (q *defaultPriorityQueue[T]) Foreach(f func(e T) bool) {
	es := q.queue
	for i, n := 0, len(q.queue); i < n && f(es[i]); i++ {
	}
}

func (q *defaultPriorityQueue[T]) Len() int {
	return len(q.queue)
}

func (q *defaultPriorityQueue[T]) Clear() {
	q.queue = nil
}

func (q *defaultPriorityQueue[T]) Poll() T {
	es := q.queue
	result := es[0]
	n := uint(len(es)) - 1
	x := es[n]
	q.queue = q.queue[:n]
	if n > 0 {
		q.siftDown(0, x)
	}
	return result
}

func (q *defaultPriorityQueue[T]) removeAt(i int) T {
	es := q.queue
	n := uint(len(es)) - 1
	moved := es[n]
	q.queue = q.queue[:n]
	s := len(q.queue)
	if s != i {
		q.siftDown(i, moved)
		if es[i] == moved {
			q.siftUp(i, moved)
			if es[i] != moved {
				return moved
			}
		}
	}
	return moved
}

func (q *defaultPriorityQueue[T]) siftUp(k int, e T) {
	for k > 0 {
		parent := int((uint(k) - 1) >> 1)
		e1 := q.queue[parent]
		if e >= e1 {
			break
		}
		q.queue[k] = e1
		k = parent
	}
	q.queue[k] = e
}

func (q *defaultPriorityQueue[T]) siftDown(k int, x T) {
	es := q.queue
	n := len(es)
	half := int(uint(n) >> 1)
	for k < half {
		child := (k << 1) + 1
		c := es[child]
		right := child + 1
		if right < n && c > es[right] {
			child = right
			c = es[child]
		}
		if x <= c {
			break
		}
		es[k] = c
		k = child
	}
	es[k] = x
}

func (q *defaultPriorityQueue[T]) heapify() {
	es := q.queue
	n := len(es)
	if n > 1 {
		for i := int((uint(n) >> 1) - 1); i >= 0; i-- {
			q.siftDown(i, es[i])
		}
	}
}
