package goutil

import (
	"sync"
	"sync/atomic"
)

type node[T any] struct {
	item T
	next *node[T]
}

// BlockingQueue 线程安全的队列，容量无限，用链表实现
type BlockingQueue[T any] struct {
	count             atomic.Int64
	head, last        *node[T]
	takeLock, putLock sync.Mutex
	notEmpty          *sync.Cond
}

// NewBlockingQueue 新建一个线程安全的队列
func NewBlockingQueue[T any]() *BlockingQueue[T] {
	q := &BlockingQueue[T]{
		head: new(node[T]),
	}
	q.last = q.head
	q.notEmpty = sync.NewCond(&q.takeLock)
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

// Put 向队尾插入一个元素。因为容量无限，所以不会阻塞
func (q *BlockingQueue[T]) Put(e T) {
	n := &node[T]{item: e}
	q.putLock.Lock()
	q.enqueue(n)
	c := q.count.Add(1)
	q.putLock.Unlock()
	if c == 1 {
		q.notEmpty.Signal()
	}
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
