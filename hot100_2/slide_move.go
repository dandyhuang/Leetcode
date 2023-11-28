package hot100_2

// 3. 无重复字符的最长子串
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
func lengthOfLongestSubstring(s string) int {
	res := 0
	m := make(map[byte]bool, 0)
	l := 0
	for i := range s {
		c := s[i]
		for m[c] {
			delete(m, s[l])
			l++
		}
		m[c] = true
		res = max(res, i-l+1)
	}
	return res
}

// 438. 找到字符串中所有字母异位词 记住这里p会有重复的情况
// 输入: s = "cbaebabacd", p = "abc"
// 输出: [0,6]
// 解释:
// 起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
// 起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
func findAnagrams(s string, p string) []int {
	var res []int
	if len(s) < len(p) {
		return res
	}
	ms := make(map[byte]int)
	mp := make(map[byte]int)
	for i := range p {
		mp[p[i]]++
		ms[s[i]]++
	}
	if isEqual(ms, mp) {
		res = append(res, 0)
	}
	left := 0
	for i := len(p); i < len(s); i++ {
		ms[s[i]]++
		ms[s[left]]--
		if ms[s[left]] == 0 {
			delete(ms, s[left])
		}
		left++
		if isEqual(ms, mp) {
			res = append(res, left)
		}
	}
	return res
}

func isEqual(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
