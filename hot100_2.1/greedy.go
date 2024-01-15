package hot100_2

import "math"

// 121. 买卖股票的最佳时机
// 输入：[7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
// 注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
func maxProfit(prices []int) int {
	cost := math.MaxInt
	profit := 0
	for i := range prices {
		profit = max(profit, prices[i]-cost)
		cost = min(cost, prices[i])
	}
	return profit
}

// 55. 跳跃游戏
// 输入：nums = [2,3,1,1,4]
// 输出：true
// 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
func canJump(nums []int) bool {
	maxJump := nums[0]
	for i := 0; i < len(nums); i++ {
		if i > maxJump {
			return false
		}
		maxJump = max(maxJump, i+nums[i])
	}
	return true
}

// 45. 跳跃游戏 II
// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
// 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
func jump(nums []int) int {
}

// 763. 划分字母区间
// 输入：s = "ababcbacadefegdehijhklij"
// 输出：[9,7,8]
// 解释：
// 划分结果为 "ababcbaca"、"defegde"、"hijhklij" 。
func partitionLabels(s string) []int {
}
