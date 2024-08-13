package __array

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
