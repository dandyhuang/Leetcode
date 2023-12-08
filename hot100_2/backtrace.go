package hot100_2

import "sort"

// 46. 全排列
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func permute(nums []int) [][]int {
	var res [][]int
	var arr []int
	n := len(nums)
	used := make([]int, n+1)
	var dfs func(arr []int)
	dfs = func(arr []int) {
		if len(arr) == n {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] == 1 {
				continue
			}
			used[i] = 1
			arr = append(arr, nums[i])
			dfs(arr)
			arr = arr[:len(arr)-1]
			used[i] = 0
		}
	}
	dfs(arr)
	return res
}

// 78. 子集
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
func subsets(nums []int) [][]int {
	var res [][]int
	var arr []int
	var dfs func(start int, arr []int)
	dfs = func(start int, arr []int) {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		res = append(res, tmp)
		if start >= len(nums) {
			return
		}
		for i := start; i < len(nums); i++ {
			arr = append(arr, nums[i])
			dfs(i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(0, arr)
	return res
}

// 90. 子集 II
// 输入：nums = [1,2,2]
// 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
func subsetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] < nums[j] {
			return true
		}
		return false
	})
	var res [][]int
	var arr []int
	var dfs func(start int, arr []int)
	dfs = func(start int, arr []int) {
		res = append(res, append([]int{}, arr...))
		for i := start; i < len(nums); i++ {
			// 相同的去除
			if i-1 >= start && nums[i] == nums[i-1] {
				continue
			}
			arr = append(arr, nums[i])
			dfs(i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(0, arr)
	return res
}

// 组合排序
// 77. 组合
// 输入：n = 4, k = 2
// 输出：
// [[2,4], [3,4], [2,3],[1,2], [1,3], [1,4]]
func combine(n int, k int) [][]int {
	var res [][]int
	var arr []int
	var dfs func(start int, arr []int)
	dfs = func(start int, arr []int) {
		if len(arr) == k {
			var tmp []int
			copy(tmp, arr)
			res = append(res, tmp)
		}
		for i := start; i < start; i++ {
			arr = append(arr, i)
			dfs(start+1, arr)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(1, arr)
	return res
}

// 39. 组合总和
// 输入：candidates = [2,3,6,7], target = 7
// 输出：[[2,2,3],[7]]
func combinationSum(candidates []int, target int) [][]int {
	sum := 0
	var arr []int
	var res [][]int
	var dfs func(arr []int)
	dfs = func(arr []int) {

		for i := 0; i < len(candidates); i++ {

		}
		return
	}
	dfs(arr)
	return res
}

// 22. 括号生成 太抽象了
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
func generateParenthesis(n int) []string {

}

// 79. 单词搜索 和岛屿数量是一样的
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// 输出：true
func exist(board [][]byte, word string) bool {

}

// 131. 分割回文串
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
func partition(s string) [][]string {

}

// 51. N 皇后
// 输入：n = 4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
func isQueens(row, col int, arr [][]byte) bool {

}
