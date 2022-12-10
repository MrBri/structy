package main

func PairProduct(values []int, product int) []int {
	seen := map[int]int{}

	for i, n := range values {
		compliment := product / n
		if in, ok := seen[compliment]; ok {
			return []int{in, i}
		}
		seen[n] = i
	}
	return []int{}
}
