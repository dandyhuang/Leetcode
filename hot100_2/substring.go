package hot100_2

// 560. 和为 K 的子数组 前缀和哈希表
// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
// 子数组是数组中元素的连续非空序列。
// 输入：nums = [1,1,1], k = 2
// 输出：2
func subarraySum(nums []int, k int) int {

}

// 560. 和为 K 的子数组
func subarraySum2(nums []int, k int) int {

}

type MyQueue struct {
	q []int
}

func (q *MyQueue) Push(val int) {
	for len(q.q) > 0 && val > q.q[len(q.q)-1] {
		q.q = q.q[:len(q.q)-1]
	}
	q.q = append(q.q, val)
}

func (q *MyQueue) Pop(val int) {
	if len(q.q) > 0 && q.q[0] == val {
		q.q = q.q[1:]
	}
}

// 239. 滑动窗口最大值
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
func maxSlidingWindow(nums []int, k int) []int {
	q := &MyQueue{}
	res := make([]int, 0)
	for i := 0; i < k && i < len(nums); i++ {
		q.Push(nums[i])
	}
	res = append(res, q.q[0])
	for i := k; i < len(nums); i++ {
		q.Pop(nums[i-k])
		q.Push(nums[i])
		res = append(res, q.q[0])
	}
	return res
}

// 76. 最小覆盖子串
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
func minWindow(s string, t string) string {

}
