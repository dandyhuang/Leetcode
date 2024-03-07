package leetcode_hot100

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

// 135. 分发糖果
// n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
// 你需要按照以下要求，给这些孩子分发糖果：
// 每个孩子至少分配到 1 个糖果。
// 相邻两个孩子评分更高的孩子会获得更多的糖果。
// 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
// 输入：ratings = [1,0,2]
// 输出：5
// 解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
func candy(ratings []int) int {
	res := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		res[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			res[i] = res[i-1] + 1
		}
	}
	for j := len(ratings) - 2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] && res[j] <= res[j+1] {
			res[j] = res[j+1] + 1
		}
	}
	sum := 0
	for i := range res {
		sum += res[i]
	}
	return sum
}
