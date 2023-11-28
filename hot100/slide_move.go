package leetcode_hot100

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

// 438. 找到字符串中所有字母异位词
func isEqual(pM, sM map[byte]int) bool {
	if len(pM) != len(sM) {
		return false
	}
	for k, _ := range pM {
		if pM[k] != sM[k] {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	var result []int

	if len(s) < len(p) {
		return result
	}

	pMap := make(map[byte]int)
	sMap := make(map[byte]int)
	for i, _ := range p {
		pMap[p[i]]++
		sMap[s[i]]++
	}
	if isEqual(pMap, sMap) {
		result = append(result, 0)
	}

	for i := len(p); i < len(s); i++ {
		sMap[s[i]]++
		sMap[s[i-len(p)]]--
		if sMap[s[i-len(p)]] == 0 {
			// 删除改元素
			delete(sMap, s[i-len(p)])
		}
		if isEqual(pMap, sMap) {
			result = append(result, i-len(p)+1)
		}
	}
	return result
}
