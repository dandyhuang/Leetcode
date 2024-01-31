package hot100_2

import (
	"strconv"
	"strings"
)

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
// 比如我图中，当栈里是 1 5 6 这三个数时，此时 索引 i 代表的数是 2 ，触发计算。
// 此时第一次计算是因为 栈顶 6 比 2 大。所以出栈 6 ，6 为矩形的 h ，意思就是计算高度为 6 的矩形面积。
// 所以 left 就是出栈之后 5 的右边开始算起，5 不能包含在内。
// right 就是 i - 1 ，就是 2 的左边。
// 宽度计算完就是 1 格，高度 6 。
// 接下来，会对比栈顶 5 ，与 2 比较也会出栈。
// 此时要计算是高度为 5 的矩形。
// left 是 1 的右边。
// right 还是 i - 1，2 的左边。 宽度 2 ， 高度 5 。
func largestRectangleArea(heights []int) int {
	res := 0
	stack := make([]int, 0)
	// 第一个0是为了计算left的值
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	for i := range heights {
		// 如果是2，1，2的场景是怎么计算的呢, 最后一个最小值，一定是从最开始计算到某位所以一定会计算出来
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			l := stack[len(stack)-1] + 1
			r := i - 1
			res = max(res, (r-l+1)*h)
		}
		stack = append(stack, i)
	}
	return res
}
