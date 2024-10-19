package Tree

// IterationPreorderTraversal 进行二叉树的统一迭代前序遍历。
// 它使用栈来模拟递归过程，并返回遍历结果的整数切片。
func IterationPreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil // 如果根节点为空，返回空结果
	}
	res := []int{}             // 初始化结果切片
	stack := []*TreeNode{root} // 初始化栈，将根节点加入栈
	for len(stack) > 0 {
		node := stack[len(stack)-1]  // 从栈顶取出一个节点
		stack = stack[:len(stack)-1] // 移除已取出的节点
		res = append(res, node.Val)  // 将节点的值添加到结果中
		if node.Right != nil {
			stack = append(stack, node.Right) // 如果右子节点不为空，将其加入栈
		}
		if node.Left != nil {
			stack = append(stack, node.Left) // 如果左子节点不为空，将其加入栈
		}
	}
	return res // 返回包含前序遍历结果的切片
}

// IterationInorderTraversal 进行二叉树的统一迭代中序遍历。
// 它使用栈来模拟递归过程，并返回遍历结果的整数切片。
func IterationInorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil // 如果根节点为空，返回空结果
	}
	res := []int{}         // 初始化结果切片
	stack := []*TreeNode{} // 初始化栈
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node) // 将节点加入栈
			node = node.Left            // 处理左子节点
		}
		node = stack[len(stack)-1]   // 从栈顶取出一个节点
		stack = stack[:len(stack)-1] // 移除已取出的节点
		res = append(res, node.Val)  // 将节点的值添加到结果中
		node = node.Right            // 处理右子节点
	}
	return res // 返回包含中序遍历结果的切片
}

// IterationPostorderTraversal 进行二叉树的统一迭代后序遍历。
// 它使用栈来模拟递归过程，并返回遍历结果的整数切片。
func IterationPostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := []*TreeNode{}
	var lastVisited *TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		peekNode := stack[len(stack)-1]
		if peekNode.Right != nil && lastVisited != peekNode.Right {
			node = peekNode.Right
		} else {
			res = append(res, peekNode.Val)
			lastVisited = peekNode
			stack = stack[:len(stack)-1]
		}
	}
	return res // 返回包含后序遍历结果的切片
}
