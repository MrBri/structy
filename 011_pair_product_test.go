package main

import (
	"strconv"
	"testing"
)

func TestPairProduct(t *testing.T) {

	tests := []struct {
		values  []int
		product int
		pair    []int // indeces
	}{
		{[]int{3, 2, 5, 4, 1}, 8, []int{1, 3}},
		{[]int{3, 2, 5, 4, 1}, 10, []int{1, 2}},
		{[]int{4, 7, 9, 2, 5, 1}, 5, []int{4, 5}},
		{[]int{4, 7, 9, 2, 5, 1}, 35, []int{1, 4}},
		{[]int{3, 2, 5, 4, 1}, 10, []int{1, 2}},
		{[]int{4, 6, 8, 2}, 16, []int{2, 3}},
	}

	test := -01
	for _, tt := range tests {
		test++
		t.Run(strconv.Itoa(test), func(t *testing.T) {
			ans := PairProduct(tt.values, tt.product)
			if ans[0] != tt.pair[0] || ans[1] != tt.pair[1] {
				t.Errorf("got: %v, want %v\n", ans, tt.pair)
			}
		})
	}

}

/*
test_00:
pair_product([3, 2, 5, 4, 1], 8) # -> (1, 3)
test_01:
pair_product([3, 2, 5, 4, 1], 10) # -> (1, 2)
test_02:
pair_product([4, 7, 9, 2, 5, 1], 5) # -> (4, 5)
test_03:
pair_product([4, 7, 9, 2, 5, 1], 35) # -> (1, 4)
test_04:
pair_product([3, 2, 5, 4, 1], 10) # -> (1, 2)
test_05:
pair_product([4, 6, 8, 2], 16) # -> (2, 3)
*/
