package main

import (
	"strconv"
)

func Compress(s string) string {
	// result := []string{}
	result := ""
	i, j := 0, 0
	count := 0
	for j < len(s) {
		if s[i] == s[j] {
			j++
		} else {
			count = j - i
			if count == 1 {
				result += string(s[i])
				// result = append(result, string(s[i]))
			} else {
				result += strconv.Itoa(count)
				result += string(s[i])
				// result = append(result, strconv.Itoa(count), string(s[i]))
			}
			i = j
		}
	}
	count = j - i
	if count == 1 {
		result += string(s[i])
		// result = append(result, string(s[i]))
	} else {
		result += strconv.Itoa(count)
		result += string(s[i])
		// result = append(result, strconv.Itoa(count), string(s[i]))
	}
	// return strings.Join(result, "")
	return result
}
