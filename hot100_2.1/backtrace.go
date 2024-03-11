package hot100_2

import "sort"

// 组合排序
// 77. 组合
// 输入：n = 4, k = 2
// 输出：
// [[2,4], [3,4], [2,3],[1,2], [1,3], [1,4]]
func combine(n int, k int) [][]int {

}

// 46. 全排列
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func permute(nums []int) [][]int {

}

// 78. 子集
// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 说明：解集不能包含重复的子集。
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
func subsets(nums []int) [][]int {
	var res [][]int
	var arr []int
	var dfs func(start int, arr []int)
	dfs = func(start int, arr []int) {
		res = append(res, append([]int{}, arr...))
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

// 90.子集II
// 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 说明：解集不能包含重复的子集。
// 示例:
// 输入: [1,2,2]
// 输出: [ [2], [1], [1,2,2], [2,2], [1,2], [] ]
func subsetsWithDup(nums []int) [][]int {
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
			// i >= 1
			if i > start && nums[i-1] == nums[i] {
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

func subsetsWithDupV2(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] < nums[j] {
			return true
		}
		return false
	})
	var res [][]int
	var arr []int
	var dfs func(int, []int)
	dfs = func(start int, arr []int) {
		res = append(res, append([]int{}, arr...))
		if start >= len(nums) {
			return
		}
		used := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			arr = append(arr, nums[i])
			dfs(i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(0, arr)
	return res
}

// 39. 组合总和
// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的数字可以无限制重复被选取。
// 输入：candidates = [2,3,6,7], target = 7
// 输出：[[2,2,3],[7]]
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var arr []int
	var dfs func(start, sum int, arr []int)
	dfs = func(start, sum int, arr []int) {
		if sum >= target {
			if sum == target {
				res = append(res, append([]int{}, arr...))
			}
			return
		}
		for i := start; i < len(candidates); i++ {
			arr = append(arr, candidates[i])
			dfs(i, sum+candidates[i], arr)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(0, 0, arr)
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
	rows := len(board)
	cols := len(board[0])
	var dfs func(row, col, index int) bool
	dfs = func(row, col, index int) bool {
		if row >= rows || row < 0 || col >= cols || col < 0 {
			return false
		}
		// 匹配上了
		if board[row][col] == word[index] && index == len(word)-1 {
			return true
		}

		// 匹配不上
		if board[row][col] != word[index] {
			return false
		}
		// 清空当前查找
		tmp := board[row][col]
		board[row][col] = 0
		// 继续搜索
		if dfs(row+1, col, index+1) ||
			dfs(row-1, col, index+1) ||
			dfs(row, col+1, index+1) ||
			dfs(row, col-1, index+1) {
			return true
		}
		// 恢复查找，提供后续使用
		board[row][col] = tmp
		return false
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// 131. 分割回文串
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
func partition(s string) [][]string {
	var res [][]string
	var dfs func(start int, arr []string)
	dfs = func(start int, arr []string) {
		if start == len(s) && len(arr) > 0 {
			res = append(res, append([]string{}, arr...))
		}
		for i := start; i < len(s); i++ {
			if isPalindromeStr(s[start : i+1]) {
				arr = append(arr, s[start:i+1])
				dfs(i+1, arr)
				arr = arr[:len(arr)-1]
			}
		}
	}
	var arr []string
	dfs(0, arr)
	return res
}

func isPalindromeStr(s string) bool {
	var i = 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 51. N 皇后
// 输入：n = 4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
func isQueens(row, col int, arr [][]byte) bool {

}
