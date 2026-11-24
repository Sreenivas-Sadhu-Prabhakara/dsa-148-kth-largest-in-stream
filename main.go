package main

import (
	"container/heap"
	"fmt"
)

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

// KthLargest tracks the k-th largest element of a growing stream: a min-heap
// of the k largest values seen, whose root is always the answer.
type KthLargest struct {
	k int
	h intHeap
}

// NewKthLargest seeds the tracker with initial values.
func NewKthLargest(k int, nums []int) *KthLargest {
	s := &KthLargest{k: k}
	for _, v := range nums {
		s.Add(v)
	}
	return s
}

// Add inserts a value and returns the current k-th largest; ok is false
// while fewer than k values have been seen.
func (s *KthLargest) Add(v int) (int, bool) {
	if s.h.Len() < s.k {
		heap.Push(&s.h, v)
	} else if v > s.h[0] {
		s.h[0] = v
		heap.Fix(&s.h, 0)
	}
	if s.h.Len() < s.k {
		return 0, false
	}
	return s.h[0], true
}

func main() {
	s := NewKthLargest(3, []int{4, 5, 8, 2})
	for _, v := range []int{3, 5, 10, 9, 4} {
		kth, _ := s.Add(v)
		fmt.Print(kth, " ")
	}
	fmt.Println()
}
