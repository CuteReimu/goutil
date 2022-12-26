package util

// PriorityQueue 优先队列
type PriorityQueue[T comparable] struct {
	queue      []T
	Comparator func(o1, o2 T) int
}

// NewPriorityQueue 用给定的初始值和比较函数新建优先队列
func NewPriorityQueue[T comparable](values []T, comparator func(o1, o2 T) int) *PriorityQueue[T] {
	q := &PriorityQueue[T]{queue: values, Comparator: comparator}
	q.heapify()
	return q
}

func (q *PriorityQueue[T]) Add(e T) {
	q.queue = append(q.queue, e)
	q.siftUp(len(q.queue)-1, e)
}

func (q *PriorityQueue[T]) Peek() T {
	return q.queue[0]
}

func (q *PriorityQueue[T]) indexOf(o T) int {
	es := q.queue
	for i, n := 0, len(q.queue); i < n; i++ {
		if q.Comparator(o, es[i]) == 0 {
			return i
		}
	}
	return -1
}

func (q *PriorityQueue[T]) Remove(o T) bool {
	i := q.indexOf(o)
	if i == -1 {
		return false
	} else {
		q.removeAt(i)
		return true
	}
}

func (q *PriorityQueue[T]) Contains(o T) bool {
	return q.indexOf(o) >= 0
}

// ToSlice 返回队列中的所有元素。
//
// 注意，ToSlice并不一定按照大小顺序返回
func (q *PriorityQueue[T]) ToSlice(in []T) []T {
	l := len(q.queue)
	if len(in) < l {
		in = make([]T, l)
	}
	copy(in, q.queue)
	return in
}

// Foreach 循环对队列中的每个元素调用函数f。对于f的返回值，若返回true表示继续循环，返回false表示跳出循环。
//
// 注意，Foreach并不一定按照大小顺序执行
func (q *PriorityQueue[T]) Foreach(f func(e T) bool) {
	es := q.queue
	for i, n := 0, len(q.queue); i < n && f(es[i]); i++ {
	}
}

func (q *PriorityQueue[T]) Len() int {
	return len(q.queue)
}

func (q *PriorityQueue[T]) Clear() {
	q.queue = nil
}

func (q *PriorityQueue[T]) Poll() T {
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

func (q *PriorityQueue[T]) removeAt(i int) T {
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

func (q *PriorityQueue[T]) siftUp(k int, e T) {
	for k > 0 {
		parent := int((uint(k) - 1) >> 1)
		e1 := q.queue[parent]
		if q.Comparator(e, e1) >= 0 {
			break
		}
		q.queue[k] = e1
		k = parent
	}
	q.queue[k] = e
}

func (q *PriorityQueue[T]) siftDown(k int, x T) {
	es := q.queue
	n := len(es)
	half := int(uint(n) >> 1)
	for k < half {
		child := (k << 1) + 1
		c := es[child]
		right := child + 1
		if right < n && q.Comparator(c, es[right]) > 0 {
			child = right
			c = es[child]
		}
		if q.Comparator(x, c) <= 0 {
			break
		}
		es[k] = c
		k = child
	}
	es[k] = x
}

func (q *PriorityQueue[T]) heapify() {
	es := q.queue
	n := len(es)
	if n > 1 {
		for i := int((uint(n) >> 1) - 1); i >= 0; i-- {
			q.siftDown(i, es[i])
		}
	}
}
