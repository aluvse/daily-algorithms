package main

func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))

	for _, x := range nums {
		if _, exists := seen[x]; exists {
			return true
		}
		seen[x] = struct{}{}
	}
	return false
}
