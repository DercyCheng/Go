package Stack_Queue

import "sort"

func topKFrequent(nums []int, k int) []int {
	var ans []int
	mapNum := map[int]int64{}
	for _, item := range nums {
		mapNum[item]++
	}
	for key, _ := range mapNum {
		ans = append(ans, key)
	}
	sort.Slice(ans, func(i, j int) bool {
		return mapNum[ans[i]] > mapNum[ans[j]]
	})
	return ans[:k]
}
