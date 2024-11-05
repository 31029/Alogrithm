package binarytree


// LeetCode-501: 二叉搜索树中的众数. [https://leetcode.cn/problems/find-mode-in-binary-search-tree/description/]
/* 描述：
给你一个含重复值的二叉搜索树（BST）的根节点 root ，找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。
如果树中有不止一个众数，可以按 任意顺序 返回。
*/
func findMode(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	hmap := make(map[int]int)
	findModeRecursion(root, hmap)

	most := 0
	for _, v := range hmap {
		if v > most {
			most = v
		}
	}

	for k, v := range hmap {
		if v == most {
			res = append(res, k)
		}
	}

	return res
}


func findModeRecursion(root *TreeNode, hmap map[int]int)  {
	if root == nil {
		return 
	}

	hmap[root.Val] += 1
	findModeRecursion(root.Left, hmap)
	findModeRecursion(root.Right, hmap)
}