package main

import "testing"

func TestKthLargestStream(t *testing.T) {
	s := NewKthLargest(3, []int{4, 5, 8, 2})
	steps := []struct {
		add, want int
	}{
		{3, 4}, {5, 5}, {10, 5}, {9, 8}, {4, 8},
	}
	for _, st := range steps {
		got, ok := s.Add(st.add)
		if !ok || got != st.want {
			t.Errorf("Add(%d) = (%d,%v), want %d", st.add, got, ok, st.want)
		}
	}
}

func TestUnderfilled(t *testing.T) {
	s := NewKthLargest(2, nil)
	if _, ok := s.Add(5); ok {
		t.Error("one value cannot have a 2nd largest")
	}
	if got, ok := s.Add(3); !ok || got != 3 {
		t.Errorf("second value: (%d,%v), want 3", got, ok)
	}
	if got, _ := s.Add(10); got != 5 {
		t.Errorf("after 10: %d, want 5", got)
	}
}
