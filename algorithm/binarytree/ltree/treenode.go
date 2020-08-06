package ltree



type BinaryTree struct {
	Value int
	LeftChild *BinaryTree
	RightChild *BinaryTree
}



func init()  {

}

func Btree(level int) *BinaryTree {
	if level == 1 {
		return &BinaryTree{
			Value: level,
			LeftChild: nil,
			RightChild: nil,
		}
	}

	root := BinaryTree{
		Value:      level,
		LeftChild:  Btree(level - 1),
		RightChild: Btree(level - 1),
	}

	return &root
}