package main

import "hash"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var nextHead *ListNode

	var head *ListNode

	if l1.Val < l2.Val {
		head = l1
		nextHead = mergeTwoLists(l1.Next, l2)

	} else {
		head = l2
		nextHead = mergeTwoLists(l1, l2.Next)

	}
	head.Next = nextHead
	return head
}