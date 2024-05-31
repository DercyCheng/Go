package _7_linkedlist

import "fmt"

// 单链表节点
type ListNode struct {
	Next *ListNode
	Val  interface{}
}

// 单链表
type LinkedList struct {
	head *ListNode
}

// 打印链表
func (l *LinkedList) Print() {
	cur := l.head.Next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.Val)
		cur = cur.Next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

/*
单链表反转
时间复杂度：O(N)
*/
func (l *LinkedList) Reverse() {
	if nil == l.head || nil == l.head.Next || nil == l.head.Next.Next {
		return
	}

	var pre *ListNode = nil
	cur := l.head.Next
	for nil != cur {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	l.head.Next = pre
}

/*
判断单链表是否有环
*/
func (l *LinkedList) HasCycle() bool {
	if nil != l.head {
		slow := l.head
		fast := l.head
		for nil != fast && nil != fast.Next {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

/*
两个有序单链表合并
*/
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if nil == l1 || nil == l1.head || nil == l1.head.Next {
		return l2
	}
	if nil == l2 || nil == l2.head || nil == l2.head.Next {
		return l1
	}

	l := &LinkedList{head: &ListNode{}}
	cur := l.head
	curl1 := l1.head.Next
	curl2 := l2.head.Next
	for nil != curl1 && nil != curl2 {
		if curl1.Val.(int) > curl2.Val.(int) {
			cur.Next = curl2
			curl2 = curl2.Next
		} else {
			cur.Next = curl1
			curl1 = curl1.Next
		}
		cur = cur.Next
	}

	if nil != curl1 {
		cur.Next = curl1
	} else if nil != curl2 {
		cur.Next = curl2
	}

	return l
}

/*
删除倒数第N个节点
*/
func (l *LinkedList) DeleteBottomN(n int) {
	if n <= 0 || nil == l.head || nil == l.head.Next {
		return
	}

	fast := l.head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.Next
	}

	if nil == fast {
		return
	}

	slow := l.head
	for nil != fast.Next {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
}

/*
获取中间节点
*/
func (l *LinkedList) FindMiddleNode() *ListNode {
	if nil == l.head || nil == l.head.Next {
		return nil
	}
	if nil == l.head.Next.Next {
		return l.head.Next
	}

	slow, fast := l.head, l.head
	for nil != fast && nil != fast.Next {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
