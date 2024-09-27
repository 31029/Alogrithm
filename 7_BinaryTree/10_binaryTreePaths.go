package binarytree

import (
	"fmt"
	"strconv"
)

func binaryTreePathsRecursion(root *TreeNode, cur string, paths *[]string) {
	if root == nil {
		*paths = append(*paths, cur)
		return
	}

	cur = fmt.Sprintf("%s->", strconv.Itoa(root.Val)) 
	binaryTreePathsRecursion(root.Left, cur, paths)
	binaryTreePathsRecursion(root.Right, cur, paths)
}

func binaryTreePaths(root *TreeNode) []string {
	cur := ""
	res := []string{}
	binaryTreePathsRecursion(root, cur, &res)
	return res
}