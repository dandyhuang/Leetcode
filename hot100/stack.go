package leetcode_hot100

import (
	"strconv"
	"strings"
)

// 20. 有效的括号
func isValid(s string) bool {
	stack := make([]byte, 0)
	m := map[int32]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	for i, str := range s {
		if v, ok := m[str]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == v {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// 394. 字符串解码
// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"
func decodeString(s string) string {
	stack := make([]string, 0)
	num := 0
	res := ""
	for _, char := range s {
		if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		} else if char == '[' {
			// 保存原来的字符串
			stack = append(stack, strconv.Itoa(num), res)
			// 留给后续构造使用
			res = ""
			num = 0
		} else if char == ']' {
			preRes := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			times, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			res = preRes + strings.Repeat(res, times)
		} else {
			res += string(char)
		}
	}
	return res
}

// 739. 每日温度
func dailyTemperatures(T []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(T))
	for i, _ := range T {
		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[index] = i - index
		}
		stack = append(stack, i)
	}

	return res
}
