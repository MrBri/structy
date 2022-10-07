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
anagrams('restful', 'fluster'); // -> true
test_01:
anagrams('cats', 'tocs'); // -> false
test_02:
anagrams('monkeyswrite', 'newyorktimes'); // -> true
test_03:
anagrams('paper', 'reapa'); // -> false
test_04:
anagrams('elbow', 'below'); // -> true
test_05:
anagrams('tax', 'taxi'); // -> false
test_06:
anagrams('taxi', 'tax'); // -> false
test_07:
anagrams('night', 'thing'); // -> true
test_08:
anagrams('abbc', 'aabc'); // -> false
test_09:
anagrams('po', 'popp'); // -> false
test_10:
anagrams('pp', 'oo') // -> false
*/
