package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return helper(root)
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left)
	right := helper(root.Right)
	return 1 + max(left, right)
}

func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

