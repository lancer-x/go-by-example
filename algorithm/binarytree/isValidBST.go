package main

import "math"

type res struct {
	Flag bool
	MaxLeftVal int
	MinRightVal int
}

func main1()  {

}

func isValidBST(root *TreeNode) bool {
	return helper1(root, math.MinInt64, math.MaxInt64)
}

func helper1(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val < min || root.Val > max {
		return false
	}

	if helper1(root.Left, min, root.Val) && helper1(root.Right, root.Val, max) {
		return true
	}
	return false
}