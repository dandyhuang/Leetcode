package leetcode_hot100

import (
	"sort"
)

// 53. 最大子数组和 动态规划
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = maxs(nums[i], dp[i-1]+nums[i])
		res = maxs(res, dp[i])
	}

	return res
}

// 合并区间
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	l, r := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > r {
			arr := []int{l, r}
			res = append(res, arr)
			l = intervals[i][0]
			r = intervals[i][1]
		} else {
			r = maxs(r, intervals[i][1])
		}
	}
	return res
}

// 189. 轮转数组
func reserver(nums []int) {
	for l, r := 0, len(nums)-1; l < r; {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
func rotate(nums []int, k int) {
	k = k % len(nums)
	reserver(nums)
	reserver(nums[:k])
	reserver(nums[k:])
}

// 238. 除自身以外数组的乘积 前后缀积
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	left, right := make([]int, len(nums)), make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			left[0] = 1
		} else {
			left[i] = left[i-1] * nums[i-1]
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			right[i] = 1
		} else {
			right[i] = right[i+1] * nums[i+1]
		}
	}
	for i := 0; i < len(nums); i++ {
		res[i] = left[i] * right[i]
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
func firstMissingPositive(nums []int) int {
	// 第一轮: 将所有小于等于0的数替换为n+1
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	// 第二轮: 根据数值将对应的索引处的数替换为负值。说明已经存在了
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	// 第三轮: 找到第一个正数的索引
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}
