package main

import (
	"strconv"
	"testing"
)

func TestIsPrime(t *testing.T) {
	var tests = []struct {
		n       int
		isprime bool
	}{{1, false}, {7, true}, {9, false}, {11, true}, {8, false}, {6, false}}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			ans := IsPrime(tt.n)
			if ans != tt.isprime {
				t.Errorf("got %t, want %t", ans, tt.isprime)
			}

		})
	}
}
