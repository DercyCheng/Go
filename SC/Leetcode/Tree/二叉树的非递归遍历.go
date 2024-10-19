package Tree

import "container/list"

func literate_preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		node := l.Remove(l.Back()).(*TreeNode)
		ans = append(ans, node.Val)
		if node.Right != nil {
			l.PushBack(node.Right)
		}
		if node.Right != nil {
			l.PushBack(node.Left)
		}
	}
	return ans
}
func literate_postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		node := l.Remove(l.Back()).(*TreeNode)
		ans = append(ans, node.Val)
		if node.Right != nil {
			l.PushBack(node.Left)
		}
		if node.Right != nil {
			l.PushBack(node.Right)
		}

	}
	reverse := func(a []int) {
		l, r := 0, len(a)-1
		for l < r {
			a[l], a[r] = a[r], a[l]
			l, r = l+1, r-1
		}
	}
	reverse(ans)
	return ans
}

func literate_inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	l := list.New()
	cur := root
	for cur != nil || l.Len() > 0 {
		if cur != nil {
			l.PushBack(cur)
			cur = cur.Left
		} else {
			cur = l.Remove(l.Back()).(*TreeNode)
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}
	return ans
}
