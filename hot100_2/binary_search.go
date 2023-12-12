package hot100_2

// 35. 搜索插入位置
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

// 74. 搜索二维矩阵
// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// 输出：true
func searchMatrixV2(matrix [][]int, target int) bool {
	rows, _ := len(matrix), len(matrix[0])
	for i := 0; i < rows; i++ {
		left, right := 0, len(matrix[i])-1
		for left <= right {
			mid := left + (right-left)/2
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]
func searchRange(nums []int, target int) []int {
	findFirst := func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		first := -1
		for left <= right {
			mid := left + (right-left)/2

			if nums[mid] == target {
				// 更新最后一个位置，并继续在右侧搜索
				first = mid
				right = mid - 1
			} else if nums[mid] < target {
				// 目标值在右侧，更新左边界
				left = mid + 1
			} else {
				// 目标值在左侧，更新右边界
				right = mid - 1
			}
		}
		return first
	}
	findLast := func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		last := -1
		for left <= right {
			mid := left + (right-left)/2

			if nums[mid] == target {
				// 更新最后一个位置，并继续在右侧搜索
				last = mid
				left = mid + 1
			} else if nums[mid] < target {
				// 目标值在右侧，更新左边界
				left = mid + 1
			} else {
				// 目标值在左侧，更新右边界
				right = mid - 1
			}
		}
		return last
	}
	first := findFirst(nums, target)
	last := findLast(nums, target)

	return []int{first, last}
}

// 33. 搜索旋转排序数组
// 输入：nums = [4,5,6,7,0,1,2], target = 0
// 输出：4
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {
			// 左边有序
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// 右边有序
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

// 153. 寻找旋转排序数组中的最小值
// 输入：nums = [3,4,5,1,2]
// 输出：1
// 解释：原数组为 [1,2,3,4,5] ，旋转 3 次得到输入数组。
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if l == r {
			return nums[l]
		}
		mid := (r-l)/2 + l
		// 如果对比左边，可能3-5和1-2都是一样的
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// 4. 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

}
