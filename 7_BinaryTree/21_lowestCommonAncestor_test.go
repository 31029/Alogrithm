package binarytree

import "testing"


func TestLowestCommonAncestor(t *testing.T){
	root := BuildTreeByLevelOrder([]any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5, nil, nil, nil, nil})
	p := &TreeNode{Val: 2}
	q := &TreeNode{Val: 4}
	LowestCommonAncestor(root, p , q)
}

