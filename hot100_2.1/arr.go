package hot100_2

import "sort"

// 53. 最大子数组和
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。!
func maxSubArray(nums []int) int {
	res := nums[0]
	sum := 0
	for i := range nums {
		sum += nums[i]
		if sum < nums[i] {
			sum = nums[i]
		}
		res = max(res, sum)
	}
	return res
}

func maxSubArrayV2(nums []int) int {
	res := 0
	dp := make([]int, len(nums))
	dp[0] = max(nums[0], 0)
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], 0)
		res = max(res, dp[i])
	}
	return res
}

// 56. 合并区间
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6]. !
// 1, 10
// 2, 3
func merge(intervals [][]int) [][]int {
	var res [][]int
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	l, r := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if r < intervals[i][0] {
			res = append(res, []int{l, r})
			l = intervals[i][0]
			r = intervals[i][1]
		} else {
			r = max(r, intervals[i][1])
		}
	}
	res = append(res, []int{l, r})
	return res
}

// 189. 轮转数组
// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右轮转 1 步: [7,1,2,3,4,5,6]
// 向右轮转 2 步: [6,7,1,2,3,4,5]
// 向右轮转 3 步: [5,6,7,1,2,3,4] !
func rotate(nums []int, k int) {
	k = k % len(nums)
	reversRotate(nums)
	reversRotate(nums[:k])
	reversRotate(nums[k:])
}

func reversRotate(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// 238. 除自身以外数组的乘积
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

// 41. 缺失的第一个正数
// 输入：nums = [3,-1,4,1]
// 输出：2
// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
func firstMissingPositive(nums []int) int {

}
