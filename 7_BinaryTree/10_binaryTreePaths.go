package binarytree

import (
	"fmt"
	"strconv"
)

func BinaryTreePathsRecursion(root *TreeNode, cur string, paths *[]string) {
	if root == nil {
		*paths = append(*paths, cur)
		return
	}

	cur = fmt.Sprintf("%s->", strconv.Itoa(root.Val)) 
	BinaryTreePathsRecursion(root.Left, cur, paths)
	BinaryTreePathsRecursion(root.Right, cur, paths)
}

func BinaryTreePaths(root *TreeNode) []string {
	cur := ""
	res := []string{}
	BinaryTreePathsRecursion(root, cur, &res)
	return res
}