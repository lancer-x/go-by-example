package main 

type Node struct {
    val int //
    left *Node //
    right *Node //
}



func main() {
    root := &Node{}
    imageTree(root)
}

func imageTree(root *Node) {
    if root == nil {
        return
    }

    root.left = imageTree(root.right)
    root.right = imageTree(root.left)
}