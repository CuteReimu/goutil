package goutil

import (
	"math"
	"sync"
	"sync/atomic"
)

type node[T any] struct {
	item T
	next *node[T]
}

// BlockingQueue 线程安全的队列，用链表实现
type BlockingQueue[T any] struct {
	capacity          int64
	count             atomic.Int64
	head, last        *node[T]
	takeLock, putLock sync.Mutex
	notEmpty, notFull *sync.Cond
}

// NewBlockingQueue 新建一个线程安全的队列， capacity ≤ 0 表示无限（最多 math.MaxInt64 ）
func NewBlockingQueue[T any](capacity int64) *BlockingQueue[T] {
	if capacity <= 0 {
		capacity = math.MaxInt64
	}
	q := &BlockingQueue[T]{
		capacity: capacity,
		head:     new(node[T]),
	}
	q.last = q.head
	q.notEmpty = sync.NewCond(&q.takeLock)
	q.notFull = sync.NewCond(&q.putLock)
	return q
}

// enqueue Links node at end of queue. 调用前必须先获取putLock
func (q *BlockingQueue[T]) enqueue(node *node[T]) {
	q.last.next = node
	q.last = node
}

// dequeue Removes a node from head of queue. 调用前必须先获取takeLock
func (q *BlockingQueue[T]) dequeue() T {
	h := q.head
	first := h.next
	h.next = h // help GC
	q.head = first
	return first.item
}

func (q *BlockingQueue[T]) Len() int64 {
	return q.count.Load()
}

func (q *BlockingQueue[T]) Cap() int64 {
	return q.capacity
}

func (q *BlockingQueue[T]) RemainingCapacity() int64 {
	return q.capacity - q.count.Load()
}

// Put 向队尾插入一个元素，如果满了则阻塞
func (q *BlockingQueue[T]) Put(e T) {
	n := &node[T]{item: e}
	q.putLock.Lock()
	for q.count.Load() == q.capacity {
		q.notFull.Wait()
	}
	q.enqueue(n)
	c := q.count.Add(1)
	if c < q.capacity {
		q.notFull.Signal()
	}
	q.putLock.Unlock()
	if c == 1 {
		q.notEmpty.Signal()
	}
}

// Offer 向队尾插入一个元素，不阻塞，如果满了就返回 false
func (q *BlockingQueue[T]) Offer(e T) bool {
	if q.count.Load() == q.capacity {
		return false
	}
	n := &node[T]{item: e}
	q.putLock.Lock()
	if q.count.Load() == q.capacity {
		q.putLock.Unlock()
		return false
	}
	q.enqueue(n)
	c := q.count.Add(1)
	if c < q.capacity {
		q.notFull.Signal()
	}
	q.putLock.Unlock()
	if c == 1 {
		q.notEmpty.Signal()
	}
	return true
}

// Take 从队首获取一个元素，如果队列为空则阻塞
func (q *BlockingQueue[T]) Take() T {
	q.takeLock.Lock()
	for q.count.Load() == 0 {
		q.notEmpty.Wait()
	}
	x := q.dequeue()
	c := q.count.Add(-1)
	if c > 0 {
		q.notEmpty.Signal()
	}
	q.takeLock.Unlock()
	if c == q.capacity-1 {
		q.notFull.Signal()
	}
	return x
}

// Poll 从队首获取一个元素，不阻塞
func (q *BlockingQueue[T]) Poll() (e T, ok bool) {
	if q.count.Load() == 0 {
		return
	}
	q.takeLock.Lock()
	if q.count.Load() == 0 {
		q.takeLock.Unlock()
		return
	}
	e, ok = q.dequeue(), true
	c := q.count.Add(-1)
	if c > 0 {
		q.notEmpty.Signal()
	}
	q.takeLock.Unlock()
	if c == q.capacity-1 {
		q.notFull.Signal()
	}
	return
}

// Peek 返回队首的元素，不取出
func (q *BlockingQueue[T]) Peek() (e T, ok bool) {
	if q.count.Load() == 0 {
		return
	}
	q.takeLock.Lock()
	if q.count.Load() > 0 {
		e, ok = q.head.next.item, true
	}
	q.takeLock.Unlock()
	return
}
