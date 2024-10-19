package Tree

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0 // 如果根节点为空，返回0
	}
	queue := []*TreeNode{root} // 初始化队列并将根节点添加到队列中
	count := 0                 // 初始化计数器
	for len(queue) > 0 {
		node := queue[0]  // 从队列的前端取出一个节点
		queue = queue[1:] // 移除已取出的节点

		if node.Left != nil {
			queue = append(queue, node.Left) // 如果有左子节点，将其加入队列
		}
		if node.Right != nil {
			queue = append(queue, node.Right) // 如果有右子节点，将其加入队列
		}
		count++ // 计数器加一
	}
	return count // 返回节点总数
}
