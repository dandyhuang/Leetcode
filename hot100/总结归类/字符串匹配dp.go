package 总结归类

// 10. 正则表达式匹配
// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
// 输入：s = "aa", p = "a"
// 输出：false
// 解释："a" 无法匹配 "aa" 整个字符串。
func isMatch(s string, p string) bool {

}

// 44. 通配符匹配
// 给你一个输入字符串 (s) 和一个字符模式 (p) ，请你实现一个支持 '?' 和 '*' 匹配规则的通配符匹配：
//'?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符序列（包括空字符序列）
// 输入：s = "cb", p = "?a"
//输出：false
//解释：'?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。

func isMatchV2(s string, p string) bool {

}

// 115. 不同的子序列
// 输入：s = "rabbbit", t = "rabbit"
// 输出：3
// 解释：
// 如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
// rabbbit
// rabbbit
// rabbbit
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	// 以i-1为结尾的s子序列中出现以j-1为结尾的t的个数为dp[i][j]。
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		// t为空
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				// i-1和j-1字符相同,需要在看i-2和j-2的字符是否也相同。即dp[i-1][j-1]有多少种
				// 这里计算的是有多少种，包括匹配和不匹配的时候，不是计算匹配了多少个字符
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}

// 583. 两个字符串的删除操作
// 给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。每步 可以删除任意一个字符串中的一个字符。
// 输入: word1 = "sea", word2 = "eat"
// 输出: 2
// 解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
func minDistanceDelete(word1 string, word2 string) int {

}

// 72. 编辑距离
// 你可以对一个单词进行如下三种操作：
// 插入一个字符
// 删除一个字符
// 替换一个字符
// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
func minDistance(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	// dp[i][j] 表示以下标i-1为结尾的字符串word1，和以下标j-1为结尾的字符串word2，最近编辑距离为dp[i][j]。
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {

			}
		}
	}
	return dp[m][n]
}

// 1143. 最长公共子序列
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace" ，它的长度为 3 。
func longestCommonSubsequence(text1 string, text2 string) int {

}
