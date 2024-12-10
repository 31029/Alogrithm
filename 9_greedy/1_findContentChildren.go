package greedy

// leetcode - 455:分发饼干  [https://leetcode.cn/problems/assign-cookies/description/]
/*
假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；
并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。
你的目标是满足尽可能多的孩子，并输出这个最大数值。
*/
func bubbleSort(tgt []int) {
	// 优化过后的bubbleSort
	for i, change := 0, true; i < len(tgt) && change; i++ {
		change = false
		for j := 0; j < len(tgt) - 1 - i; j++ {
			if tgt[j] > tgt[j+1] {
				tmp := tgt[j]
				tgt[j] = tgt[j+1]
				tgt[j+1] = tmp
				change = true
			}
		}
	}
}


func findContentChildren(g []int, s []int)(sat int){
	bubbleSort(g)
	bubbleSort(s)

	gI, sI := 0, 0
	for gI < len(g) && sI < len(s) {
		if s[sI] >= g[gI] {
			sat ++
			gI ++
			sI ++
		} else {
			sI ++
		}
	}
	return
}