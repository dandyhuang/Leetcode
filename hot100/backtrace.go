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

// 79. 单词搜索 和岛屿数量是一样的
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

// 131. 分割回文串
func partition(s string) [][]string {
	var res [][]string
	var arr []string
	var dfs func(s string, start int)
	dfs = func(s string, start int) {
		if start == len(s) {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(s); i++ {
			fmt.Println(s[start : i+1])
			if isPalindromeStr(s[start : i+1]) {
				arr = append(arr, s[start:i+1])
				dfs(s, i+1)
				arr = arr[:len(arr)-1]
			}
		}
	}
	dfs(s, 0)
	return res
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
