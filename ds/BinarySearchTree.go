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

func (b *BinarySearchTree) Remove(v int) bool {
	currentNode := b.root
	if currentNode == nil {
		return false
	}

	parentNode := &bstNode{}
	for currentNode != nil {
		if v < currentNode.value {
			parentNode = currentNode
			currentNode = currentNode.left
		} else if v > currentNode.value {
			parentNode = currentNode
			currentNode = currentNode.right
		} else if v == currentNode.value {
			// Option 1: No right child
			if currentNode.right == nil {
				if parentNode == nil {
					b.root = currentNode.left
				} else {
					if currentNode.value < parentNode.value {
						parentNode.left = currentNode.left
					} else if currentNode.value > parentNode.value {
						parentNode.right = currentNode.left
					}
				}
				// Option 2: Right child that doesn't have a left child
			} else if currentNode.right.left == nil {
				if parentNode == nil {
					b.root = currentNode.left
				} else {
					currentNode.right.left = currentNode.left
					if currentNode.value < parentNode.value {
						parentNode.left = currentNode.right
					} else if currentNode.value > parentNode.value {
						parentNode.right = currentNode.right
					}
				}
				// Option 3: right child that has a left child
			} else {
				leftmost := currentNode.right.left
				leftmostParent := currentNode.right
				for leftmost.left != nil {
					leftmostParent = leftmost
					leftmost = leftmost.left
				}

				leftmostParent.left = leftmost.right
				leftmost.left = currentNode.left
				leftmost.right = currentNode.right

				if parentNode == nil {
					b.root = leftmost
				} else {
					if currentNode.value < parentNode.value {
						parentNode.left = leftmost
					} else if currentNode.value > parentNode.value {
						parentNode.right = leftmost
					}
				}
			}
			return true
		}
	}

	return false
}
