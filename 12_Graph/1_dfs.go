package graph

import (
	"fmt"
)

var result [][]int // 收集符合条件的路径
var path []int     // 1节点到终点的路径

// karma - 98:所有可达路径  [https://kamacoder.com/problempage.php?pid=1170]
/*
给定一个有 n 个节点的有向无环图，节点编号从 1 到 n。请编写一个函数，找出并返回所有从节点 1 到节点 n 的路径。每条路径应以节点编号的列表形式表示。

【输入描述】
第一行包含两个整数 N，M，表示图中拥有 N 个节点，M 条边
后续 M 行，每行包含两个整数 s 和 t，表示图中的 s 节点与 t 节点中有一条路径
*/
func dfs(graph [][]int, x, n int) {
	if x == n {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)
		return
	}

	// 每层都有 n 个节点
	for i:= 0; i <= n; i ++ {
		if graph[x][i] == 1 {
			path = append(path, i)
			dfs(graph, i, n)
			path = path[:len(path)-1]
		}
	}
}