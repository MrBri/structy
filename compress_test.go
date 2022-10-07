package main

import "testing"

func TestCompress(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{{"ccaaatsss", "2c3at3s"}, {"ssssbbz", "4s2bz"}}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			ans := Compress(tt.in)
			if ans != tt.out {
				t.Errorf("got %s, want %s", ans, tt.out)
			}
		})
	}
}

/**
compress('ccaaatsss'); // -> '2c3at3s'
test_01:
compress('ssssbbz'); // -> '4s2bz'
test_02:
compress('ppoppppp'); // -> '2po5p'
test_03:
compress('nnneeeeeeeeeeeezz'); // -> '3n12e2z'
test_04:
compress('yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy');
// -> '127y'
 **/
