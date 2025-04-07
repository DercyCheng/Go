package Tree

// buildTree 根据中序遍历和后序遍历构建二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	var inorderIndex int
	for i, v := range inorder {
		if v == rootVal {
			inorderIndex = i
			break
		}
	}
	buildTree(inorder[:inorderIndex], postorder[:inorderIndex])
	buildTree(inorder[inorderIndex+1:], postorder[inorderIndex:len(postorder)-1])
	return root
}
