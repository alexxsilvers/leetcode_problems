package binary_search_tree

import (
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	root := Insert(nil, 10)
	Insert(root, 4)
	Insert(root, 18)
	Insert(root, 45)
	Insert(root, 34)
	Insert(root, 78)
	Insert(root, 6)
	Insert(root, 8)
	Insert(root, 15)

	InOrder(root)
}
