package _7_linkedlist

import "testing"

var l *LinkedList

func init() {
	n5 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	l = &LinkedList{head: &ListNode{Next: n1}}
}

func TestReverse(t *testing.T) {
	l.Print()
	l.Reverse()
	l.Print()
}

func TestHasCycle(t *testing.T) {
	t.Log(l.HasCycle())
	l.head.Next.Next.Next.Next.Next.Next = l.head.Next.Next.Next
	t.Log(l.HasCycle())
}

func TestMergeSortedList(t *testing.T) {
	n5 := &ListNode{Val: 9}
	n4 := &ListNode{Val: 7, Next: n5}
	n3 := &ListNode{Val: 5, Next: n4}
	n2 := &ListNode{Val: 3, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	l1 := &LinkedList{head: &ListNode{Next: n1}}

	n10 := &ListNode{Val: 10}
	n9 := &ListNode{Val: 8, Next: n10}
	n8 := &ListNode{Val: 6, Next: n9}
	n7 := &ListNode{Val: 4, Next: n8}
	n6 := &ListNode{Val: 2, Next: n7}
	l2 := &LinkedList{head: &ListNode{Next: n6}}

	MergeSortedList(l1, l2).Print()
}

func TestDeleteBottomN(t *testing.T) {
	l.Print()
	l.DeleteBottomN(3)
	l.Print()
}

func TestFindMiddleNode(t *testing.T) {
	l.DeleteBottomN(1)
	l.DeleteBottomN(1)
	l.DeleteBottomN(1)
	l.DeleteBottomN(1)
	l.Print()
	t.Log(l.FindMiddleNode())
}
