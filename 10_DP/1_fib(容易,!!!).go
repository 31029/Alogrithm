package dp

func fib(n int) int {
    if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fib2(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = 1
	for i := 0; i < n; i++ {
		// 优化：只需要用到数组的两个连续元素空间
		sum := dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = sum
	}
	return dp[1]
}