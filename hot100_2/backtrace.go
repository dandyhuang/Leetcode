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
func combinationSum(nums []int, target int) [][]int {
	var arr []int
	var res [][]int
	var dfs func(sum, start int, arr []int)
	dfs = func(sum, start int, arr []int) {
		if sum > target {
			return
		}
		if sum == target {
			res = append(res, append([]int{}, arr...))
			return
		}
		for i := start; i < len(nums); i++ {
			arr = append(arr, nums[i])
			// 这里因为可以重复，从i开始
			dfs(sum+nums[i], i, arr)
			arr = arr[:len(arr)-1]
		}
		return
	}
	dfs(0, 0, arr)
	return res
}

// 17. 电话号码的字母组合
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
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
	var dfs func(start int, arr []byte)
	dfs = func(start int, arr []byte) {
		if len(arr) == len(digits) {
			tmp := make([]byte, len(arr))
			copy(tmp, arr)
			res = append(res, string(tmp))
			return
		}
		letters := m[digits[start]]
		for i := 0; i < len(letters); i++ {
			arr = append(arr, letters[i])
			dfs(start+1, arr)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(0, arr)
	return res
}

// 22. 括号生成
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
func generateParenthesis(n int) []string {
	var res []string
	var s []byte
	var dfs func(l, r int, s []byte)
	dfs = func(l, r int, s []byte) {
		if l == n && r == n {
			tmp := make([]byte, len(s))
			copy(tmp, s)
			res = append(res, string(s))
			return
		}
		if l < n {
			s = append(s, '(')
			dfs(l+1, r, s)
			s = s[:len(s)-1]
		}
		if r < l {
			s = append(s, ')')
			dfs(l, r+1, s)
			s = s[:len(s)-1]
		}
	}
	dfs(0, 0, s)
	return res
}
func isPalindromeStr(str string) bool {
	l, r := 0, len(str)-1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// 131. 分割回文串
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
func partition(s string) [][]string {
	var res [][]string
	var arr []string
	var dfs func(start int, arr []string)
	dfs = func(start int, arr []string) {
		if start == len(s) {
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
	dfs(0, arr)
	return res
}

// 79. 单词搜索 和岛屿数量是一样的 pass
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

// 51. N 皇后 pass
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
