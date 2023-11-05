package leetcode_hot100

import "sort"

// 283. 移动零
func moveZeroes(nums []int) {
	start := 0
	for _, v := range nums {
		if v != 0 {
			nums[start] = v
			start++
		}
	}
	for i := start; i < len(nums); i++ {
		nums[start] = 0
		start++
	}
}

func mins(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 11. 盛最多水的容器
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	mArea := 0
	for l < r {
		min := mins(height[l], height[r])
		area := min * (r - l)
		mArea = maxs(area, mArea)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return mArea
}

// 15. 三数之和
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res [][]int
	// l== i+1所以-2
	for i := 0; i < len(nums)-2; i++ {
		// 因为单调递增，相等说明之前验证过,不重复的三元组
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				arr := []int{nums[i], nums[l], nums[r]}
				res = append(res, arr)
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}

	return res
}

// 接雨水
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	lMax := make([]int, len(height))
	rMax := make([]int, len(height))
	total := 0
	lMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		lMax[i] = maxs(lMax[i-1], height[i])
	}
	rMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rMax[i] = maxs(rMax[i+1], height[i])
	}
	for i := 0; i < len(height); i++ {
		size := mins(lMax[i], rMax[i]) - height[i]
		total += size
	}
	return total
}

// 84. 柱状图中最大的矩形
func largestRectangleArea(heights []int) int {
	total := 0
	return total
}
