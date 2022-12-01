package main

import (
	"strconv"
	"strings"
)

func Uncompress(s string) string {
	builder := []string{}
	i, j := 0, 0
	for _, r := range s {
		_, err := strconv.Atoi(string(r))
		if err == nil {
			j++
		} else {
			num := s[i:j]
			m, _ := strconv.Atoi(num)
			for k := 0; k < m; k++ {
				builder = append(builder, string(s[j]))
			}
			j++
			i = j
		}
	}
	return strings.Join(builder, "")
}
