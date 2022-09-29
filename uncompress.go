package main

import (
	"strconv"
	"strings"
)

func Uncompress(s string) string {
	result := []string{}
	for i, j := 0, 0; j < len(s); {
		if s[j] > 47 && s[j] < 58 {
			j++
		} else {
			num, _ := strconv.Atoi(s[i:j])
			for k := 0; k < num; k++ {
				result = append(result, string(s[j]))
			}
			j++
			i = j
		}
	}
	return strings.Join(result, "")
}
