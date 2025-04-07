package Tree

// constructMaximumBinaryTree 根据给定的数组构建最大二叉树
// 参数：
// nums - 输入的整数数组
// 返回值：
// 返回构建的最大二叉树的根节点
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxVal := nums[0]
	maxIndex := 0
	for i, num := range nums {
		if num > maxVal {
			maxVal = num
			maxIndex = i
		}
	}
	root := &TreeNode{Val: maxVal}
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	return root
}
