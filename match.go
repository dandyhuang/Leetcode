package main

import (
	"fmt"
)

// isMatch 判断字符串是否匹配通配符
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	// dp[i][j] 表示 s 的前 i 个字符和 p 的前 j 个字符是否匹配
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// 空字符串和空模式匹配
	dp[0][0] = true

	// dp[0][j] 需要分情况讨论：
	// 因为星号才能匹配空字符串，所以只有当模式 p的前 j 个字符均为星号时，dp[0][j] 才为真。
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		} else {
			break
		}
	}

	// 动态规划填表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				// '*' 匹配 0 次或多次的情况
				// dp[i-1][j] 匹配多次情况
				// dp[i][j-1] 删除，不匹配
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				// '?' 匹配单个字符的情况，或字符相等的情况
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}

func main() {
	// 示例字符串和模式
	s := "aa"
	p := "*"

	// 判断是否匹配
	result := isMatch(s, p)

	// 输出结果
	fmt.Printf("Is '%s' matched with pattern '%s'? %v\n", s, p, result)
}
