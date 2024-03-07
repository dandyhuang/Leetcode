package hot100_2

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 70. 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 118. 杨辉三角
// 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		arr := make([]int, i+1)
		arr[0], arr[i] = 1, 1
		for j := 1; j < i; j++ {
			arr[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = arr
	}
	return res
}

// 198. 打家劫舍
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[1], nums[0])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
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
// 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，
// 其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

// 输入：n = 12
// 输出：3
// 解释：12 = 4 + 4 + 4
// 完全平方数就是物品（可以无限件使用），凑个正整数n就是背包，问凑满这个背包最少有多少物品？
func numSquares(n int) int {
	// dp[i] 表示数字 i 的最小完全平方数数量
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
	}
	for i := 1; i < n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
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
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			// 只有dp[j-coins[i]]不是初始最大值时，该位才有选择的必要
			if dp[j-coins[i]] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}

// 139. 单词拆分
// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
// 输入: s = "applepenapple", wordDict = ["apple", "pen"]
// 输出: true
// 解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
// 注意，你可以重复使用字典中的单词。
// 感觉不是背包问题。
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s))
	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	dp[0] = true
	for i := range wordDict {
		for j := 0; j < len(s); j++ {
			if j-len(wordDict) >= 0 && wordMap[s[j-len(wordDict):j]] {
				dp[j] = dp[j] || dp[j-len(wordDict[i])]
			}
		}
	}
	return dp[len(s)-1]
}

func canBreak(start int, s string, wordMap map[string]bool) bool {
	if start == len(s) {
		return true
	}
	for i := start + 1; i <= len(s); i++ {
		prefix := s[start:i]
		if wordMap[prefix] && canBreak(i, s, wordMap) {
			return true
		}
	}
	return false
}

func wordBreakV2(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}

	var dfs func(start int) bool
	dfs = func(start int) bool {
		if start == len(s) {
			return true
		}
		for i := start + 1; i <= len(s); i++ {
			if !wordMap[s[start:i]] {
				continue
			}
			if dfs(i) {
				return true
			}
		}
		return false
	}

	return dfs(0)
}

// 3. 无重复字符的最长子串
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
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
	n := len(s)
	dp := make([]int, n)
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] > 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		res = max(res, dp[i])
	}
	return res
}
