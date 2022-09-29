package main

import "testing"

func TestUncompress(t *testing.T) {
	var tests = []struct {
		code         string
		uncompressed string
	}{{"2c3a1t", "ccaaat"}, {"3n12e2z", "nnneeeeeeeeeeeezz"}, {"127y", "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy"}}

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
//uncompress("3n12e2z"); // -> 'nnneeeeeeeeeeeezz'
// uncompress("127y"); // ->'yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy'
