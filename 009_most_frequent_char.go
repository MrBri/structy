package main

func MostFrequentChar(s string) string {
	seen := map[string]int{} // go randomizes map order
	var max string
	order := []string{} // order must be maintained separately, https://nathanleclaire.com/blog/2014/04/27/a-surprising-feature-of-golang-that-colored-me-impressed/

	for _, r := range s {
		order = append(order, string(r))
		seen[string(r)]++
	}

	for _, r := range order {
		if seen[string(r)] > seen[max] {
			max = string(r)
		}
	}
	return max
}
