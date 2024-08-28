package __array

import (
	"math"
	"sort"
)

func max(i, j int) int { return int(math.Max(float64(i), float64(j))) }
func min(i, j int) int { return int(math.Min(float64(i), float64(j))) }

// 数组
// 704. 二分查找
// 力扣题目链接
// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，
// 写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
// 输入: nums = [-1,0,3,5,9,12], target = 9
// 输出: 4
// 解释: 9 出现在 nums 中并且下标为 4
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// 27. 移除元素
// 力扣题目链接
// 示例 1: 给定 nums = [3,2,2,3], val = 3, 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
// 你不需要考虑数组中超出新长度后面的元素。
func removeElement(nums []int, val int) int {
	index := 0
	for i := range nums {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

// 977.有序数组的平方
// 输入：nums = [-4,-1,0,3,10]
// 输出：[0,1,9,16,100]
// 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]
func sortedSquares(nums []int) []int {
	res := make([]int, 0)
	slow, fast := 0, len(nums)-1
	if nums[fast] < 0 {
		for i := len(nums) - 1; i >= 0; i-- {
			res = append(res, nums[i]*nums[i])
		}
		return res
	}
	for slow <= fast {
		if nums[slow] < 0 {
			if int(math.Abs(float64(nums[slow]))) > nums[fast] {
				res = append(res, nums[slow]*nums[slow])
				slow++
			} else {
				res = append(res, nums[fast]*nums[fast])
				fast--
			}
		} else {
			res = append(res, nums[slow]*nums[slow])
			slow++
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}

// 209.长度最小的子数组
// 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，
// 并返回其长度。如果不存在符合条件的子数组，返回 0。
// 输入：s = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
func minSubArrayLen(s int, nums []int) int {
	slow := 0
	res := 0
	sum := math.MaxInt
	for start := 0; start < len(nums); start++ {
		res += nums[start]
		for res >= s {
			sum = min(sum, start-slow+1)
			res -= nums[slow]
			slow++
		}
	}
	if sum == math.MaxInt {
		return 0
	}
	return sum
}

// 904. 水果成篮
func totalFruit(fruits []int) int {
	left := 0
	res := math.MinInt
	mkind := make(map[int]int)
	for i := 0; i < len(fruits); i++ {
		mkind[fruits[i]]++
		for len(mkind) > 2 {
			if mkind[fruits[left]] == 1 {
				delete(mkind, fruits[left])
				left++
				break
			}
			mkind[fruits[left]]--
			left++
		}
		res = max(res, i-left+1)
	}
	return res
}

// 59.螺旋矩阵II
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	size := 1
	for i := range res {
		res[i] = make([]int, n)
	}
	t, b, l, r := 0, n-1, 0, n-1
	for t <= b && l <= r {
		for i := l; i <= r; i++ {
			res[t][i] = size
			size++
		}
		t++
		for i := t; i <= b; i++ {
			res[i][r] = size
			size++
		}
		r--
		for i := r; i >= l; i-- {
			res[b][i] = size
			size++
		}
		b--
		for i := b; i >= t; i-- {
			res[i][l] = size
			size++
		}
		l++
	}
	return res
}

// 58. 区间和
//本题为代码随想录后续扩充题目，还没有视频讲解，顺便让大家练习一下ACM输入输出模式（笔试面试必备）
// 前缀和
