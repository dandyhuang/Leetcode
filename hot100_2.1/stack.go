package hot100_2

// 20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 输入：s = "()[]{}"
// 输出：true
func isValid(s string) bool {

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
	queue := make([]int, 0)
	for i := 0; i < len(T); i++ {
		for len(queue) > 0 && T[queue[len(queue)-1]] < T[i] {
			res[queue[len(queue)-1]] = i - queue[len(queue)-1]
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	return res
}

// 84. 柱状图中最大的矩形 和每日温度刚好相反
// 输入：heights = [2,1,5,6,2,3]
// 输出：10
// 当前栈，左右两边，找到最小值
func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	stack := make([]int, 0)
	res := 0
	for i := range heights {
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
