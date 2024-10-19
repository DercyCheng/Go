package Tree

// sumOfLeftLeaves 计算所有左叶子节点的和
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0 // 如果根节点为空，返回0
	}
	res := 0               // 初始化结果变量
	q := []*TreeNode{root} // 初始化队列，将根节点加入队列
	for len(q) > 0 {
		node := q[0] // 从队列前端取出一个节点
		q = q[1:]    // 移除已取出的节点
		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				res += node.Left.Val // 如果左子节点是叶子节点，将其值加到结果变量中
			} else {
				q = append(q, node.Left) // 如果左子节点不是叶子节点，将其加入队列
			}
		}
		if node.Right != nil {
			q = append(q, node.Right) // 如果右子节点不为空，将其加入队列
		}
	}
	return res // 返回结果变量
}
