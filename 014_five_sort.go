package main

func FiveSort(arr []int) []int {
	i := 0
	j := len(arr) - 1
	for i < j {
		if arr[j] == 5 {
			j--
		} else if arr[i] == 5 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		} else {
			i++
		}
	}
	return arr
}
