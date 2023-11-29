package hot100_2

// 73. 矩阵置零
// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0
// 输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
// 输出：[[1,0,1],[0,0,0],[1,0,1]]
func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	var res [][]int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				res = append(res, []int{i, j})
			}
		}
	}

	for i := range res {
		row, col := res[i][0], res[i][1]
		for j := 0; j < cols; j++ {
			matrix[row][j] = 0
		}
		for j := 0; j < rows; j++ {
			matrix[j][col] = 0
		}
	}
}

// 54. 螺旋矩阵
// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]
func spiralOrder(matrix [][]int) []int {

}

// 48. 旋转图像
// 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[[7,4,1],[8,5,2],[9,6,3]]
func mRotate(matrix [][]int) {

}

// 240. 搜索二维矩阵 II
// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
func searchMatrix(matrix [][]int, target int) bool {

	return false
}
