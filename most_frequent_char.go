package main

func MostFrequentChar(s string) string {
	seen := map[rune]int{}
	var max rune
	for _, r := range s {
		seen[r]++
	}
	for k, v := range seen {
		if v > seen[max] {
			max = k
		}
	}

	return string(max)
}
