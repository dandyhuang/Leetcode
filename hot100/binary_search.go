package leetcode_hot100

// 35. 搜索插入位置
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 如果循环结束时仍未找到目标值，返回插入位置
	return left
}

// 74. 搜索二维矩阵
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
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {
			// 左半段有序
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
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
// [3,4,5,1,2]
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if l == r {
			return nums[l]
		}
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			// 中间元素大于右边界，说明最小值在右半段
			l = mid + 1
		} else {
			// 中间元素小于等于右边界，说明最小值在左半段或就是当前元素
			r = mid
		}
	}
	return nums[l]
}

// 4. 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
	return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
	return 0
}
