package binarytree

func findBottomLeftValue(root *TreeNode) int {
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