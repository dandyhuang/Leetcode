package main

import (
	"fmt"
	"math"
)

// 给定一个数n，如23121，给定一组数字a,如｛2，4，9｝，求a中元素组成的小于n的最大数，如小于23121的最大数为22999
func main1() {
	fmt.Println(1^1, 1^1^1, 0^0, 0^0^1)
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxScore(grid [][]int) int {
	ans := math.MinInt
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)
	dp[0] = make([]int, n+1)
	for j := range dp[0] {
		dp[0][j] = math.MaxInt
	}

	for i := 0; i < m; i++ {
		dp[i+1] = make([]int, n+1)
		dp[i+1][0] = math.MaxInt
		for j := 0; j < n; j++ {
			last := min(dp[i+1][j], dp[i][j+1])
			ans = max(ans, grid[i][j]-last)
			dp[i+1][j+1] = min(grid[i][j], last)
		}
	}
	return ans
}

func maxScore1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	res := math.MinInt
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		res = max(res, grid[0][i]-dp[0][i-1])
		dp[0][i] = min(dp[0][i-1], grid[0][i])
	}
	for i := 1; i < m; i++ {
		res = max(res, grid[i][0]-dp[i-1][0])
		dp[i][0] = min(dp[i-1][0], grid[0][i])
		for j := 1; j < n; j++ {
			tmp := min(dp[i-1][j], dp[i][j-1])
			res = max(res, grid[i][j]-tmp)
			dp[i][j] = min(tmp, grid[i][j])
		}
	}
	return res
}
