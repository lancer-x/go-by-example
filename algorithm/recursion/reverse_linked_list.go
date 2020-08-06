package main

type ListNode struct {
	Next *ListNode
	Value int
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	newHead := reverseList(next)
	newHead.Next = head
	return newHead
}