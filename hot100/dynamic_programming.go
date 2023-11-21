package leetcode_hot100

import "math"

// 70. 爬楼梯
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
		return nums[1]
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
	// 定义 and 初始化
	dp := make([]int, bagWeight+1)
	// 初始化 dp[j]表示：容量为j的背包，所背的物品价值可以最大为dp[j]
	for i := 0; i < len(weight); i++ { // 遍历物品
		// 这里必须倒序,区别二维,因为二维dp保存了i的状态
		for j := bagWeight; j >= weight[i]; j-- { // 遍历背包容量
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	//fmt.Println(dp)
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
	// 遍历物品
	for i := 1; i <= n; i++ {
		// 遍历背包
		for j := i * i; j <= n; j++ {
			// 这里一定会有所有值为1的情况，可以满足
			dp[j] = min(dp[j], dp[j-i*i]+1)
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
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}
	for i := 0; i < len(coins); i++ {
		// 遍历背包 总额数为j
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
// 输入: s = "leetcode", wordDict = ["leet", "code"]
// 输出: true
// 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
// 感觉不是背包问题。
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	// 字符长度i，能被单词拆分
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// 300. 最长递增子序列
// 输入：nums = [10,9,2,5,3,7,101,18]
// 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	ans := dp[0]
	// dp[i]表示i之前包括i的以nums[i]结尾的最长递增子序列的长度
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
i = 1
ans = 0
