package binarytree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历
func BuildTreeByLevelOrder(arr []any) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	// 根节点
	var root *TreeNode
	if arr[0] != nil {
		root = &TreeNode{Val: arr[0].(int)}
	} else {
		return nil
	}

	// 用队列进行层序遍历构造二叉树
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(arr) {
		// 模拟队列pop操作
		current := queue[0]
		queue = queue[1:]

		// 处理左子节点
		if i < len(arr) && arr[i] != nil {
			current.Left = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, current.Left)
		}
		i++

		// 处理右子节点
		if i < len(arr) && arr[i] != nil {
			current.Right = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, current.Right)
		}
		i++
	}

	return root
}

// 中序遍历，验证生成的树
func InorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	InorderTraversal(root.Left)
	fmt.Printf("%d ", root.Val)
	InorderTraversal(root.Right)
}
