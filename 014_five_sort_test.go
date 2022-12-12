package main

import (
	"strconv"
	"testing"
)

func TestFiveSort(t *testing.T) {

	tests := []struct {
		arr    []int
		sorted []int
	}{
		{[]int{12, 5, 1, 5, 12, 7}, []int{12, 7, 1, 12, 5, 5}},
		{[]int{5, 2, 5, 6, 5, 1, 10, 2, 5, 5}, []int{2, 2, 10, 6, 1, 5, 5, 5, 5, 5}},
		{[]int{5, 5, 5, 1, 1, 1, 4}, []int{4, 1, 1, 1, 5, 5, 5}},
		{[]int{5, 5, 6, 5, 5, 5, 5}, []int{6, 5, 5, 5, 5, 5, 5}},
		{[]int{5, 1, 2, 5, 5, 3, 2, 5, 1, 5, 5, 5, 4, 5}, []int{4, 1, 2, 1, 2, 3, 5, 5, 5, 5, 5, 5, 5, 5}},
	}

	test := -01
	for _, tt := range tests {
		test++
		t.Run(strconv.Itoa(test), func(t *testing.T) {
			ans := FiveSort(tt.arr)
			if len(ans) != len(tt.sorted) {
				t.Errorf("got: %v, want %v\n", ans, tt.sorted)
			}
			for i, v := range ans {
				if v != tt.sorted[i] {
					t.Errorf("got: %v, want %v\n", ans, tt.sorted)
				}
			}
		})
	}

	// var a, b []int
	// for i := 0; i < 50000; i++ {
	// 	a = append(a, i)
	// 	b = append(b, i)
	// }

	// intersect := Intersection(a, b)

	// if len(intersect) != len(b) {
	// 	t.Errorf("got: %v, want %v\n", intersect, b)
	// }
}

/*
test_00
five_sort([12, 5, 1, 5, 12, 7])
# -> [12, 7, 1, 12, 5, 5]
test_01
five_sort([5, 2, 5, 6, 5, 1, 10, 2, 5, 5])
# -> [2, 2, 10, 6, 1, 5, 5, 5, 5, 5]
test_02
five_sort([5, 5, 5, 1, 1, 1, 4])
# -> [4, 1, 1, 1, 5, 5, 5]
test_03
five_sort([5, 5, 6, 5, 5, 5, 5])
# -> [6, 5, 5, 5, 5, 5, 5]
test_04
five_sort([5, 1, 2, 5, 5, 3, 2, 5, 1, 5, 5, 5, 4, 5])
# -> [4, 1, 2, 1, 2, 3, 5, 5, 5, 5, 5, 5, 5, 5]
test_05
fours = [4] * 20000
fives = [5] * 20000
nums = fours + fives
five_sort(nums)
# twenty-thousand 4s followed by twenty-thousand 5s
# -> [4, 4, 4, 4, ..., 5, 5, 5, 5]
*/
