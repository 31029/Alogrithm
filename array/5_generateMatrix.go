package array

func GenerateMatrix(n int) [][]int {
	// 解题思路关键：用循环模拟转圈的过程 + 循环不变量原则：每一条边是左闭右开 [a, b)？还是左闭右闭 [a, b]？
	  
	// 生成二维切片，并初始化为0
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	odd := false
	circule := 0
	if n%2 == 0 {
		circule = n / 2
	} else {
		circule = (n - 1) / 2
		odd = true
	}

	// 特殊值处理
	if n == 1 {
		matrix[0][0] = 1
	}

	current_num := 1
	for i := 1; i <= circule; i++ {
		line_num := n - 1 - 2*(i-1) // 每次循环后，一条边上带填充的num的个数

		for l1 := 0; l1 < line_num; l1++ {
			matrix[i-1][l1+i-1] = current_num
			current_num++
		}

		for l2 := 0; l2 < line_num; l2++ {
			matrix[l2+i-1][n-i] = current_num
			current_num++
		}

		for l3 := line_num; l3 > 0; l3-- {
			matrix[n-i][l3+i-1] = current_num
			current_num++
		}

		for l4 := line_num; l4 > 0; l4-- {
			matrix[l4+i-1][i-1] = current_num
			current_num++
		}
	}

	if odd {
		matrix[n/2][n/2] = n * n
	}

	return matrix
}
