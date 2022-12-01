package main

import "testing"

func TestAnagrams(t *testing.T) {
	tests := []struct {
		first  string
		second string
		answer bool
	}{{"restful", "fluster", true}, {"cats", "tocs", false}, {"monkeyswrite", "newyorktimes", true},
		{"paper", "reapa", false},
		{"elbow", "below", true},
		{"tax", "taxi", false},
		{"taxi", "tax", false},
		{"night", "thing", true},
		{"abbc", "aabc", false},
		{"po", "popp", false},
		{"pp", "oo", false},
	}

	for _, tt := range tests {
		t.Run(tt.first, func(t *testing.T) {
			ans := Anagrams(tt.first, tt.second)
			if ans != tt.answer {
				t.Errorf("got %t, want %t", ans, tt.answer)
			}
		})
	}
}

/*
test_00:
{"restful", "fluster", {true
test_01:
{"cats", "tocs", {false
test_02:
{"monkeyswrite", "newyorktimes", {true
test_03:
{"paper", "reapa", {false
test_04:
{"elbow", "below", {true
test_05:
{"tax", "taxi", {false
test_06:
{"taxi", "tax", {false
test_07:
{"night", "thing", {true
test_08:
{"abbc", "aabc", {false
test_09:
{"po", "popp", {false
test_10:
{"pp", "oo") // -> false
*/
