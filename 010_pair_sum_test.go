package main

import (
	"strconv"
	"testing"
)

func TestPairSum(t *testing.T) {

	tests := []struct {
		values []int
		sum    int
		pair   []int // indeces
	}{
		{[]int{3, 2, 5, 4, 1}, 8, []int{0, 2}},
		{[]int{4, 7, 9, 2, 5, 1}, 5, []int{0, 5}},
		{[]int{4, 7, 9, 2, 5, 1}, 3, []int{3, 5}},
		{[]int{1, 6, 7, 2}, 13, []int{1, 2}},
		{[]int{9, 9}, 18, []int{0, 1}},
		{[]int{6, 4, 2, 8}, 12, []int{1, 3}},
	}

	test := -01
	for _, tt := range tests {
		test++
		t.Run(strconv.Itoa(test), func(t *testing.T) {
			ans := PairSum(tt.values, tt.sum)
			if ans[0] != tt.pair[0] || ans[1] != tt.pair[1] {
				t.Errorf("got: %v, want %v\n", ans, tt.pair)
			}
		})
	}

}

/*
test_00:
pair_sum([3, 2, 5, 4, 1], 8) # -> (0, 2)
test_01:
pair_sum([4, 7, 9, 2, 5, 1], 5) # -> (0, 5)
test_02:
pair_sum([4, 7, 9, 2, 5, 1], 3) # -> (3, 5)
test_03:
pair_sum([1, 6, 7, 2], 13) # -> (1, 2)
test_04:
pair_sum([9, 9], 18) # -> (0, 1)
test_05:
pair_sum([6, 4, 2, 8 ], 12) # -> (1, 3)
*/
