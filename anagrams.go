package main

import (
	"reflect"
)

func Anagrams(s1, s2 string) bool {
	return reflect.DeepEqual(makeMap(s1), makeMap(s2))
}

func makeMap(s string) map[rune]int {
	seen := make(map[rune]int)
	for _, r := range s {
		seen[r]++
	}
	return seen
}
