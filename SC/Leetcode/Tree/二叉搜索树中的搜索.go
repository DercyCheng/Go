package Tree

// searchBST 在二叉搜索树中查找值为 val 的节点
// 参数：
// root - 二叉搜索树的根节点
// val - 要查找的值
// 返回值：
// 返回值为 val 的节点，如果不存在则返回 nil
func searchBST(root *TreeNode, val int) *TreeNode {
	// 如果根节点为空或根节点的值等于要查找的值，返回根节点
	if root == nil || root.Val == val {
		return root
	}
	// 如果根节点的值大于要查找的值，在左子树中递归查找
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	// 如果根节点的值小于要查找的值，在右子树中递归查找
	return searchBST(root.Right, val)
}
