package hot100_2

import "math"

// 121. 买卖股票的最佳时机
func maxProfit(prices []int) int {
	profit := 0
	cost := math.MaxInt
	for _, v := range prices {
		cost = min(cost, v)
		profit = max(profit, v-cost)
	}
	return profit
}

// 55. 跳跃游戏
func canJump(nums []int) bool {
	maxJump := 0
	for i, v := range nums {
		// 判断条件
		if i > maxJump {
			return false
		}
		maxJump = max(maxJump, i+v)
		if maxJump >= len(nums)-1 {
			return true
		}
	}
	return false
}

// 45. 跳跃游戏 II
func jump(nums []int) int {
	maxJump := 0
	step := 0
	end := 0
	for i, v := range nums {
		//  最后一步不需要在 + step了
		if i == len(nums)-1 {
			break
		}
		maxJump = max(maxJump, i+v)
		if i == end {
			end = maxJump
			step++
		}
	}
	return step
}

// 763. 划分字母区间
// 输入：s = "ababcbacadefegdehijhklij"
// 输出：[9,7,8]
// 解释：
// 划分结果为 "ababcbaca"、"defegde"、"hijhklij" 。
func partitionLabels(s string) []int {
	lastIndex := make(map[int32]int)
	for i, v := range s {
		lastIndex[v] = i
	}
	var res []int
	start, end := 0, 0
	for i, v := range s {
		end = max(end, lastIndex[v])
		if end == i {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}
