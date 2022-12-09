package main

func PairSum(values []int, sum int) []int {
	seen := map[int]int{}

	for i, n := range values {
		compliment := sum - n
		if in, ok := seen[compliment]; ok {
			return []int{in, i}
		}
		seen[n] = i
	}
	return []int{}
}
