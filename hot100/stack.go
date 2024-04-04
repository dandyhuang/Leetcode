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
// 输入：s = "3[a]2[bc]" | s = "3[a2[c]]"
// 输出："aaabcbc" ｜ "accaccacc"
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

// 42. 接雨水
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
func trapV2(height []int) int {
	var res = 0
	var stack []int
	for i, h := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			bottomH := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			// 栈为空，接不住雨水了
			if len(stack) == 0 {
				break
			}
			// 如 5 2 3 1 6数组，在index=4的时候数组为6时，栈内剩下[5,3,1], 出栈计算3的节点面积，就需要从5的起点开始，
			// 但是5这个点又不能计算，所以要-1。 并且由5到6构成一个凹槽，积累雨水，即3的左右两边高度取较低的5减去自身3。
			// 最后只剩下5这个数组，是不计算面积的，因为没有左边界
			// 如果不从5开始计算，2的这个结点的面积的计算不到了
			wide := i - stack[len(stack)-1] - 1
			// 5和6取5的高度，减去3
			leftHeight := min(h, height[stack[len(stack)-1]]) - bottomH
			res += leftHeight * wide
		}
		stack = append(stack, i)
	}
	return res
}

func largestRectangleAreaV2(heights []int) int {
	// 这里是为了不需要判断栈为空的时候。 和接雨水的问题，同样有这个问题，区分开来为什么
	heights = append(heights, 0)
	// 为了让单调递增的栈，最后的出栈计算
	heights = append([]int{0}, heights...)
	res := 0
	var stack []int
	for i, _ := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			bottomH := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			wide := i - stack[len(stack)-1] - 1
			res = max(res, bottomH*wide)
		}
		stack = append(stack, i)
	}
	return res
}
