package main

import (
	"reflect"
)

func Anagrams(f, s string) bool {
	seenf, seens := map[string]int{}, map[string]int{}
	for _, r := range f {
		seenf[string(r)]++
	}
	for _, r := range s {
		seens[string(r)]++
	}
	return reflect.DeepEqual(seenf, seens)
}
