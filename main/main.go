package main

import (
	"ds.com/ds/ds"
	"fmt"
)

func main() {
	bst := ds.NewBinarySearchTree(20)
	bst.Insert(45)
	bst.Insert(11)
	bst.Insert(1)
	bst.Insert(6)
	bst.Insert(111)
	bst.Insert(19)
	bst.Insert(26)

	treeContains := bst.Lookup(26)
	fmt.Println(treeContains)

	bst.Remove(26)
	treeContains = bst.Lookup(26)
	fmt.Println(treeContains)

}
