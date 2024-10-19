package Tree

// hasPathSum 判断是否存在从根节点到叶子节点的路径，使得路径上所有节点值之和等于目标和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false // 如果根节点为空，返回false
	}

	// 初始化队列，将根节点和目标和加入队列
	q := []struct {
		node *TreeNode
		sum  int
	}{{root, targetSum}}

	// 当队列不为空时，继续遍历
	for len(q) > 0 {
		cur := q[0] // 取出队列前端的元素
		q = q[1:]   // 移除已取出的元素

		node := cur.node
		sum := cur.sum

		// 如果当前节点是叶子节点，且路径和等于目标和，返回true
		if node.Left == nil && node.Right == nil && node.Val == sum {
			return true
		}

		// 将左子节点加入队列，并更新路径和
		if node.Left != nil {
			q = append(q, struct {
				node *TreeNode
				sum  int
			}{node.Left, sum - node.Val})
		}

		// 将右子节点加入队列，并更新路径和
		if node.Right != nil {
			q = append(q, struct {
				node *TreeNode
				sum  int
			}{node.Right, sum - node.Val})
		}
	}

	return false // 如果没有找到满足条件的路径，返回false
}
