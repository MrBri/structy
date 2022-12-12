package main

import (
	"strconv"
	"testing"
)

func TestIntersection(t *testing.T) {

	tests := []struct {
		arr1      []int
		arr2      []int
		intersect []int // indeces
	}{
		{[]int{4, 2, 1, 6}, []int{3, 6, 9, 2, 10}, []int{2, 6}},
		{[]int{2, 4, 6}, []int{4, 2}, []int{2, 4}},
		{[]int{4, 2, 1}, []int{1, 2, 4, 6}, []int{1, 2, 4}},
		{[]int{0, 1, 2}, []int{10, 11}, []int{}},
	}

	test := -01
	for _, tt := range tests {
		test++
		t.Run(strconv.Itoa(test), func(t *testing.T) {
			ans := Intersection(tt.arr1, tt.arr2)
			if len(ans) != len(tt.intersect) {
				t.Errorf("got: %v, want %v\n", ans, tt.intersect)
			}
			for i, v := range ans {
				if v != tt.intersect[i] {
					t.Errorf("got: %v, want %v\n", ans, tt.intersect)
				}
			}
		})
	}

	var a, b []int
	for i := 0; i < 50000; i++ {
		a = append(a, i)
		b = append(b, i)
	}

	intersect := Intersection(a, b)

	if len(intersect) != len(b) {
		t.Errorf("got: %v, want %v\n", intersect, b)
	}
}

/*
test_00:
intersection([4,2,1,6], [3,6,9,2,10]) // -> [2,6]
test_01:
intersection([2,4,6], [4,2]) // -> [2,4]
test_02:
intersection([4,2,1], [1,2,4,6]) // -> [1,2,4]
test_03:
intersection([0,1,2], [10,11]) // -> []
test_04:
const a = [];
const b = [];
for (let i = 0; i < 50000; i += 1) {
  a.push(i);
  b.push(i);
}
intersection(a, b) // -> [0,1,2,3,..., 49999]
*/
