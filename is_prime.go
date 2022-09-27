package main

import "math"

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	sq := math.Sqrt(float64(n))
	for i := 2; i <= int(sq); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
