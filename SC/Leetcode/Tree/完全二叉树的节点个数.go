package Tree

// completeCountNodes 计算完全二叉树的节点个数
// 使用层次遍历（广度优先搜索）的方法
func completeCountNodes(root *TreeNode) int {
	if root == nil {
		return 0 // 如果根节点为空，返回0
	}
	q := []*TreeNode{root} // 初始化队列，将根节点加入队列
	res := 0               // 初始化计数器
	for len(q) > 0 {
		size := len(q) // 获取当前层的节点数量
		for i := 0; i < size; i++ {
			node := q[0] // 从队列前端取出一个节点
			q = q[1:]    // 移除已取出的节点
			if node.Left != nil {
				q = append(q, node.Left) // 如果左子节点不为空，将其加入队列
			}
			if node.Right != nil {
				q = append(q, node.Right) // 如果右子节点不为空，将其加入队列
			}
			res++ // 计数器加1
		}
	}
	return res // 返回节点总数
}
