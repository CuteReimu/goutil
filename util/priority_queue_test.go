package util

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
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
