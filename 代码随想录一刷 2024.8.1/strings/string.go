package strings

import (
	"strings"
)

// 344.反转字符串
// 输入：s = ["h","e","l","l","o"]
// 输出：["o","l","l","e","h"]
func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

// 541. 反转字符串II
// 每隔k个反转k个，末尾不够k个时全部反转
// 输入：s = "abcdefg", k = 2
// 输出："bacdfeg"
func reverseStr(s string, k int) string {
	t := []byte(s)

	for i := 0; i < len(t); i += 2 * k {
		if k+i < len(t) {
			reverseString(t[i : k+i])
		} else {
			reverseString(t[i:len(t)])
		}
	}
	return string(t)
}

// 151.翻转字符串里的单词
// 输入：s = "the sky is blue"
// 输出："blue is sky the"
// 从后向前后去单词
func reverseWords(str string) string {
	res := ""
	s := strings.Trim(str, " ")
	i, j := len(s)-1, len(s)-1
	for i >= 0 {
		for i >= 0 && s[i] != ' ' {
			i--
		}
		res += s[i+1:j+1] + " "
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
	}

	return res[:len(res)-1]
}

func reverseWordsV2(s string) string {
	sliceStr := strings.Split(s, " ")
	res := ""
	l, r := 0, len(sliceStr)-1
	for l < r {
		sliceStr[l], sliceStr[r] = sliceStr[r], sliceStr[l]
		l++
		r--
	}
	for _, str := range sliceStr {
		if str == "" {
			continue
		}
		res += str + " "
	}
	return res[:len(res)-1]
}

// 28. 实现 strStr() KMP实现 ？？？？
func strStr(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)

	if l2 == 0 {
		return 0
	}
	if l1 == 0 || l1 < l2 {
		return -1
	}

	for i := 0; i <= l1-l2; i++ {
		if haystack[i:i+l2] == needle {
			return i
		}
	}
	return -1
}

// 459.重复的子字符串
// 输入: s = "abcabcabcabc"
// 输出: true
// 解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)
func repeatedSubstringPattern(s string) bool {
	// 如果长度为1，直接返回false
	if len(s) == 1 {
		return false
	}
	// 遍历每种可能符合的子串
	for i := 1; i < len(s); i++ {
		// 如果子串长度不能整除字符串长度则跳过
		// if len(s) % i != 0 {
		//     continue
		// }
		j := i
		n := 1

		// 移动每个区间，判断是否都相同
		for j*(n+1) <= len(s) && s[0:i] == s[j*n:j*(n+1)] {
			n++
			if j*n == len(s) {
				return true
			}
		}
	}
	return false
}
