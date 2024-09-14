package binarytree

import "testing"

func TestTraversal(t *testing.T)  {
	nums := []any{3, 9, 20, nil, nil, 15, 7}
	root := BuildTreeByLevelOrder(nums)
	LevelOrder(root)
}