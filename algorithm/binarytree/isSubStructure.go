package main


func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return findBRoot(A, B)
}




func findBRoot(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}

	return helper3(A, B) || findBRoot(A.Left, B) || findBRoot(A.Right, B)

}

func helper3(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}

	return A.Val == B.Val && helper3(A.Left, B.Left) && helper3(A.Right, B.Right)
}
