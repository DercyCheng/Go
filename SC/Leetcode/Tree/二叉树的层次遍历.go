package Tree

// levelOrder 进行二叉树的层次遍历。
// 它使用队列来实现广度优先搜索，并返回每一层的节点值。
func levelOrder(root *TreeNode) (result [][]int) {
	if root == nil {
		return // 如果根节点为空，返回空结果
	}

	queue := []*TreeNode{root} // 初始化队列，将根节点推入队列

	for len(queue) > 0 {
		level := []int{}           // 初始化当前层的节点值切片
		nextQueue := []*TreeNode{} // 初始化下一层的节点队列

		for _, node := range queue {
			level = append(level, node.Val) // 将当前节点的值添加到当前层的结果中

			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left) // 将左子节点推入下一层队列
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right) // 将右子节点推入下一层队列
			}
		}

		result = append(result, level) // 将当前层的结果添加到最终结果中
		queue = nextQueue              // 更新队列为下一层的节点
	}

	return result // 返回包含每一层节点值的结果切片
}
