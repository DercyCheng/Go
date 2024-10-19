package Tree

// isSymmetric 检查二叉树是否对称
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true // 空树是对称的
	}

	var queue []*TreeNode
	queue = append(queue, root.Left, root.Right) // 初始化队列，将根节点的左子节点和右子节点加入队列

	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:] // 从队列前端取出两个节点

		if left == nil && right == nil {
			continue // 如果两个节点都为空，继续下一次循环
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false // 如果其中一个节点为空或者两个节点的值不相等，返回 false
		}

		// 将左节点的左子节点和右节点的右子节点加入队列
		queue = append(queue, left.Left, right.Right)
		// 将左节点的右子节点和右节点的左子节点加入队列
		queue = append(queue, left.Right, right.Left)
	}

	return true // 如果遍历完所有节点都没有发现不对称的情况，返回 true
}
