package hot100_2

// 560. 和为 K 的子数组 前缀和哈希表
// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
// 子数组是数组中元素的连续非空序列。
// 输入：nums = [1,1,1], k = 2
// 输出：2
// 因为 preRight-preLeft=k
// nums 的 第 i 到 j 项 的和，有：
// nums[i]+…+nums[j]=prefixSum[j]−prefixSum[i−1]
// 当 i 为 0，此时 i-1 为 -1，我们故意让 prefixSum[-1] 为 0，使得通式在i=0时也成立：
// nums[0]+…+nums[j]=prefixSum[j]
func subarraySum(nums []int, k int) int {

}

// 560. 和为 K 的子数组 前缀和哈希表
func subarraySum2(nums []int, k int) int {
	count := 0
	m := make(map[int]int)
	m[0] = 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		num := sum - k
		if v, ok := m[num]; ok {
			count += v
		}
		m[sum]++
	}
	return count
}

// 239. 滑动窗口最大值
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
func maxSlidingWindow(nums []int, k int) []int {

}

// 76. 最小覆盖子串 还是滑动窗口思想
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
func minWindow(s string, t string) string {

}
