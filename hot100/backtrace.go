package leetcode_hot100

import "fmt"

// 组合排序
// 77. 组合
// 输入：n = 4, k = 2
// 输出：
// [[2,4], [3,4], [2,3],[1,2], [1,3], [1,4]]
func combine(n int, k int) [][]int {
	var res [][]int
	var arr []int
	var dfs func(int, int, int, []int)
	dfs = func(n, k, start int, arr []int) {
		if len(arr)+n-start+1 < k {
			return
		}
		if len(arr) == k {
			tmp := make([]int, k)
			copy(tmp, arr)
			res = append(res, tmp)
			// 直接赋值有问题  拷贝的都是最后一个值
			// res = append(res, arr)
			return
		}
		for i := start; i <= n; i++ {
			arr = append(arr, i)
			dfs(n, k, i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(n, k, 1, arr)
	return res
}

// 46. 全排列
func permute(nums []int) [][]int {
	var res [][]int
	var arr []int
	used := make([]int, len(nums))
	var dfs func([]int, []int, []int)
	dfs = func(nums, used []int, arr []int) {
		if len(nums) == len(arr) {
			tmp := make([]int, len(nums))
			copy(tmp, arr)
			res = append(res, tmp)
			return
		}
		// i始终从0开始递归
		for i := 0; i < len(nums); i++ {
			if used[i] == 1 {
				continue
			}
			used[i] = 1
			arr = append(arr, nums[i])
			dfs(nums, used, arr)
			arr = arr[:len(arr)-1]
			used[i] = 0
		}
	}
	dfs(nums, used, arr)
	return res
}

// 78. 子集
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
func subsets(nums []int) [][]int {
	var res [][]int
	var arr []int
	var dfs func(int, []int)
	dfs = func(start int, arr []int) {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		res = append(res, tmp)
		if start >= len(nums) {
			return
		}
		for i := start; i < len(nums); i++ {
			arr = append(arr, nums[i])
			fmt.Println("in:", i, arr)
			dfs(i+1, arr)
			arr = arr[:len(arr)-1]
			fmt.Println("out:", i, arr)
		}
	}
	dfs(0, arr)
	return res
}

// 39. 组合总和 很难理解的，i为index
// 输入：candidates = [2,3,6,7], target = 7
// 输出：[[2,2,3],[7]]
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var arr []int
	var dfs func(int, int, []int)
	dfs = func(target, index int, arr []int) {
		if index == len(candidates) {
			return
		}

		if target == 0 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
			return
		}
		if target < 0 {
			return
		}

		for i := index; i < len(candidates); i++ {
			arr = append(arr, candidates[i])
			dfs(target-candidates[i], i, arr)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(target, 0, arr)
	return res
}

// 22. 括号生成 太抽象了
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
func generateParenthesis(n int) []string {
	var res []string
	var s []byte
	var dfs func(l, r, n int, s []byte)
	dfs = func(l, r, n int, s []byte) {
		if l == n && r == n {
			tmp := make([]byte, len(s))
			copy(tmp, s)
			res = append(res, string(s))
			return
		}
		if l < n {
			s = append(s, '(')
			dfs(l+1, r, n, s)
			s = s[:len(s)-1]
		}
		if r < l {
			s = append(s, ')')
			dfs(l, r+1, n, s)
			s = s[:len(s)-1]
		}
	}
	dfs(0, 0, n, s)
	return res
}
