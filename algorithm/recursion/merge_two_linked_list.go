package main

//import "hash"

type ListNode1 struct {
	Val int
	Next *ListNode1
}

func mergeTwoLists(l1 *ListNode1, l2 *ListNode1) *ListNode1 {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var nextHead *ListNode1

	var head *ListNode1

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