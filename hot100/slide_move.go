package leetcode_hot100

// 3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	total := 0
	str := []byte(s)
	m := make(map[byte]int)
	start := 0
	for i := 0; i < len(str); i++ {
		// abba情况
		if v, ok := m[str[i]]; ok && v >= start {
			// 相等说明找到了
			num := i - start
			total = maxs(num, total)
			start = m[str[i]] + 1
		}
		m[str[i]] = i
	}
	fina := len(str) - start
	total = maxs(fina, total)
	return total
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
