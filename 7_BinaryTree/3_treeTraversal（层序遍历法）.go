package binarytree

import "container/list"


func LevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	q := list.New()
	q.PushBack(root)

	layerSize := 1
	for q.Len() > 0 {
		nums := []int{}
		for i := 0; i < layerSize && q.Len()>0; i++ {
			node , ok := q.Remove(q.Front()).(*TreeNode)
			if !ok {
				q.PushBack(nil)
			}
			if node != nil {
				nums = append(nums, node.Val)
				q.PushBack(node.Left)
				q.PushBack(node.Right)
			}
		}

		res = append(res, nums)
		layerSize = layerSize * 2
	}
	return nil
}

func LevelOrderWithEmpty(root *TreeNode) [][]interface{} {
	res := [][]interface{}{}
	if root == nil {
		return res
	}

	q := list.New()
	q.PushBack(root)

	layerSize := 1
	for q.Len() > 0 {
		nums := []interface{}{}
		for i := 0; i < layerSize && q.Len()>0; i++ {
			node , ok := q.Remove(q.Front()).(*TreeNode)
			if !ok {
				nums = append(nums, nil)
				q.PushBack(nil)
			}
			if node != nil {
				nums = append(nums, node.Val)
				q.PushBack(node.Left)
				q.PushBack(node.Right)
			}
		}

		res = append(res, nums)
		layerSize = layerSize * 2
	}
	return nil
}