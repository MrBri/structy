package main

import "testing"

func TestMostFrequentChar(t *testing.T) {
	tests := []struct {
		word string
		char string
	}{{"bookeeper", "e"}, {"david", "d"}, {"abby", "b"}, {"mississippi", "i"}, {"potato", "o"},
		{"eleventennine", "e"}, {"riverbed", "r"},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			ans := MostFrequentChar(tt.word)
			if ans != tt.char {
				t.Errorf("got %s, want %s", ans, tt.char)
			}
		})
	}
}

/*
test_00:
most_frequent_char('bookeeper') # -> 'e'
test_01:
most_frequent_char('david') # -> 'd'
test_02:
most_frequent_char('abby') # -> 'b'
test_03:
most_frequent_char('mississippi') # -> 'i'
test_04:
most_frequent_char('potato') # -> 'o'
test_05:
most_frequent_char('eleventennine') # -> 'e'
test_06:
most_frequent_char('riverbed') # -> 'r'
*/
