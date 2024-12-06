package codedaily

// leetcode:999 可以被一步捕获的棋子数 [https://leetcode.cn/problems/available-captures-for-rook/description/]
/*
给定一个 8 x 8 的棋盘，只有一个 白色的车，用字符 'R' 表示。棋盘上还可能存在白色的象 'B' 以及黑色的卒 'p'。空方块用字符 '.' 表示。
车可以按水平或竖直方向（上，下，左，右）移动任意个方格直到它遇到另一个棋子或棋盘的边界。如果它能够在一次移动中移动到棋子的方格，则能够 吃掉 棋子。

注意：车不能穿过其它棋子，比如象和卒。这意味着如果有其它棋子挡住了路径，车就不能够吃掉棋子。
返回白车将能 吃掉 的 卒的数量。

*/
func numRookCaptures(board [][]byte) int {
	// 解题思路：常规思路
	Rrow, Rcol := 0, 0

	// 找到R的坐标
	for row := range board {
		for col, v := range board[row] {
			if v == []byte("R")[0] {
				Rrow = row
				Rcol = col
			}
		}
	}

	eatCount := 0
	// 横
	for i := Rcol - 1; i >= 0; i-- {
		// 直接使用单引号即可！！！！！
		if board[Rrow][i] == 'B' {
			break
		}
		if board[Rrow][i] == 'p' {
			eatCount++
			break
		}
	}
	for i := Rcol + 1; i <= len(board[Rrow])-1; i++ {
		if board[Rrow][i] == 'B' {
			break
		}
		if board[Rrow][i] == 'p' {
			eatCount++
			break
		}
	}

	// 竖
	for i := Rrow - 1; i >= 0; i-- {
		if board[i][Rcol] == 'B' {
			break
		}
		if board[i][Rcol] == 'p' {
			eatCount++
			break
		}
	}
	for i := Rrow + 1; i <= len(board)-1; i++ {
		// 不要用这种方式，底层会发生一次拷贝？
		if board[i][Rcol] == []byte("B")[0] {
			break
		}
		if board[i][Rcol] == []byte("p")[0] {
			eatCount++
			break
		}
	}

	return eatCount

}
