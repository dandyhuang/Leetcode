package leetcode_hot100

import "math"

// 560. 和为 K 的子数组 前缀和哈希表
// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
// 子数组是数组中元素的连续非空序列。
// 输入：nums = [1,1,1], k = 2
// 输出：2
func subarraySum(nums []int, k int) int {
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

// 560. 和为 K 的子数组 前缀和哈希表
func subarraySum2(nums []int, k int) int {
	m := make(map[int]int)
	// 初始前缀和为 0 的个数为 1
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
	sM := make(map[byte]int)
	tM := make(map[byte]int)

	for i, _ := range t {
		tM[t[i]]++
	}

	count := 0
	var res string
	for i, j := 0, 0; i < len(s); i++ {
		sM[s[i]]++
		if sM[s[i]] <= tM[s[i]] {
			count++
		}
		for j < len(s) && sM[s[j]] > tM[s[j]] {
			sM[s[j]]--
			j++
		}
		if count == len(t) {
			if len(res) == 0 || i-j+1 < len(res) {
				res = s[j : i+1]
			}
		}
	}
	return res
}

func minWindow1(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
