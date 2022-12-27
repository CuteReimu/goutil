package util

import (
	"fmt"
	"golang.org/x/exp/slices"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestCreatorAndPoll(t *testing.T) {
	q := NewPriorityQueue([]int{4, 6, 3, 5, 9}, func(o1, o2 int) int { return o1 - o2 })
	var ss []string
	for q.Len() > 0 {
		ss = append(ss, strconv.Itoa(q.Poll()))
	}
	s := fmt.Sprint(strings.Join(ss, " "))
	if s != "3 4 5 6 9" {
		t.Log("incorrect Foreach: ", s, "| 3 4 5 6 9")
		t.Fail()
	}
}

func TestCreatorAndPoll2(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	for i := 0; i < 100; i++ {
		s := make([]int, 0, 100)
		for j := 0; j < 100; j++ {
			s = append(s, r.Intn(99999))
		}
		q := NewPriorityQueue(append([]int(nil), s...), func(o1, o2 int) int { return o1 - o2 })
		s1 := make([]int, 0, 100)
		for q.Len() > 0 {
			s1 = append(s1, q.Poll())
		}
		s2 := append([]int(nil), s...)
		slices.Sort(s2)
		if !slices.Equal(s1, s2) {
			t.Log("incorrect order: ", s)
			t.Fail()
		}
	}
}

func TestAddAndPoll(t *testing.T) {
	q := &PriorityQueue[int]{
		Comparator: func(o1, o2 int) int { return o1 - o2 },
	}
	q.Add(4)
	q.Add(6)
	q.Add(3)
	q.Add(5)
	q.Add(9)
	var ss []string
	for q.Len() > 0 {
		ss = append(ss, strconv.Itoa(q.Poll()))
	}
	s := fmt.Sprint(strings.Join(ss, " "))
	if s != "3 4 5 6 9" {
		t.Log("incorrect Foreach: ", s, "| 3 4 5 6 9")
		t.Fail()
	}
}

func TestAddAndPoll2(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	for i := 0; i < 100; i++ {
		q := &PriorityQueue[int]{
			Comparator: func(o1, o2 int) int { return o1 - o2 },
		}
		s := make([]int, 0, 100)
		for j := 0; j < 100; j++ {
			v := r.Intn(99999)
			s = append(s, v)
			q.Add(v)
		}
		s1 := make([]int, 0, 100)
		for q.Len() > 0 {
			s1 = append(s1, q.Poll())
		}
		s2 := append([]int(nil), s...)
		slices.Sort(s2)
		if !slices.Equal(s1, s2) {
			t.Log("incorrect order: ", s)
			t.Fail()
		}
	}
}
