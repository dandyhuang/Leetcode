package leetcode_hot100

import (
	"sort"
	"strconv"
)

// 组合排序
// 77. 组合
// 输入：n = 4, k = 2
// 输出：
// [[2,4], [3,4], [2,3],[1,2], [1,3], [1,4]]
func combine(n int, k int) [][]int {
	var res [][]int
	var arr []int
	var dfs func(int, []int)
	dfs = func(start int, arr []int) {
		if len(arr)+n-start+1 < k {
			return
		}
		if len(arr) == k {
			res = append(res, append([]int{}, arr...))
			// 直接赋值有问题  拷贝的都是最后一个值
			// res = append(res, arr)
			return
		}
		for i := start; i <= n; i++ {
			arr = append(arr, i)
			dfs(i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(1, arr)
	return res
}

// 216.组合总和III
// 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
// 示例 1: 输入: k = 3, n = 7 输出: [[1,2,4]]
// 示例 2: 输入: k = 3, n = 9 输出: [[1,2,6], [1,3,5], [2,3,4]]

// 39. 组合总和
// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的数字可以无限制重复被选取。
// 输入：candidates = [2,3,6,7], target = 7
// 输出：[[2,2,3],[7]]
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var arr []int
	var dfs func(int, int, []int)
	dfs = func(sum, index int, arr []int) {
		if sum > target {
			return
		}
		if sum == target {
			res = append(res, append([]int{}, arr...))
			return
		}

		for i := index; i < len(candidates); i++ {
			arr = append(arr, candidates[i])
			dfs(sum+candidates[i], i, arr)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(0, 0, arr)
	return res
}

// 40.组合总和II
// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用一次。
// 说明： 所有数字（包括目标数）都是正整数。解集不能包含重复的组合。
// 示例 1:
// 输入: candidates = [10,1,2,7,6,1,5], target = 8, // 1,1,2,5,6,7,10
// 输出:
// [  [1,1,6],  [1,2,5], [1,7], [2,6] ]
func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i] < candidates[j] {
			return true
		}
		return false
	})
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
			dfs(i+1, sum+candidates[i], arr)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(0, 0, arr)
	return res
}

// 17. 电话号码的字母组合
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}
	m := map[byte]string{
		'1': "",
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var dfs func(start int, str string)
	dfs = func(start int, str string) {
		if len(str) == len(digits) {
			res = append(res, str)
			return
		}
		for i := 0; i < len(m[digits[start]]); i++ {
			dfs(start+1, str+string(m[digits[start]][i]))
		}
	}
	dfs(0, "")
	return res
}

// 131. 分割回文串
// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文串。返回 s 所有可能的分割方案。
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
func partition(s string) [][]string {
	var res [][]string
	var arr []string
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(s) {
			res = append(res, append([]string{}, arr...))
			return
		}
		// 如abcd，切割a后， 在从bcd开始切割...
		for i := start; i < len(s); i++ {
			if isPalindromeStr(s[start : i+1]) {
				arr = append(arr, s[start:i+1])
				dfs(i + 1)
				arr = arr[:len(arr)-1]
			}
		}
	}
	dfs(0)
	return res
}

func isPalindromeStr(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 93. 复原 IP 地址 分隔
// 输入：s = "101023"
// 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
func restoreIpAddresses(s string) []string {
	var res []string
	var dfs func(start int, arr []string)
	dfs = func(start int, arr []string) {
		if start == len(s) && len(arr) == 4 {
			res = append(res, arr[0]+"."+arr[1]+"."+arr[2]+"."+arr[3])
			return
		}
		for i := start; i < len(s); i++ {
			if isNormalIp(s, start, i) {
				arr = append(arr, s[start:i+1])
				// 大于三位数的ip
				if len(arr) > 4 || i-start+1 > 3 {
					return
				}
				dfs(i+1, arr)
				arr = arr[:len(arr)-1]
			}
		}
	}
	var arr []string
	dfs(0, arr)
	return res
}

func isNormalIp(s string, start, end int) bool {
	// 01等 情况排出
	if end-start+1 > 1 && s[start] == '0' {
		return false
	}
	num, _ := strconv.Atoi(s[start : end+1])
	if num > 255 {
		return false
	}
	return true
}

// 78. 子集
// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 说明：解集不能包含重复的子集。
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
			// i - 1 >= start
			if i >= 1 && nums[i-1] == nums[i] {
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

// 46. 全排列
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
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

// 79. 单词搜索 和岛屿数量是一样的
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// 输出：true
func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}
	var dfs func(board [][]byte, word string, row, col, rows, cols, index int, visited [][]bool) bool
	dfs = func(board [][]byte, word string, row, col, rows, cols, index int, visited [][]bool) bool {
		if index == len(word) {
			return true
		}
		if row < 0 || row >= rows || col < 0 || col >= cols ||
			visited[row][col] || board[row][col] != word[index] {
			return false
		}
		visited[row][col] = true
		if dfs(board, word, row+1, col, rows, cols, index+1, visited) ||
			dfs(board, word, row-1, col, rows, cols, index+1, visited) ||
			dfs(board, word, row, col+1, rows, cols, index+1, visited) ||
			dfs(board, word, row, col-1, rows, cols, index+1, visited) {
			return true
		}
		visited[row][col] = false
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, rows, cols, 0, visited) {
				return true
			}
		}
	}
	return false
}

// 51. N 皇后
// 输入：n = 4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
func isQueens(row, col int, arr [][]byte) bool {
	n := len(arr)
	// 检查同一列是否有皇后
	// 因为我们是每行每行的放置，因此当前行数以下都没有放置皇后所以扫描当前行数以上的就行了
	for i := 0; i < row; i++ {
		if arr[i][col] == 'Q' {
			return false
		}
	}
	// 检查左上方是否有皇后, 以下都还没放,后续增加也会检查到
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if arr[i][j] == 'Q' {
			return false
		}
	}

	// 检查右上方是否有皇后
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if arr[i][j] == 'Q' {
			return false
		}
	}
	return true
}
func solveNQueens(n int) [][]string {
	var res [][]string
	arr := make([][]byte, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]byte, n)
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = '.'
		}
	}

	var dfs func(n, row int)
	dfs = func(n, row int) {
		if row == n {
			tmp := make([]string, 0)
			for i := 0; i < n; i++ {
				tmp = append(tmp, string(arr[i]))
			}
			res = append(res, tmp)
		}
		for col := 0; col < n; col++ {
			if isQueens(row, col, arr) {
				arr[row][col] = 'Q'
				// 递归每一行
				dfs(n, row+1)
				arr[row][col] = '.'
			}
		}
	}
	dfs(n, 0)
	return res
}
