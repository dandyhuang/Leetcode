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
	// 定义 and 初始化， 一维dp数组（滚动数组）
	dp := make([]int, bagWeight+1)
	// 初始化 dp[j]表示：容量为j的背包，所背的物品价值可以最大为dp[j]
	for i := 0; i < len(weight); i++ { // 遍历物品
		// 这里必须倒序,区别二维,因为二维dp保存了i的状态
		// 倒序遍历是为了保证物品i只被放入一次！。但如果一旦正序遍历了，那么物品0就会被重复加入多次！
		for j := bagWeight; j >= weight[i]; j-- { // 遍历背包容量
			// 一个是取自己dp[j] 相当于 二维dp数组中的dp[i-1][j]，即不放物品i，
			// 一个是取dp[j - weight[i]] + value[i]，即放物品i，指定是取最大的，毕竟是求最大价值，
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	//fmt.Println(dp)
	return dp[bagWeight]
}
func test2WeiBagProblem(weight, value []int, bagWeight int) int {
	// 定义 and 初始化
	dp := make([][]int, len(weight))
	for i := range dp {
		dp[i] = make([]int, bagWeight)
	}
	// 背包小于基本重量的,可以不用考虑，暂时
	//for j := 0; j < weight[0]; j++ {
	//	dp[0][j] = 0
	//}
	//背包重量大于基本重量的时候
	for i := weight[0]; i <= bagWeight; i++ {
		dp[0][i] = value[0]
	}
	for i := 1; i < len(weight); i++ { // 先遍历物品
		for j := 0; j <= bagWeight; j++ { // 遍历背包
			if j < weight[i] {
				// 放不下物品i，那就等于前一个物品的最大价值
				dp[i][j] = dp[i-1][j]
			} else {
				// 放的下物品， 当前容量去除物品i重量，和不放物品i看哪个价值更大。
				dp[i][j] = max(dp[i-1][j-weight[i]]+value[i], dp[i-1][j])
			}
		}
	}

	//fmt.Println(dp)
	return dp[len(weight)-1][bagWeight]
}

// 0 1背包问题
func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	test1WeiBagProblem(weight, value, 4)
}

// 518. 零钱兑换 II 完全背包问题
// 输入：amount = 5, coins = [1, 2, 5]
// 输出：4
// 解释：有四种方式可以凑成总金额：
// 5=5
// 5=2+2+1
// 5=2+1+1+1
// 5=1+1+1+1+1
func change(amount int, coins []int) int {
	// 凑成总金额j的货币组合数为dp[j]
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

// 二维数组 原始叠加
func changeOrigin(amount int, coins []int) int {
	// 凑成总金额j的货币组合数为dp[j]
	// dp[i][j]为考虑前 i 件物品，凑成总和为 j 的方案数量。
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	// 初始化：当没有物品，背包容量也为0时，组合数为1
	dp[0][0] = 1
	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			// 从上一个物品开始
			dp[i][j] = dp[i-1][j]
			// 每个硬币可以被选择多次
			// 从第一个硬币开始，所以要扣除i-1
			for k := 1; k*coins[i-1] <= j; k++ {
				dp[i][j] += dp[i-1][j-k*coins[i-1]]
			}

		}
	}
	return dp[len(coins)][amount]
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
// 输入: s = "goalspecial",wordDict = ["go","goal","goals","special"]
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

// 674. 最长连续递增序列
// 输入：nums = [1,3,5,4,7]
// 输出：3
// 解释：最长连续递增序列是 [1,3,5], 长度为3。尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。
func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	ans := dp[0]
	// dp[i]以下标i为结尾的连续递增的子序列长度为dp[i]。
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

// 718. 最长重复子数组
// A: [1,2,3,2,1]
// B: [3,2,1,4,7]
// 输出：3
// 解释：长度最长的公共子数组是 [3, 2, 1] 。
func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	res := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums2[j-1] == nums1[i-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}

// 1143.最长公共子序列
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace"，它的长度为 3。
func longestCommonSubsequence(text1 string, text2 string) int {
	t1 := len(text1)
	t2 := len(text2)
	dp := make([][]int, t1+1)
	for i := range dp {
		dp[i] = make([]int, t2+1)
	}

	for i := 1; i <= t1; i++ {
		for j := 1; j <= t2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 如text1=acf,text2=cfe
				// f[2][3]=1, f[3][2] =2
				// 我们比较text1[3]与text2[3]，发现'f'不等于'e'，这样f[3][3]无法在原先的基础上延长，
				// 因此继承"ac"与"cfe" ，"acf"与"cf"的最长公共子序列中的较大值，
				// 即 f[3][3] = max(f[2][3] ,f[3][2]) = 2。
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[t1][t2]
}

// 53. 最大子序和
// 输入: [-2,1,-3,4,-1,2,1,-5,4]
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
func maxSubArrayV1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	res := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		res = maxs(res, dp[i])
	}
	return res
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
// 0 , 1背包问题
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
	dp := make([]int, target+1)

	for _, num := range nums { // 物品 不能重复
		for j := target; j >= num; j-- { // 背包
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}
	return dp[target] == target
}

// 32. 最长有效括号
// 输入：s = "(()(())"
// 输出：6
// 解释：最长有效括号子串是 "()(())"
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
