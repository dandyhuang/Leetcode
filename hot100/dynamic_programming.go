package leetcode_hot100

import (
	"math"
	"sort"
)

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
	// 凑成总金额j的货币组合数为dp[j]个
	dp := make([]int, amount+1)
	// 当不选取任何硬币时，金额之和才为0，因此有1种硬币组合。
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			// 当前的币种的组合数就是dp[i - coin]
			// 别的币种组合数就是上一次计算的dp[i]，所以本次dp[i]就是上次的dp[i]+当前的组合数
			/* 比如j=5,i=1时
			   coins[i]=2,那么dp[5]是由dp[3]+coins[i]当前这枚硬币，
			   也能组合成了这次。dp[5]取决于dp[3]的组合数有多少种。

			   之前i=0的时候，dp[5]也是由dp[4]+coin[i]组合了这次。这时候的dp[5]也存在一个值，
			   是由i=0,coin等于1的时候，的一个结果
			*/
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

// 二维数组 原始叠加
func changeOriginV2(amount int, coins []int) int {
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
			dp[i][j] = dp[i-1][j]
			// 每个硬币可以被选择多次
			// 2.2 dp[i - 1][j]表示不加第i个硬币就能使总金额达到j的所有方案
			// 2.3 dp[i][j - k * coins[i]]表示加上k个第i个金币，能够使硬币总金额达到j的所有方案
			if j >= coins[i-1] {
				dp[i][j] += dp[i][j-coins[i-1]]
			}

		}
	}
	return dp[len(coins)][amount]
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

// 494. 目标和
// 输入：nums = [1,1,1,1,1], target = 3
// 输出：5
// 解释：一共有 5 种方法让最终目标和为 3 。
// -1 + 1 + 1 + 1 + 1 = 3
// +1 - 1 + 1 + 1 + 1 = 3
// +1 + 1 - 1 + 1 + 1 = 3
// +1 + 1 + 1 - 1 + 1 = 3
// +1 + 1 + 1 + 1 - 1 = 3
// 回溯
func findTargetSumWays(nums []int, target int) int {
	count := 0
	var dfs func(index, sum int)
	dfs = func(index, sum int) {
		if index == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		dfs(index+1, sum+nums[index])
		dfs(index+1, sum-nums[index])
	}
	dfs(0, 0)
	return count
}

// 回溯V2
func findTargetSumWaysV2(nums []int, target int) int {
	// 因为
	// l-r = target
	// l+r = sum
	// l-(sum-l) = target
	// l = (target+sum) /2
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if target > sum {
		return 0
	} // 此时没有方案
	if (target+sum)%2 != 0 {
		return 0 // 此时没有方案，两个int相加的时候要各位小心数值溢出的问题
	}
	bagSize := (target + sum) / 2 // 转变为组合总和问题，bagsize就是要求的和
	var arr []int
	var res [][]int
	var dfs func(index, sum int, arr []int)
	dfs = func(index, sum int, arr []int) {
		if sum > bagSize {
			return
		}
		if sum == bagSize {
			res = append(res, append([]int{}, arr...))
		}
		for i := index; i < len(nums); i++ {
			arr = append(arr, nums[i])
			dfs(i+1, sum+nums[i], arr)
			arr = arr[:len(arr)-1]
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	dfs(0, 0, arr)
	return len(res)
}

// 动态规划
func findTargetSumWaysDpTwo(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	n, neg := len(nums), diff/2
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	// 如果j<num,num不能选 dp[i][j] = 	dp[i-1][j]
	// 如果j>=num， num能选或者不选 dp[i][j] = dp[i-1][j]+dp[i-1][j-nums[i]]
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= neg; j++ {
			// 从上一个物品开始
			dp[i][j] = dp[i-1][j]
			// 每个硬币可以被选择多次
			// 从第一个硬币开始，所以要扣除i-1
			if j >= nums[i-1] {
				// 用i-1个数凑成和为 j-num的总数，给这些方案都加上一个num，和就是 j

				// dp[i][j]的含义，它表示的是在前i个数中选出一些数使其和为j，如果选择num，
				// 相当于延长已经选择过的数据序列，而不是增加方案数，
				// 方案数的变化应该是考虑选择当前的num是一种方案，
				// 不选择又是另外一种方案，所以就该把二者相加，作为总的方案数。
				dp[i][j] += dp[i-1][j-nums[i-1]]
			}

		}
	}
	return dp[n][neg]
}

// 完全背包问题呢
func findTargetSumWaysDp(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if abs(target) > sum {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	// 计算背包大小
	bag := (sum + target) / 2
	// dp[j] 表示：填满j（包括j）这么大容积的包，有dp[j]种方法
	dp := make([]int, bag+1)
	// 初始化
	dp[0] = 1
	// 遍历顺序
	for i := 0; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			//推导公式
			dp[j] += dp[j-nums[i]]
			//fmt.Println(dp)
		}
	}
	return dp[bag]
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

// dsf单词拆分
func wordBreakV2(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}
	// 记忆搜索
	memo := make(map[int]bool)
	var dfs func(start int) bool
	dfs = func(start int) bool {
		if start == len(s) {
			return true
		}
		if res, ok := memo[start]; ok {
			return res
		}
		for i := start + 1; i <= len(s); i++ {
			if !wordMap[s[start:i]] {
				continue
			}
			if dfs(i) {
				memo[start] = true
				return true
			}
		}
		memo[start] = false
		return false
	}

	return dfs(0)
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

// 392.判断子序列
// 输入：s = "abc", t = "ahbgdc"
// 输出：true
// 双指针
func isSubsequenceV2(s string, t string) bool {
	left := 0
	for i := 0; i < len(t) && left < len(s); i++ {
		if s[left] == t[i] {
			left++
		}
	}
	return left == len(s)
}

func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[n][m] == n
}

// 115.不同的子序列
// 给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数，结果需要对 109 + 7 取模。
// 输入：s = "rabbbit", t = "rabbit"  输出：3
// 解释：
// 如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
// rab bbit
// rabb bit
// rabbb it
func numDistinct(s string, t string) int {

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
	// dp[i] 表示以 s[i] 结尾的最长有效括号的长度,以 ‘(’ 结尾的子串对应的 dp 值必定为0
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
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' { // '()(())' 看最后一个）情况
				if i-dp[i-1] >= 2 { // '()))'
					// 除了判断前面的i-1的数据是否是有效括号，还要判断i-dp[i-1]-1前面的i-dp[i-1]-2是否也是有效括号
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
