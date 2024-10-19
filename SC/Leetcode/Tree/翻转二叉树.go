package Tree

// 递归前序
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// 递归后序
func invertTreeV1(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	invertTree(node.Left)
	invertTree(node.Right)
	node.Left, node.Right = node.Right, node.Left

	return node
}
