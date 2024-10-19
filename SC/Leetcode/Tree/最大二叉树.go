package Tree

// constructMaximumBinaryTree 根据给定的数组构建最大二叉树
// 参数：
// nums - 输入的整数数组
// 返回值：
// 返回构建的最大二叉树的根节点
func constructMaximumBinaryTree(nums []int) *TreeNode {
	// 如果数组为空，返回 nil
	if len(nums) == 0 {
		return nil
	}

	// 定义一个内部函数 findMax，用于找到数组中的最大值及其索引
	findMax := func(nums []int) (idx int) {
		// 遍历数组，找到最大值的索引
		for i, v := range nums {
			if nums[idx] < v {
				idx = i
			}
		}
		return
	}

	// 找到数组中的最大值的索引
	idx := findMax(nums)

	// 创建根节点，值为数组中的最大值
	root := &TreeNode{
		Val:   nums[idx],                                // 节点值为最大值
		Left:  constructMaximumBinaryTree(nums[:idx]),   // 递归构建左子树
		Right: constructMaximumBinaryTree(nums[idx+1:]), // 递归构建右子树
	}

	// 返回根节点
	return root
}
