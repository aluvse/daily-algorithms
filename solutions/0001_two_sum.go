package main

func twoSum(nums []int, target int) []int {
	prevMap := make(map[int]int)

	for i, x := range nums {
		complement := target - x

		if idx, exists := prevMap[complement]; exists {
			return []int{idx, i}
		}
		prevMap[x] = i
	}
	return nil

}
