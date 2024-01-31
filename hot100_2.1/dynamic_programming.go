package hot100_2

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 70. 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
func climbStairs(n int) int {

}

// 118. 杨辉三角
// 输入: numRows = 5
// 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
func generate(numRows int) [][]int {

}

// 198. 打家劫舍
func rob(nums []int) int {

}

// 0 1背包问题
func test1WeiBagProblem(weight, value []int, bagWeight int) int {
	//  初始化 dp[j]表示：容量为j的背包，所背的物品价值可以最大为dp[j]
	dp := make([]int, bagWeight+1)
	for i := range weight {
		for j := bagWeight; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[bagWeight]
}

// 0 1背包问题
func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	test1WeiBagProblem(weight, value, 4)
}

// 279. 完全平方数
// 输入：n = 12
// 输出：3
// 解释：12 = 4 + 4 + 4
// 完全平方数就是物品（可以无限件使用），凑个正整数n就是背包，问凑满这个背包最少有多少物品？
func numSquares(n int) int {
	// dp[i] 表示数字 i 的最小完全平方数数量
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

// 322. 零钱兑换
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
// 最少的硬币个数
// amount = 3 , coins = [2]
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range coins {
		for j := amount; j > 0; j-- {
			dp[j] = min(dp[j], dp[j-1]+1)
		}
	}
}

// 139. 单词拆分
// 输入: s = "leetcode", wordDict = ["leet", "code"]
// 输出: true
// 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
// 感觉不是背包问题。
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	// dp[i] 表示字符串 s 的前 i 个字符能否被拆分
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {

	}
	return dp[len(s)]
}

// 3. 无重复字符的最长子串
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 滑动窗口，统一放过来求最长xxx
func lengthOfLongestSubstringV2(s string) int {
}

// 300. 最长递增子序列
// 输入：nums = [10,9,2,5,3,7,101,18]
// 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4
func lengthOfLIS(nums []int) int {

}

// 674. 最长连续递增序列
// 输入：nums = [1,3,5,4,7]
// 输出：3
// 解释：最长连续递增序列是 [1,3,5], 长度为3。尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。
func findLengthOfLCIS(nums []int) int {

}

// 718. 最长重复子数组
// A: [1,2,3,2,1]
// B: [3,2,1,4,7]
// 输出：3
// 解释：长度最长的公共子数组是 [3, 2, 1] 。
func findLength(nums1 []int, nums2 []int) int {

}

// 1143.最长公共子序列
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace"，它的长度为 3。
func longestCommonSubsequence(text1 string, text2 string) int {

}

// 53. 最大子序和
// 输入: [-2,1,-3,4,-1,2,1,-5,4]
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
func maxSubArrayV1(nums []int) int {

}

// 152. 乘积最大子数组
// 输入: nums = [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。
func maxProduct(nums []int) int {

}

// 416. 分割等和子集
// 输入：nums = [1,5,11,5]
// 输出：true
// 解释：数组可以分割成 [1, 5, 5] 和 [11] 。
func canPartition(nums []int) bool {
	// dp[j] 表示： 容量为j的背包，所背的物品价值最大可以为dp[j]。
	// dp[j]表示 背包总容量（所能装的总重量）是j，放进物品后，背的最大重量为dp[j]。
}

// 32. 最长有效括号
// 输入：s = "(()(())"
// 输出：6
// 解释：最长有效括号子串是 "()()"
func longestValidParentheses(s string) int {

}
