package Tree

// buildTree 根据中序遍历和后序遍历构建二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// 创建一个 map 来存储中序遍历中每个值的索引
	inorderIndexMap := make(map[int]int)
	for i, val := range inorder {
		inorderIndexMap[val] = i
	}

	// 使用后序遍历的最后一个元素初始化根节点
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	queue := []*TreeNode{root}
	postIndex := len(postorder) - 2

	// 反向遍历后序遍历数组来构建树
	for postIndex >= 0 {
		// 从队列中取出当前节点
		node := queue[0]
		queue = queue[1:]

		// 获取当前节点在中序遍历中的索引
		inIndex := inorderIndexMap[node.Val]

		// 获取当前后序遍历的值
		postVal := postorder[postIndex]
		postIndex--

		// 判断下一个节点是右子节点还是左子节点
		if inorderIndexMap[postVal] > inIndex {
			// 如果在中序遍历中，postVal 的索引大于当前节点的索引，则为右子节点
			node.Right = &TreeNode{Val: postVal}
			// 将右子节点加入队列
			queue = append(queue, node.Right)
		} else {
			// 如果在中序遍历中，postVal 的索引小于当前节点的索引，则为左子节点
			var parent *TreeNode
			// 从队列中找到合适的父节点
			for len(queue) > 0 && inorderIndexMap[postVal] < inorderIndexMap[queue[0].Val] {
				parent = queue[0]
				queue = queue[1:]
			}
			// 将左子节点加入父节点
			parent.Left = &TreeNode{Val: postVal}
			// 将左子节点加入队列
			queue = append(queue, parent.Left)
		}
	}
	return root
}
