package hot100_2

import "strconv"

// 20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 输入：s = "()[]{}"
// 输出：true
func isValid(s string) bool {
	stack := make([]rune, 0)
	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, str := range s {
		if v, ok := m[str]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == v {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, str)
		}
	}
	return len(stack) == 0
}

// 394. 字符串解码
// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"
func decodeString(s string) string {
	var res string
	stack := make([]string, 0)
	num := 0
	for _, str := range s {
		if str >= '0' && str <= '9' {
			num = num*10 + int(str-'0')
		} else if str == '[' {
			stack = append(stack, strconv.Itoa(num), res)
			num = 0
			res = ""
		} else if str == ']' {
			preStr := stack[:len(stack)-1]
			size := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

		} else {
			res += string(str)
		}
	}
	return res
}

// 739. 每日温度
// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，
// 下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
// 输入: temperatures = [73,74,75,71,69,72,76,73]
// 输出: [1,1,4,2,1,1,0,0]
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	stack := make([]int, 0)
	for i := range T {
		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			index := stack[len(stack)-1]
			res[index] = i - index
			stack = stack[:len(stack)-1]
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
	res := 0
	stack := make([]int, 0)
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	for i := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 找左边比它小的
			l := stack[len(stack)-1] + 1
			// 找右边比它小的
			r := i - 1
			res = max(res, (r-l+1)*heights[cur])
		}
		stack = append(stack, i)
	}
	return res
}
