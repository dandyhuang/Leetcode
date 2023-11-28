package hot100_2

import "sort"

// 283. 移动零
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
func moveZeroes(nums []int) {
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 11. 盛最多水的容器
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49 7*7
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
func maxArea(height []int) int {
	res := 0
	l, r := 0, len(height)-1
	for l < r {
		w := r - l
		h := min(height[l], height[r])
		res = max(res, w*h)
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return res
}

// 15. 三数之和
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 解释：
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
// 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
// 注意，输出的顺序和三元组的顺序并不重要。
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
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
func trap(height []int) int {
	n := len(height)
	lHeight := make([]int, n)
	rHeight := make([]int, n)
	lHeight[0] = height[0]
	rHeight[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		lHeight[i] = max(lHeight[i-1], height[i])
	}
	for j := n - 2; j >= 0; j-- {
		rHeight[j] = max(rHeight[j+1], height[j])
	}
	var res int
	for i := 0; i < n; i++ {
		res += min(lHeight[i], rHeight[i]) - height[i]
	}
	return res
}
