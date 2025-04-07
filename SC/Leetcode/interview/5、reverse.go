package interview

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// 创建虚拟头节点，处理m=1的情况
	dummy := &ListNode{Next: head}
	prev := dummy

	// 定位到反转区间的前驱节点（走m-1步）
	for i := 0; i < m-1; i++ {
		prev = prev.Next
	}

	// start指向待反转区间的第一个节点
	start := prev.Next

	// 头插法反转区间，共n-m次操作
	for i := 0; i < n-m; i++ {
		// 暂存start的下一个节点（要移动的节点）
		removed := start.Next
		// 将start的next指向下一个节点的next，跳过被移动的节点
		start.Next = removed.Next
		// 将移动的节点插入到prev的后方
		removed.Next = prev.Next
		prev.Next = removed
	}

	return dummy.Next
}
