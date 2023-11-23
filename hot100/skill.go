package leetcode_hot100

// 136. 只出现一次的数字
// 输入：nums = [4,1,2,1,2]
// 输出：4
func singleNumber(nums []int) int {
	// 任何数与自身异或结果为0
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

// 169. 多数元素
// 输入：nums = [3,2,3]
// 输出：3  Boyer-Moore投票算法
func majorityElement(nums []int) int {
	var candidate, count int

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if candidate == num {
			count++
		} else {
			count--
		}
	}

	return candidate
}

// 75. 颜色分类
// 输入：nums = [2,0,2,1,1,0]
// 输出：[0,0,1,1,2,2] 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
// 荷兰国旗问题
func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	start := 0
	for start <= r {
		switch nums[start] {
		case 0:
			nums[l], nums[start] = nums[start], nums[l]
			start++
			l++
		case 1:
			start++
		case 2:
			nums[r], nums[start] = nums[start], nums[r]
			r--
		}
	}
}

// 31. 下一个排列
// 例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
// 类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
// 而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
func nextPermutation(nums []int) {

}

// 287. 寻找重复数
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]

	// 快慢指针相遇
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// 重新设置慢指针，并与快指针相遇
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	// 返回相遇点即为重复数
	return slow
}
