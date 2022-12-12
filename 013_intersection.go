package main

import (
	"sort"
)

func Intersection(arr1, arr2 []int) []int {
	seen := map[int]bool{}
	result := []int{}
	for _, n := range arr1 {
		seen[n] = true
	}

	for _, n := range arr2 {
		if seen[n] {
			result = append(result, n)
		}
	}
	sort.Ints(result)
	return result
}
