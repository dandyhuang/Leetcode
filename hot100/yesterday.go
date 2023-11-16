package leetcode_hot100

import "sort"

func yThreeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				for l < r && nums[l] == nums[l+1] {
					l++
				}
			}
		}
	}
	return res
}

// 接雨水
func yTrap(height []int) int {
	total := 0
	return total
}

// 49. 字母异位词分组
func yGroupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[string][]string)
	for _, v := range strs {
		str := []byte(v)
		sort.Slice(str, func(i, j int) bool {
			return str[i] > str[j]
		})
		m[string(str)] = append(m[string(str)], v)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// 560. 和为 K 的子数组 前缀和h和哈希表
func ySubarraySum(nums []int, k int) int {
	return 0
}

// 2. 两数相加
func yAddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	plus := 0
	dump := &ListNode{}
	cur := dump
	for l1 != nil || l2 != nil {
		sum := 0
		sum += plus
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		plus = sum / 10
		dump.Next = &ListNode{Val: sum % 10}
		dump = dump.Next
	}
	if plus > 0 {
		dump.Next = &ListNode{Val: plus}
	}
	return cur.Next
}

// 25. K 个一组翻转链表
func yReverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	prevGroupHead := dummy
	for head != nil {
		groupHead := head
		groupTail := head
		for i := 1; i < k && groupTail != nil; i++ {
			groupTail = groupTail.Next
		}
		if groupTail == nil {
			break
		}
		nextGroup := groupTail.Next
		groupTail.Next = nil
		prevGroupHead.Next = reverseList(groupHead)
		groupHead.Next = nextGroup
		prevGroupHead = groupHead
		head = nextGroup
	}
	return dummy.Next
}

// 78. 子集
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
func ySubsets(nums []int) [][]int {
	var res [][]int
	var arr []int
	var dfs func(int, []int, []int)
	dfs = func(start int, nums []int, arr []int) {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		res = append(res, tmp)
		if start >= len(nums) {
			return
		}
		for i := start; i < len(nums); i++ {
			arr = append(arr, nums[i])
			dfs(start+1, nums, arr)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(0, nums, arr)
	return res
}

// 17. 电话号码的字母组合
func letterCombinations(digits string) []string {
	var res []string
	if digits == "" {
		return res
	}
	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var arr []byte
	var dfs func(string, []byte, int)
	dfs = func(digits string, arr []byte, start int) {
		if len(arr) == len(digits) {
			tmp := make([]byte, len(arr))
			copy(tmp, arr)
			res = append(res, string(tmp))
			return
		}
		letters := m[digits[start]]
		for i := 0; i < len(letters); i++ {
			arr = append(arr, letters[i])
			dfs(digits, arr, start+1)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(digits, arr, 0)

	return res
}

//
func ySolveNQueens(n int) [][]string {
	var res [][]string
	return res
}
