package goutil

import (
	"golang.org/x/exp/slices"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestBlockingQueuePut(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	values := make([][]int, 100)
	for i := 0; i < len(values); i++ {
		for j := 0; j < 10000; j++ {
			values[i] = append(values[i], r.Int())
		}
	}
	var wg sync.WaitGroup
	wg.Add(len(values))
	q := NewBlockingQueue[int]()
	for i := range values {
		value := values[i]
		go func() {
			for _, v := range value {
				q.Put(v)
			}
			wg.Done()
		}()
	}
	ch := make(chan struct{})
	var values2 []int
	go func() {
		for {
			e, ok := q.Poll()
			if ok {
				values2 = append(values2, e)
				continue
			}
			select {
			case <-ch:
				return
			default:
			}
		}
	}()
	wg.Wait()
	close(ch)
	var values1 []int
	for _, v := range values {
		values1 = append(values1, v...)
	}
	slices.Sort(values1)
	slices.Sort(values2)
	if !slices.Equal(values1, values2) {
		t.Fail()
	}
}
