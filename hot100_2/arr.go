package hot100_2

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 53. 最大子数组和 动态规划
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	res := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}
	return res
}

// 56. 合并区间 先排序，right始终保持上一个的最大值
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	l, r := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if r >= intervals[i][0] {
			r = max(intervals[i][1], r)
		} else {
			res = append(res, []int{l, r})
			l = intervals[i][0]
			r = intervals[i][1]
		}
	}
	res = append(res, []int{l, r})
	return res
}

// 189. 轮转数组 多个反转
// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右轮转 1 步: [7,1,2,3,4,5,6]
// 向右轮转 2 步: [6,7,1,2,3,4,5]
// 向右轮转 3 步: [5,6,7,1,2,3,4]
func revert(nums []int) {
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
func rotate(nums []int, k int) {
	k = k % len(nums)
	revert(nums)
	revert(nums[:k])
	revert(nums[k:])
}

// 238. 除自身以外数组的乘积 前后缀积
// 输入: nums = [1,2,3,4]
// 输出: [24,12,8,6]
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	lPlus, rPlus := make([]int, len(nums)), make([]int, len(nums))
	plus := 1
	for i := range nums {
		lPlus[i] = plus
		plus *= nums[i]
	}
	plus = 1
	for i := len(nums) - 1; i >= 0; i-- {
		rPlus[i] = plus
		plus *= nums[i]
	}
	for i := 0; i < len(nums); i++ {
		res[i] = lPlus[i] * rPlus[i]
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 41. 缺失的第一个正数
// 输入：nums = [3,4,-1,1]
// 输出：2
// [3 4 5 1]
// [-3 4 -5 -1]
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	// 第二轮: 根据数值将对应的索引处的数替换为负值。说明已经存在了
	// 如果3，4，2，1场景, 3的index=2被设置为-2，2的index就是-2了
	for i := range nums {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	// 第三轮: 找到第一个正数的索引
	for i := range nums {
		if nums[i] >= 0 {
			return i + 1
		}
	}
	return n + 1
}
