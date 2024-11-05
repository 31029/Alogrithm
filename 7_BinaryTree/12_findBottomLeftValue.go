package binarytree

// leetcode - 513: 找二叉树左下角的值
// 描述： 给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
// 假设二叉树中至少有一个节点。
func FindBottomLeftValue(root *TreeNode) int {
	// 思路1：通过层序遍历法，找到最后一层第一个不为空的节点的值。
	res := 0
	tN := LevelOrderWithEmpty(root)
	for _, v := range tN[len(tN)-1] {
		v , ok := v.(int)
		if ok {
			res = v
			return res
		}
	}
	return res
}