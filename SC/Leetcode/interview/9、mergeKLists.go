package interview

// 主函数：合并K个链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	// 分治：拆分成左右两部分分别合并
	mid := len(lists) / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return mergeTwo(left, right) // 合并两个链表
}

// 辅助函数：合并两个有序链表（递归版）
func mergeTwo(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 选择较小的头节点，递归合并剩余部分
	if l1.Val < l2.Val {
		l1.Next = mergeTwo(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwo(l1, l2.Next)
		return l2
	}
}
