package ds

type node struct {
	value    int
	next     *node
	previous *node
}

type bstNode struct {
	value int
	left  *bstNode
	right *bstNode
}
