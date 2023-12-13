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
			// 去除之前叠加的数据
			stack = stack[:len(stack)-1]
			times, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			// 叠加历史的数据
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

// 84. 柱状图中最大的矩形 和每日温度刚好相反
// 输入：heights = [2,1,5,6,2,3]
// 输出：10
// 当前栈，左右两边，找到最小值
func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	maxArea := 0
	for r := 0; r <= len(heights); r++ {
		var h int
		// 单调递增的情况
		if r == len(heights) {
			h = 0 // 处理最后一个元素，使得栈内元素全部出栈
		} else {
			h = heights[r]
		}
		// 当栈非空且当前高度小于栈顶高度时，出栈并计算面积
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := 0
			// 对头计算
			if len(stack) == 0 {
				left = -1
			} else {
				left = stack[len(stack)-1]
			}
			maxArea = max(maxArea, heights[idx]*(r-left-1))
		}
		// 将当前索引入栈
		stack = append(stack, r)
	}
	return maxArea
}
