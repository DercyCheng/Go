package interview

// 遍历数组：使用for range循环遍历每个元素。
// 计算索引：index := num - 1将元素值转换为有效索引。
// 检查标记：若nums[index]为负，说明当前元素已出现过，加入结果。
// 标记出现：若nums[index]为正，将其取反标记为已访问。
// 返回结果：最终res包含所有重复元素。

func findDuplicates(nums []int) []int {
	var res []int
	for _, num := range nums {
		index := num - 1
		if nums[index] < 0 {
			res = append(res, num)
		} else {
			nums[index] *= -1
		}
	}
	return res
}
