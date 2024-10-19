package Tree

// minDepth 计算二叉树的最小深度
// 使用层次遍历（广度优先搜索）的方法
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0 // 如果根节点为空，返回深度0
	}
	queue := []*TreeNode{root} // 初始化队列，将根节点加入队列
	depth := 1                 // 初始化深度变量为1
	for len(queue) > 0 {
		size := len(queue) // 获取当前层的节点数量
		for i := 0; i < size; i++ {
			node := queue[0]  // 从队列前端取出一个节点
			queue = queue[1:] // 移除已取出的节点
			if node.Left == nil && node.Right == nil {
				return depth // 如果当前节点是叶子节点，返回当前深度
			}
			if node.Left != nil {
				queue = append(queue, node.Left) // 如果左子节点不为空，将其加入队列
			}
			if node.Right != nil {
				queue = append(queue, node.Right) // 如果右子节点不为空，将其加入队列
			}
		}
		depth++ // 每遍历完一层，深度加1
	}
	return depth // 返回深度变量
}
