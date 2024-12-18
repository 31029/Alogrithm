package graph

// leetcode - 200: 岛屿数量 [https://leetcode.cn/problems/number-of-islands/description/]
/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
*/

var dirction [4][2]int = [4][2]int{[2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0}, [2]int{-1, 0}}

func numIslandsdfs(x, y int, grid [][]byte, visited [][]int) {
	if grid[y][x] == '0' || visited[y][x] == 1 {
		return
	}
	visited[y][x] = 1

	for i := 0; i < 4; i++ {
		nextx := x + dirction[i][0]
		nexty := y + dirction[i][1]
		if nextx > len(grid[0])-1 || nextx < 0 || nexty > len(grid)-1 || nexty < 0 {
			continue
		}
		numIslandsdfs(nextx, nexty, grid, visited)
	}
}

func numIslandsbfs(x, y int, grid [][]byte, visited [][]int) {
	path := [][]int{[]int{x, y}}
	for len(path) > 0 {
		cur := (path)[0]
		curx := cur[0]
		cury := cur[1]
		path = path[1:len(path)]

		for i := 0; i < 4; i++ {
			nextx := curx + dirction[i][0]
			nexty := cury + dirction[i][1]
			if nextx > len(grid[0])-1 || nextx < 0 || nexty > len(grid)-1 || nexty < 0 {
				continue
			}
			if grid[nexty][nextx] == '0' || visited[nexty][nextx] == 1 {
				continue
			}
			path = append(path, []int{nextx, nexty})
			// 注意：加入队列则visited被标记为1，否则会导致超时或超出内存限制
			visited[nexty][nextx] = 1
		}
	}
}

func numIslands(grid [][]byte) int {
	count := 0
	visited := [][]int{}

	for i := 0; i < len(grid); i++ {
		visited = append(visited, make([]int, len(grid[0])))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && visited[i][j] != 1 {
				count++
				numIslandsdfs(j, i, grid, visited)
			}

		}
	}
	return count
}
