package main

import (
	"strconv"
	"strings"
)

func Compress(s string) string {
	builder := []string{}
	i, j := 0, 0
	s += "t"
	for j < len(s) {
		if s[i] == s[j] {
			j++
		} else {
			count := j - i
			if count == 1 {
				builder = append(builder, string(s[i]))
			} else {
				builder = append(builder, strconv.Itoa(count), string(s[i]))
			}
			i = j
		}
	}
	return strings.Join(builder, "")
}
