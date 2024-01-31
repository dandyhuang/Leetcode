package hot100_2

import (
	"math"
)

// 70. 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
func climbStairs(n int) int {
	if n < 2 {
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
// 输入: numRows = 5
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
// 输入：[1,2,3,1]
// 输出：4
// 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
// 偷窃到的最高金额 = 1 + 3 = 4 。
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[1], nums[0])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

// 0 1背包问题
// 有N件物品和一个最多能背重量为W的背包。第i件物品的重量是weight[i]，得到的价值是value[i] 。
// 每件物品都有无限个（也就是可以放入背包多次），求解将哪些物品装入背包里物品价值总和最大。
func test1WeiBagProblem(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)
	// 初始化 dp[j]表示：容量为j的背包，所背的物品价值可以最大为dp[j]
	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[i]; j-- { // 背包遍历
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
	// dp[j]：和为j的完全平方数的最少数量为dp[j]
	dp := make([]int, n+1)
	dp[0] = 0
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

func numSquaresV2(n int) int {
	// dp[i] 表示数字 i 的最小完全平方数数量
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		// 初始值，假设每个数字都由 1 的平方组成
		dp[i] = i

		// 尝试从 1 开始找完全平方数
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
func coinChange(coins []int, amount int) int {
	// dp[j]：凑足总额为j所需钱币的最少个数为dp[j]
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}
	for i := range coins {
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

func coinChangeV2(coins []int, amount int) int {
	// dp[i] :凑足总额为i所需钱币的最少个数为dp[i]
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		// 不定义成32，到时候+1会溢出，dp里头
		dp[i] = math.MaxInt32
		for j := range coins {
			// 大于等于没关系，只需要保证dp[i-coins[j]]获取到值
			if i-coins[j] >= 0 {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// 139. 单词拆分,  最长递增子序列类似
// 输入: s = "goalspecial",wordDict = ["go","goal","goals","special"]
// 输出: true
// 解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
// 感觉不是背包问题。
func wordBreak(s string, wordDict []string) bool {
	// dp[i] 表示字符串 s 的前 i 个字符能否被拆分
	dp := make([]bool, len(s)+1)
	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// 因为j是从0开始，除了自身，还包括前面的dp[0-j-1]也是true
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// 3. 无重复字符的最长子串
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 滑动窗口，统一放过来求最长xxx
func lengthOfLongestSubstringV2(s string) int {
	res := 0
	m := make(map[byte]bool)
	left := 0
	for i := range s {
		for m[s[i]] {
			delete(m, s[left])
			left++
		}
		m[s[i]] = true
		res = max(res, i-left+1)
	}
	return res
}

// 300. 最长递增子序列
// 输入：nums = [10,9,2,5,3,7,101,18]
// 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
func lengthOfLIS(nums []int) int {
	res := 1
	// dp[i]表示i之前包括i的以nums[i]结尾的最长递增子序列的长度
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

// 674. 最长连续递增序列
// 输入：nums = [1,3,5,4,7]
// 输出：3
// 解释：最长连续递增序列是 [1,3,5], 长度为3。
// 尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。
func findLengthOfLCIS(nums []int) int {
	res := 1
	// dp[i]表示i之前包括i的以nums[i]结尾的最长递增子序列的长度
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = max(dp[i], dp[i-1]+1)
		}
		res = max(res, dp[i])
	}
	return res
}

// 1143.最长公共子序列
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace"，它的长度为 3。
func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j]：长度为[0, i-1]的字符串text1与长度为[0, j-1]的字符串text2的最长公共子序列为dp[i][j]
	dp := make([][]int, len(text1))
	for i := range dp {
		dp[i] = make([]int, len(text2))
	}

}

// 718. 最长重复子数组
// A: [1,2,3,2,1]
// B: [3,2,1,4,7]
// 输出：3
// 解释：长度最长的公共子数组是 [3, 2, 1] 。
func findLength(nums1 []int, nums2 []int) int {

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
	// 由于存在负数，需要同时维护最大值和最小值
	dpMax := make([]int, len(nums)+1)
	dpMin := make([]int, len(nums)+1)
	res := nums[0]
	dpMax[0] = nums[0]
	dpMin[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			dpMax[i] = max(nums[i], dpMax[i-1]*nums[i])
			dpMin[i] = min(nums[i], dpMin[i-1]*nums[i])
		} else {
			dpMax[i] = max(nums[i], dpMin[i-1]*nums[i])
			dpMin[i] = min(nums[i], dpMax[i-1]*nums[i])
		}
		res = max(res, dpMax[i])
	}
	return res
}

// 416. 分割等和子集
// 输入：nums = [1,5,11,5]
// 输出：true
// 解释：数组可以分割成 [1, 5, 5] 和 [11] 。
// 使得两个子集的元素和相等。
func canPartition(nums []int) bool {
	// dp[j] 表示： 容量为j的背包，所背的物品价值最大可以为dp[j]。
	// dp[j]表示 背包总容量（所能装的总重量）是j，放进物品后，背的最大重量为dp[j]。
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 奇数则不可能平分成两个子集
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	// dp[j] 表示nums的子集和是否相等j
	// dp[j]表示 背包总容量（所能装的总重量）是j，放进物品后，背的最大重量为dp[j]。
	dp := make([]int, target+1)

	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}
	return dp[target] == target
}
func canPartitionV2(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	// 如果数组元素总和为奇数，无法分割成两个相等的子集
	if sum%2 != 0 {
		return false
	}

	// target 表示目标和的一半
	target := sum / 2

	// dp[i] 表示是否能从数组中选取一些数字，使它们的和等于 i
	dp := make([]bool, target+1)
	dp[0] = true // 空集合的和为0，是合法的

	// 遍历数组，更新 dp 数组
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}

// 32. 最长有效括号
// 输入：s = "(()(())"
// 输出：6
// 解释：最长有效括号子串是 "()()"
func longestValidParentheses(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}
	// dp[i] 表示以 s[i] 结尾的最长有效括号的长度
	dp := make([]int, n)
	res := 0
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				// 当前字符为')'且前一个字符为'('时，自身为2个字符()
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					// "(()())" 这种场景
					dp[i] = dp[i-1] + 2
				}
			}
		}
		res = max(res, dp[i])
	}
	return res
}
