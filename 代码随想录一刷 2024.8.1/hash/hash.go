package hash

import (
	"fmt"
	"sort"
)

// 242.有效的字母异位词
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 示例 1: 输入: s = "anagram", t = "nagaram" 输出: true
// 示例 2: 输入: s = "rat", t = "car" 输出: false
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[uint8]int)
	m2 := make(map[uint8]int)
	for i := range s {
		m1[s[i]]++
		m2[t[i]]++
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

// 349. 两个数组的交集
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出：[9,4]
// 解释：[4,9] 也是可通过的
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}, 0) // 用map模拟set
	res := make([]int, 0)
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		//如果存在于上一个数组中，则加入结果集，并清空该set值
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}

// 第202题. 快乐数
// 重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 输入：n = 19
// 输出：true
// 解释：
// 1^2 + 9^2 = 82
// 8^2 + 2^2 = 68
// 6^2 + 8^2 = 100
// 1^2 + 0^2 + 02 = 1
func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 {
		m[n] = true
		num := n
		n = 0
		for num != 0 {
			n += (num % 10) * (num % 10)
			num = num / 10
		}
		if m[n] == true {
			return false
		}
	}
	return true
}

// 第454题.四数相加II 和为0的数
func fourSumCount(a []int, b []int, c []int, d []int) (ans int) {
	countAB := map[int]int{}
	for _, v := range a {
		for _, w := range b {
			countAB[v+w]++
		}
	}
	for _, v := range c {
		for _, w := range d {
			ans += countAB[-v-w]
		}
	}
	return

}

// 383. 赎金信
// 给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
// 输入：ransomNote = "aa", magazine = "aab"
// 输出：true
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[int32]int, 0)
	for _, c := range magazine {
		m[c]++
	}
	for _, c := range ransomNote {
		m[c]--
		if m[c] < 0 {
			return false
		}
	}
	return true
}

// 第15题. 三数之和
// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
// 满足要求的三元组集合为： [ [-1, 0, 1], [-1, -1, 2] ]
// 双指针思想
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		if i+2 < n && nums[i]+nums[i+1]+nums[i+2] > 0 {
			continue
		}
		if nums[i]+nums[n-2]+nums[n-1] < 0 {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			c := nums[i] + nums[l] + nums[r]
			if c > 0 {
				r--
				continue
			}
			if c < 0 {
				l++
				continue
			}
			fmt.Println(i, l, r)
			res = append(res, []int{nums[i], nums[l], nums[r]})
			for l+1 < n && nums[l] == nums[l+1] {
				l++
			}
			for r-1 >= 0 && nums[r] == nums[r-1] {
				r--
			}
			l++
			r--
		}
	}

	return res
}
func threeSumMap(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		m := make(map[int]bool)
		for j := i + 1; j < len(nums); j++ {
			// 0, 0, 0 的场景
			if j > i+2 && nums[j] == nums[j-1] && nums[j-2] == nums[j-1] {
				continue
			}
			c := 0 - nums[i] - nums[j]
			if m[c] {
				res = append(res, []int{nums[i], nums[j], c})
				delete(m, c)
			} else {
				m[nums[j]] = true
			}
		}
	}
	return res
}

// 第18题. 四数之和
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if i+3 < n && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			continue
		}
		if nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j+2 < n && nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				continue
			}

			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}

			l := j + 1
			r := n - 1
			for l < r {
				c := nums[i] + nums[l] + nums[r] + nums[j]
				if c > target {
					r--
					continue
				}
				if c < target {
					l++
					continue
				}
				res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
				for l+1 < n && nums[l] == nums[l+1] {
					l++
				}
				for r-1 >= 0 && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return res
}
