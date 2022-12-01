package main

import (
	"strconv"
	"testing"
)

func TestMaxValue(t *testing.T) {
	tests := []struct {
		values []float32
		out    float32
	}{
		{values: []float32{4, 7, 2, 8, 10, 9}, out: 10},
		{values: []float32{10, 5, 40, 40.3}, out: 40.3},
		{values: []float32{-5, -2, -1, -11}, out: -1},
		{values: []float32{42}, out: 42},
		{values: []float32{1000, 8}, out: 1000},
		{values: []float32{1000, 8, 9000}, out: 9000},
		{values: []float32{2, 5, 1, 1, 4}, out: 5},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			high := MaxValue(tt.values)
			if high != tt.out {
				t.Errorf("got = %v, want = %v", high, tt.out)
			}
		})
	}
}

/*
test_00:
max_value([4, 7, 2, 8, 10, 9]) # -> 10
test_01:
max_value([10, 5, 40, 40.3]) # -> 40.3
test_02:
max_value([-5, -2, -1, -11]) # -> -1
test_03:
max_value([42]) # -> 42
test_04:
max_value([1000, 8]) # -> 1000
test_05:
max_value([1000, 8, 9000]) # -> 9000
test_06:
max_value([2, 5, 1, 1, 4]) # -> 5
*/
