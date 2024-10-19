package Tree

import "math"

// isBalanced 检查二叉树是否平衡
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true // 空树是平衡的
	}

	// NodeDepth 结构体包含节点和深度
	type NodeDepth struct {
		node  *TreeNode
		depth int
	}

	q := []NodeDepth{{root, 1}} // 初始化队列，将根节点和深度1加入队列

	for len(q) > 0 {
		nd := q[0]
		q = q[1:] // 从队列前端取出一个节点和深度

		leftDepth := 0
		if nd.node.Left != nil {
			leftDepth = nd.depth + 1
			q = append(q, NodeDepth{nd.node.Left, leftDepth}) // 将左子节点和更新后的深度加入队列
		}

		rightDepth := 0
		if nd.node.Right != nil {
			rightDepth = nd.depth + 1
			q = append(q, NodeDepth{nd.node.Right, rightDepth}) // 将右子节点和更新后的深度加入队列
		}

		if math.Abs(float64(leftDepth-rightDepth)) > 1 {
			return false // 如果左右子树深度差大于1，返回false
		}
	}

	return true // 遍历完所有节点后，返回true
}
