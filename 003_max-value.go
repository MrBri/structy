package main

func MaxValue(nums []float32) float32 {
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}
