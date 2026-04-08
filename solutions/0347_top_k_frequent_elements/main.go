package main

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)

	for _, n := range nums {
		countMap[n]++
	}

	buckets := make([][]int, len(nums)+1)

	for num, freq := range countMap {
		buckets[freq] = append(buckets[freq], num)
	}

	res := make([]int, 0, k)

	for i := len(buckets) - 1; i >= 0 && len(res) < k; i-- {
		if len(buckets[i]) > 0 {
			res = append(res, buckets[i]...)
		}
	}

	return res[:k]
}
