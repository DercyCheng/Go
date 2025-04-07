package interview

//初始化技巧：
//
//dp[0] = 0（重量0无需球）
//其他初始值设为极大值（如amount+1，因最多用amount个1重量的球）
//双重循环结构：
//
//外层循环：遍历所有重量1~amount
//内层循环：尝试每个球的重量，若球重 ≤ 当前重量则更新dp[i]
//结果判断：
//最终检查dp[amount]是否超过amount，若超过则返回-1

func minBalls(balls []int, amount int) int {
	max := amount + 1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = max // 初始化dp数组，所有值设为amount+1
	}
	dp[0] = 0 // 重量0需要0个球

	for i := 1; i <= amount; i++ {
		for _, w := range balls {
			if w <= i && dp[i-w]+1 < dp[i] {
				dp[i] = dp[i-w] + 1 // 更新dp[i]为使用当前球重量后的最小值
			}
		}
	}

	if dp[amount] > amount {
		return -1 // 如果dp[amount]大于amount，说明无法凑出该重量
	}
	return dp[amount] // 返回凑出该重量所需的最少球数
}
