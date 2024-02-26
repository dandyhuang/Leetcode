package hot100_2

// 20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 输入：s = "()[]{}"
// 输出：true
func isValid(s string) bool {
	m := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			stack = append(stack, m[s[i]])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != s[i] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return true
}

// 394. 字符串解码
// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"
func decodeString(s string) string {
	
}

// 739. 每日温度
// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，
// 下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
// 输入: temperatures = [73,74,75,71,69,72,76,73]
// 输出: [1,1,4,2,1,1,0,0]
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	stack := make([]int, 0)
	for i := 0; i < len(T); i++ {
		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[start] = i - start
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
	heights = append(heights, 0)
	res := 0
	return res
}
