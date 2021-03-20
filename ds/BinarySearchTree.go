package ds

type BinarySearchTree struct {
	root  *bstNode
	left  *bstNode
	right *bstNode
}

func NewBinarySearchTree(v int) *BinarySearchTree {
	root := bstNode{value: v}

	return &BinarySearchTree{root: &root}
}

func (b *BinarySearchTree) Insert(v int) *BinarySearchTree {
	if b.root == nil {
		return NewBinarySearchTree(v)
	}
	currentNode := b.root
	for true {
		if v < currentNode.value {
			if currentNode.left == nil {
				currentNode.left = &bstNode{value: v}
				return b
			}
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				currentNode.right = &bstNode{value: v}
				return b
			}
			currentNode = currentNode.right
		}
	}

	return b
}

func (b *BinarySearchTree) Lookup(v int) bool {
	currentNode := b.root
	if currentNode == nil {
		return false
	}
	for true {
		if v == currentNode.value {
			return true
		}
		if v < currentNode.value {
			if currentNode.left == nil {
				return false
			}
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				return false
			}
			currentNode = currentNode.right
		}
	}

	return false
}
