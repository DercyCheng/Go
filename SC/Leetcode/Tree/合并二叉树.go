package Tree

// mergeTrees 合并两棵二叉树
// 参数：
// root1 - 第一棵二叉树的根节点
// root2 - 第二棵二叉树的根节点
// 返回值：
// 返回合并后的二叉树的根节点
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return nil
	}
	if root2 == nil {
		return nil
	}
	root := &TreeNode{Val: root1.Val + root2.Val}

	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}
