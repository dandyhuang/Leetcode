package leetcode_hot100

import (
	"sort"
)

// 两数之和
func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok {
			res = append(res, v)
			res = append(res, i)
		} else {
			m[nums[i]] = i
		}
	}
	return res
}

// 49. 字母异位词分组
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, v := range strs {
		str := []byte(v)
		sort.Slice(str, func(i, j int) bool {
			return str[i] < str[j]
		})
		groups[string(str)] = append(groups[string(str)], v)
	}

	res := make([][]string, 0)
	for _, group := range groups {
		res = append(res, group)
	}
	return res
}

func maxs(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 128. 最长连续序列
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	size := 0
	for v, _ := range m {
		// 存在可以等轮询到v的时候，在处理。或者之前已经处理过了。golang map是随机k
		if !m[v-1] {
			max := 1
			for m[v+1] {
				max++
				v++
			}
			size = maxs(size, max)
		}
	}
	return size
}
