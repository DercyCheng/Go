package Tree

// mergeTrees 合并两棵二叉树
// 参数：
// root1 - 第一棵二叉树的根节点
// root2 - 第二棵二叉树的根节点
// 返回值：
// 返回合并后的二叉树的根节点
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 如果 root1 为空，返回 root2
	if root1 == nil {
		return root2
	}
	// 如果 root2 为空，返回 root1
	if root2 == nil {
		return root1
	}

	// 创建一个新的根节点，其值为两个根节点值的和
	mergedRoot := &TreeNode{
		Val: root1.Val + root2.Val,
	}

	// 递归合并左子树
	mergedRoot.Left = mergeTrees(root1.Left, root2.Left)
	// 递归合并右子树
	mergedRoot.Right = mergeTrees(root1.Right, root2.Right)

	// 返回新的根节点
	return mergedRoot
}
