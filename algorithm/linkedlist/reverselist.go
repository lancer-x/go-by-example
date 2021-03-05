package main


func main()  {

}

func reverse(head *Node) *Node {
	var prev *Node
	for head != nil {
		//为了断开当前节点与后一个节点的联系，将后一个节点位置记录下来
		temp := head.Next
		//断开当前节点与后一个节点
		head.Next = prev
		//prev节点后移
		prev = head
		//更新head节点
		head = temp
	}
	return prev
}
