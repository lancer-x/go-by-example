package main

import (
	"fmt"
	"got/algorithm/binarytree/ltree"
)

func main()  {
	root := ltree.Btree(5)
	fmt.Println(*root)
	//preorderTraversal(root)
	levelOrder(root)
}

//前序bian

func preorderTraversal(root *ltree.BinaryTree) {
	if root == nil {
		return
	}
	fmt.Println(root.Value)
	preorderTraversal(root.LeftChild)
	preorderTraversal(root.RightChild)
}

//层序遍历
func levelOrder(root *ltree.BinaryTree) [][]int{
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*ltree.BinaryTree, 0)

	//其实根节点入队就是第一层
	queue = append(queue, root)
	//队列长度是变化的，需要实时计算
	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		//此处for循环表示一次处理一层的数据，因为l已经固定了
		for i := 0; i< l; i++ {
			//获取队列头部数据
			queueHead := queue[0]
			//出队
			queue = queue[1:]

			if queueHead.LeftChild != nil {
				queue = append(queue, queueHead.LeftChild)
			}
			if queueHead.RightChild != nil {
				queue = append(queue, queueHead.RightChild)
			}
			list = append(list, queueHead.Value)
		}
		//将每层的数据加入二维数组
		result = append(result, list)
	}

	for level,data := range result {
		fmt.Printf("level%d", level + 1)
		for _,v := range data {
			fmt.Printf(" %d ", v)
		}
		fmt.Println()
	}

	return result
}

//深度搜索 从上到下  先处理根节点
//注意此处第二个参数，因为是递归调用然后放到一个数组里，需要设置一个共同变量，作为参数传递
func DFSupdown(root *ltree.BinaryTree, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Value)
	//处理左节点
	DFSupdown(root.LeftChild, result)
	//处理右节点
	DFSupdown(root.RightChild, result)
}

//DFS深度搜索  从下到上 用分治
func divideAndConquer(root *ltree.BinaryTree) []int{
	result := make([]int, 0)

	if root == nil {
		return result
	}
	//分治
	left := divideAndConquer(root.LeftChild)
	right := divideAndConquer(root.RightChild)

	//合并结果
	result = append(result, root.Value)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

///// DFS 深度搜索（从上到下） 和分治法区别：前者一般将最终结果通过指针参数传入，后者一般递归返回结果最后合并