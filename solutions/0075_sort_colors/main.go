package main

func sortColors(nums []int) {
	counts := [3]int{}

	for _, num := range nums {
		counts[num]++
	}

	idx := 0

	for color, count := range counts {
		for i := 0; i < count; i++ {
			nums[idx] = color
			idx++
		}
	}
}
