package main

import "testing"

func TestMaxValue(t *testing.T) {
	high := MaxValue([]int{56, 4, 7, 72, 150, -10})
	if high != 150 {
		t.Errorf("MaxValue([]int{56, 4, 7, 72, 150, -10}) = %d; want 150", high)
	}
}
