package main


func levelOrder1(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	tmp := make([]int, 0)
	for len(queue) > 0 {
		length := len(queue)
		tmp = tmp[0:0]
		for i:=0; i < length; i++ {
			node := queue[i]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[length:]
		res = append(res, tmp)
	}
	return res
}