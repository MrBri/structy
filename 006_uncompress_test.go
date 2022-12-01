package main

import "testing"

func TestUncompress(t *testing.T) {
	var tests = []struct {
		code         string
		uncompressed string
	}{
		{"2c3a1t", "ccaaat"},
		{"4s2b", "ssssbb"},
		{"2p1o5p", "ppoppppp"},
		{"3n12e2z", "nnneeeeeeeeeeeezz"},
		{"127y", "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy"},
	}

	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			ans := Uncompress(tt.code)
			if ans != tt.uncompressed {
				t.Errorf("got %s, want %s", ans, tt.uncompressed)
			}
		})
	}
}

//uncompress("2c3a1t"); // -> 'ccaaat'
// uncompress("4s2b") # -> 'ssssbb'
// uncompress("2p1o5p") # -> 'ppoppppp'
//uncompress("3n12e2z"); // -> 'nnneeeeeeeeeeeezz'
// uncompress("127y"); // ->'yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy'
