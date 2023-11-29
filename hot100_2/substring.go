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
	m := make(map[int]int)
	sum := 0
	count := 0
	// 当第一次sum-k为0的时候，就会出现漏记录的情况，比如2,-1,-1,0 k=0的时候
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if m[sum-k] > 0 {
			// 为什么不是count++而是+v呢,因为不是m[sum]>0
			count += m[sum-k]
		}
		m[sum]++
	}
	return count
}

// 560. 和为 K 的子数组
func subarraySum2(nums []int, k int) int {
	count := 0
	preSum := make([]int, len(nums)+1)
	preSum[0] = nums[0] // preSum[0] 应该为0
	// 通过暴力请求，会发现有问题
	for i := 1; i < len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if preSum[i]-preSum[j] == k {
				count++
			}
		}
	}
	return count
}

// 239. 滑动窗口最大值
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	queue := make([]int, 0)
	for i := 0; i < k; i++ {
		for len(queue) > 0 && nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
	}
	res = append(res, queue[0])

	for i := k; i < len(nums); i++ {
		if nums[i-k] == queue[0] {
			queue = queue[1:]
		}
		for len(queue) > 0 && nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		res = append(res, queue[0])
	}
	return res
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
func maxSlidingWindowV2(nums []int, k int) []int {
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

// 76. 最小覆盖子串 还是滑动窗口思想
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
func minWindow(s string, t string) string {
	var res string
	if len(s) < len(t) {
		return res
	}
	mS := make(map[byte]int)
	mT := make(map[byte]int)
	for i := range t {
		mT[t[i]]++
	}
	check := func(mS, mT map[byte]int) bool {
		if len(mS) < len(mT) {
			return false
		}
		for k, v := range mT {
			if mS[k] < v {
				return false
			}
		}
		return true
	}
	left := 0
	n := len(s) + 1
	for i := range s {
		mS[s[i]]++
		for check(mS, mT) {
			if i-left+1 < n {
				n = i - left + 1
				res = s[left : i+1]
			}

			mS[s[left]]--

			left++
		}
	}

	return res
}
