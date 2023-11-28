package hot100_2

import "sort"

// 两数之和
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1]
func twoSum(nums []int, target int) []int {
	var res []int
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok {
			return []int{i, v}
		} else {
			m[nums[i]] = i
		}
	}
	return res
}

// 49. 字母异位词分组
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[string][]string)
	for _, v := range strs {
		str := []byte(v)
		sort.Slice(str, func(i, j int) bool {
			return str[i] < str[j]
		})
		m[string(str)] = append(m[string(str)], v)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// 128. 最长连续序列 和 3. 无重复字符的最长子串相似
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
func longestConsecutive(nums []int) int {
	res := 0
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	for k, _ := range m {
		if !m[k-1] {
			size := 1
			for m[k+1] {
				k++
				size++
			}
			res = max(res, size)
		}
	}
	return res
}
