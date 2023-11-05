package leetcode_hot100

import "sort"

// 73. 矩阵置零
func setZeroes(matrix [][]int) {
	zeroRow, zeroClos := make(map[int]bool), make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				zeroRow[i] = true
				zeroClos[j] = true
			}
		}
	}
	for row := range zeroRow {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[row][j] = 0
		}
	}
	for clos := range zeroClos {
		for i := 0; i < len(matrix); i++ {
			matrix[i][clos] = 0
		}
	}
}

// 54. 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	var res []int
	for l <= r && t <= b {
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		t++
		if t > b {
			break
		}
		for i := t; i <= b; i++ {
			res = append(res, matrix[i][r])
		}
		r--
		if r < l {
			break
		}
		for i := r; i >= l; i-- {
			res = append(res, matrix[b][i])
		}
		b--
		for i := b; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		l++
	}
	return res
}

// 48. 旋转图像
func mRotate(matrix [][]int) {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp[j][n-i-1] = matrix[i][j]
		}
	}
	copy(matrix, tmp)
}

// 240. 搜索二维矩阵 II
func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		j := sort.SearchInts(matrix[i], target)
		if j < len(matrix[i]) && matrix[i][j] == target {
			return true
		}
	}
	return false
}
