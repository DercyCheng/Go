package Tree

import "strconv"

// binaryTreePaths 返回所有从根节点到叶子节点的路径
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{} // 如果根节点为空，返回空数组
	}

	var paths []string // 存储所有路径的数组
	type NodePath struct {
		node *TreeNode // 当前节点
		path string    // 从根节点到当前节点的路径
	}

	// 初始化队列，将根节点和其值作为路径字符串加入队列
	q := []NodePath{{root, strconv.Itoa(root.Val)}}

	// 遍历队列，直到队列为空
	for len(q) > 0 {
		np := q[0] // 从队列前端取出一个节点和路径
		q = q[1:]  // 移除已取出的节点

		node := np.node // 当前节点
		path := np.path // 从根节点到当前节点的路径

		if node.Left == nil && node.Right == nil {
			paths = append(paths, path) // 如果是叶子节点，将路径添加到结果数组中
		} else {
			if node.Left != nil {
				// 将左子节点及其路径加入队列
				q = append(q, NodePath{node.Left, path + "->" + strconv.Itoa(node.Left.Val)})
			}
			if node.Right != nil {
				// 将右子节点及其路径加入队列
				q = append(q, NodePath{node.Right, path + "->" + strconv.Itoa(node.Right.Val)})
			}
		}
	}

	return paths // 返回所有路径
}
