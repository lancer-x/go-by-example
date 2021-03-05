package main

//import "got/linked"

func main1() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head.Next.Next
	slow := head.Next

	for fast != nil && slow != nil {
		if fast == slow {
			return true
		}
		if fast.Next == nil || slow.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}
