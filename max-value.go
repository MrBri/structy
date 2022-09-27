package main

func MaxValue(nums []int) int {
	high := nums[0]
	for _, n := range nums {
		if n > high {
			high = n
		}
	}
	return high
}
