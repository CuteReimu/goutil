package goutil

import (
	"fmt"
	"math/rand"
	"slices"
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

func TestCreatorAndPoll3(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	for i := 0; i < 100; i++ {
		s := make([]int, 0, 100)
		for j := 0; j < 100; j++ {
			s = append(s, r.Intn(99999))
		}
		q := NewDefaultPriorityQueue(append([]int(nil), s...))
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
	q := NewPriorityQueue[int](nil, func(o1, o2 int) int { return o1 - o2 })
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
		q := NewDefaultPriorityQueue[int](nil)
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

// TestRemoveWithComparator tests Remove method for priorityQueueWithComparator
func TestRemoveWithComparator(t *testing.T) {
	// Test removing root element
	t.Run("RemoveRoot", func(t *testing.T) {
		q := NewPriorityQueue([]int{4, 6, 3, 5, 9}, func(o1, o2 int) int { return o1 - o2 })
		if q.Len() != 5 {
			t.Errorf("Expected length 5, got %d", q.Len())
		}
		// Root should be 3 (minimum)
		if !q.Remove(3) {
			t.Error("Failed to remove root element 3")
		}
		if q.Len() != 4 {
			t.Errorf("Expected length 4 after removal, got %d", q.Len())
		}
		// Verify heap integrity by polling all elements
		var result []int
		for q.Len() > 0 {
			result = append(result, q.Poll())
		}
		expected := []int{4, 5, 6, 9}
		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test removing non-root element (middle)
	t.Run("RemoveMiddle", func(t *testing.T) {
		q := NewPriorityQueue([]int{4, 6, 3, 5, 9}, func(o1, o2 int) int { return o1 - o2 })
		if !q.Remove(5) {
			t.Error("Failed to remove middle element 5")
		}
		if q.Len() != 4 {
			t.Errorf("Expected length 4 after removal, got %d", q.Len())
		}
		// Verify heap integrity
		var result []int
		for q.Len() > 0 {
			result = append(result, q.Poll())
		}
		expected := []int{3, 4, 6, 9}
		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test removing last element
	t.Run("RemoveLast", func(t *testing.T) {
		q := NewPriorityQueue([]int{4, 6, 3, 5, 9}, func(o1, o2 int) int { return o1 - o2 })
		if !q.Remove(9) {
			t.Error("Failed to remove last element 9")
		}
		if q.Len() != 4 {
			t.Errorf("Expected length 4 after removal, got %d", q.Len())
		}
		// Verify heap integrity
		var result []int
		for q.Len() > 0 {
			result = append(result, q.Poll())
		}
		expected := []int{3, 4, 5, 6}
		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test removing non-existent element
	t.Run("RemoveNonExistent", func(t *testing.T) {
		q := NewPriorityQueue([]int{4, 6, 3, 5, 9}, func(o1, o2 int) int { return o1 - o2 })
		if q.Remove(100) {
			t.Error("Should not remove non-existent element 100")
		}
		if q.Len() != 5 {
			t.Errorf("Length should remain 5, got %d", q.Len())
		}
	})

	// Test removing with duplicates (comparator returns 0)
	t.Run("RemoveWithDuplicates", func(t *testing.T) {
		q := NewPriorityQueue([]int{4, 6, 3, 5, 3, 9}, func(o1, o2 int) int { return o1 - o2 })
		if q.Len() != 6 {
			t.Errorf("Expected length 6, got %d", q.Len())
		}
		// Remove first occurrence of 3
		if !q.Remove(3) {
			t.Error("Failed to remove element 3")
		}
		if q.Len() != 5 {
			t.Errorf("Expected length 5 after removal, got %d", q.Len())
		}
		// Verify heap still contains the other 3 and maintains order
		var result []int
		for q.Len() > 0 {
			result = append(result, q.Poll())
		}
		expected := []int{3, 4, 5, 6, 9}
		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test multiple removals
	t.Run("MultipleRemovals", func(t *testing.T) {
		q := NewPriorityQueue([]int{4, 6, 3, 5, 9, 1, 8}, func(o1, o2 int) int { return o1 - o2 })
		// Remove multiple elements
		if !q.Remove(6) {
			t.Error("Failed to remove 6")
		}
		if !q.Remove(1) {
			t.Error("Failed to remove 1")
		}
		if !q.Remove(9) {
			t.Error("Failed to remove 9")
		}
		if q.Len() != 4 {
			t.Errorf("Expected length 4 after removals, got %d", q.Len())
		}
		// Verify heap integrity
		var result []int
		for q.Len() > 0 {
			result = append(result, q.Poll())
		}
		expected := []int{3, 4, 5, 8}
		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

// TestRemoveDefault tests Remove method for defaultPriorityQueue
func TestRemoveDefault(t *testing.T) {
	// Test removing root element
	t.Run("RemoveRoot", func(t *testing.T) {
		q := NewDefaultPriorityQueue([]int{4, 6, 3, 5, 9})
		if q.Len() != 5 {
			t.Errorf("Expected length 5, got %d", q.Len())
		}
		// Root should be 3 (minimum)
		if !q.Remove(3) {
			t.Error("Failed to remove root element 3")
		}
		if q.Len() != 4 {
			t.Errorf("Expected length 4 after removal, got %d", q.Len())
		}
		// Verify heap integrity by polling all elements
		var result []int
		for q.Len() > 0 {
			result = append(result, q.Poll())
		}
		expected := []int{4, 5, 6, 9}
		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test removing non-root element (middle)
	t.Run("RemoveMiddle", func(t *testing.T) {
		q := NewDefaultPriorityQueue([]int{4, 6, 3, 5, 9})
		if !q.Remove(5) {
			t.Error("Failed to remove middle element 5")
		}
		if q.Len() != 4 {
			t.Errorf("Expected length 4 after removal, got %d", q.Len())
		}
		// Verify heap integrity
		var result []int
		for q.Len() > 0 {
			result = append(result, q.Poll())
		}
		expected := []int{3, 4, 6, 9}
		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test removing last element
	t.Run("RemoveLast", func(t *testing.T) {
		q := NewDefaultPriorityQueue([]int{4, 6, 3, 5, 9})
		if !q.Remove(9) {
			t.Error("Failed to remove last element 9")
		}
		if q.Len() != 4 {
			t.Errorf("Expected length 4 after removal, got %d", q.Len())
		}
		// Verify heap integrity
		var result []int
		for q.Len() > 0 {
			result = append(result, q.Poll())
		}
		expected := []int{3, 4, 5, 6}
		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test removing non-existent element
	t.Run("RemoveNonExistent", func(t *testing.T) {
		q := NewDefaultPriorityQueue([]int{4, 6, 3, 5, 9})
		if q.Remove(100) {
			t.Error("Should not remove non-existent element 100")
		}
		if q.Len() != 5 {
			t.Errorf("Length should remain 5, got %d", q.Len())
		}
	})

	// Test removing with duplicates
	t.Run("RemoveWithDuplicates", func(t *testing.T) {
		q := NewDefaultPriorityQueue([]int{4, 6, 3, 5, 3, 9})
		if q.Len() != 6 {
			t.Errorf("Expected length 6, got %d", q.Len())
		}
		// Remove first occurrence of 3
		if !q.Remove(3) {
			t.Error("Failed to remove element 3")
		}
		if q.Len() != 5 {
			t.Errorf("Expected length 5 after removal, got %d", q.Len())
		}
		// Verify heap still contains the other 3 and maintains order
		var result []int
		for q.Len() > 0 {
			result = append(result, q.Poll())
		}
		expected := []int{3, 4, 5, 6, 9}
		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test multiple removals
	t.Run("MultipleRemovals", func(t *testing.T) {
		q := NewDefaultPriorityQueue([]int{4, 6, 3, 5, 9, 1, 8})
		// Remove multiple elements
		if !q.Remove(6) {
			t.Error("Failed to remove 6")
		}
		if !q.Remove(1) {
			t.Error("Failed to remove 1")
		}
		if !q.Remove(9) {
			t.Error("Failed to remove 9")
		}
		if q.Len() != 4 {
			t.Errorf("Expected length 4 after removals, got %d", q.Len())
		}
		// Verify heap integrity
		var result []int
		for q.Len() > 0 {
			result = append(result, q.Poll())
		}
		expected := []int{3, 4, 5, 8}
		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test with strings to verify it works with other ordered types
	t.Run("RemoveStrings", func(t *testing.T) {
		q := NewDefaultPriorityQueue([]string{"dog", "cat", "elephant", "ant", "bear"})
		if !q.Remove("cat") {
			t.Error("Failed to remove 'cat'")
		}
		if q.Len() != 4 {
			t.Errorf("Expected length 4 after removal, got %d", q.Len())
		}
		// Verify heap integrity
		var result []string
		for q.Len() > 0 {
			result = append(result, q.Poll())
		}
		expected := []string{"ant", "bear", "dog", "elephant"}
		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

// TestToSliceWithLargerInput tests ToSlice with preallocated slice larger than queue
func TestToSliceWithLargerInput(t *testing.T) {
	// Test with comparator-based queue
	t.Run("WithComparator", func(t *testing.T) {
		q := NewPriorityQueue([]int{3, 1, 4, 1, 5}, func(o1, o2 int) int { return o1 - o2 })
		// Create a slice with stale data that's larger than the queue
		input := []int{99, 99, 99, 99, 99, 99, 99, 99, 99, 99}
		result := q.ToSlice(input)

		// Result should have length 5 (queue size), not 10 (input size)
		if len(result) != 5 {
			t.Errorf("Expected length 5, got %d", len(result))
		}

		// Result should not contain any stale 99 values
		for _, v := range result {
			if v == 99 {
				t.Errorf("Result contains stale value 99: %v", result)
				break
			}
		}

		// Verify the result contains all queue elements
		expected := []int{1, 1, 3, 4, 5}
		slices.Sort(result)
		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test with default queue
	t.Run("Default", func(t *testing.T) {
		q := NewDefaultPriorityQueue([]int{3, 1, 4, 1, 5})
		// Create a slice with stale data that's larger than the queue
		input := []int{99, 99, 99, 99, 99, 99, 99, 99, 99, 99}
		result := q.ToSlice(input)

		// Result should have length 5 (queue size), not 10 (input size)
		if len(result) != 5 {
			t.Errorf("Expected length 5, got %d", len(result))
		}

		// Result should not contain any stale 99 values
		for _, v := range result {
			if v == 99 {
				t.Errorf("Result contains stale value 99: %v", result)
				break
			}
		}

		// Verify the result contains all queue elements
		expected := []int{1, 1, 3, 4, 5}
		slices.Sort(result)
		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

// TestForeachWithMutation tests Foreach when callback mutates the queue
func TestForeachWithMutation(t *testing.T) {
	// Test with comparator-based queue
	t.Run("WithComparator", func(t *testing.T) {
		q := NewPriorityQueue([]int{1, 2, 3, 4, 5}, func(o1, o2 int) int { return o1 - o2 })
		var visited []int

		// Callback that mutates the queue during iteration
		q.Foreach(func(e int) bool {
			visited = append(visited, e)
			// Try to mutate the queue
			if e == 3 {
				q.Add(10)   // Add a new element
				q.Remove(4) // Remove an element
			}
			return true
		})

		// Should have visited exactly 5 elements (the original snapshot)
		if len(visited) != 5 {
			t.Errorf("Expected to visit 5 elements, visited %d: %v", len(visited), visited)
		}

		// The visited elements should be exactly the original 5 elements (1,2,3,4,5)
		// Sort and compare to verify each element was visited exactly once
		slices.Sort(visited)
		expected := []int{1, 2, 3, 4, 5}
		if !slices.Equal(visited, expected) {
			t.Errorf("Expected to visit %v exactly once each, got %v", expected, visited)
		}

		// Verify queue was actually mutated (should now have 10, not 4)
		if !q.Contains(10) {
			t.Error("Queue should contain 10 after mutation")
		}
		if q.Contains(4) {
			t.Error("Queue should not contain 4 after mutation")
		}
	})

	// Test with default queue
	t.Run("Default", func(t *testing.T) {
		q := NewDefaultPriorityQueue([]int{1, 2, 3, 4, 5})
		var visited []int

		// Callback that mutates the queue during iteration
		q.Foreach(func(e int) bool {
			visited = append(visited, e)
			// Try to mutate the queue
			if e == 3 {
				q.Add(10)   // Add a new element
				q.Remove(4) // Remove an element
			}
			return true
		})

		// Should have visited exactly 5 elements (the original snapshot)
		if len(visited) != 5 {
			t.Errorf("Expected to visit 5 elements, visited %d: %v", len(visited), visited)
		}

		// The visited elements should be exactly the original 5 elements (1,2,3,4,5)
		// Sort and compare to verify each element was visited exactly once
		slices.Sort(visited)
		expected := []int{1, 2, 3, 4, 5}
		if !slices.Equal(visited, expected) {
			t.Errorf("Expected to visit %v exactly once each, got %v", expected, visited)
		}

		// Verify queue was actually mutated (should now have 10, not 4)
		if !q.Contains(10) {
			t.Error("Queue should contain 10 after mutation")
		}
		if q.Contains(4) {
			t.Error("Queue should not contain 4 after mutation")
		}
	})
}
