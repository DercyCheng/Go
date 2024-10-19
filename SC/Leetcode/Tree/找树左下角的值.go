package Tree

// findBottomLeftValue 查找二叉树最左下角的值
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0 // 如果根节点为空，返回0
	}

	q := []*TreeNode{root} // 初始化队列，将根节点加入队列
	var bottomLeftValue int

	for len(q) > 0 {
		node := q[0] // 从队列前端取出一个节点
		q = q[1:]    // 移除已取出的节点

		// 更新最左下角的值
		bottomLeftValue = node.Val

		// 先将右子节点加入队列
		if node.Right != nil {
			q = append(q, node.Right)
		}
		// 再将左子节点加入队列
		if node.Left != nil {
			q = append(q, node.Left)
		}
	}

	return bottomLeftValue // 返回最左下角的值
}
