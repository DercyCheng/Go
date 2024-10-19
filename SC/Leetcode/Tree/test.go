package Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func _mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	mergeRoot := &TreeNode{Val: root1.Val + root2.Val}
	mergeRoot.Left = mergeTrees(root1.Left, root2.Left)
	mergeRoot.Right = mergeTrees(root1.Right, root2.Right)
	return mergeRoot
}
func _searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		_searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}
