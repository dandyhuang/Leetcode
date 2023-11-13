package leetcode_hot100

import (
	"fmt"
)

// 200. 岛屿数量
func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	if rows == 0 || cols == 0 {
		return 0
	}
	var dfs func([][]byte, int, int, int, int)
	dfs = func(grid [][]byte, row, col, rows, cols int) {
		if row < 0 || col < 0 || row >= rows || col >= cols || grid[row][col] == '0' {
			return
		}
		grid[row][col] = '0'
		dfs(grid, row-1, col, rows, cols)
		dfs(grid, row+1, col, rows, cols)
		dfs(grid, row, col-1, rows, cols)
		dfs(grid, row, col+1, rows, cols)
	}
	total := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				total++
				dfs(grid, i, j, rows, cols)
			}
		}
	}
	return total
}

// 994. 腐烂的橘子 bfs 层序遍历思想
func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	if rows == 0 || cols == 0 {
		return 0
	}
	var queue [][]int
	freshOranges := 0
	// 统计新鲜橘子的数量
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				freshOranges++
			} else if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	round := 0
	for freshOranges > 0 && len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			q := queue[0]
			queue = queue[1:]
			row := q[0]
			col := q[1]
			fmt.Println(row, col)
			if row-1 >= 0 && grid[row-1][col] == 1 {
				grid[row-1][col] = 2
				freshOranges--
				queue = append(queue, []int{row - 1, col})
			}
			if row+1 < rows && grid[row+1][col] == 1 {
				grid[row+1][col] = 2
				freshOranges--
				queue = append(queue, []int{row + 1, col})
			}
			if col-1 >= 0 && grid[row][col-1] == 1 {
				freshOranges--
				grid[row][col-1] = 2
				queue = append(queue, []int{row, col - 1})
			}
			if col+1 < cols && grid[row][col+1] == 1 {
				grid[row][col+1] = 2
				freshOranges--
				queue = append(queue, []int{row, col + 1})
			}
		}
		// 一轮过去了
		round++
	}
	if freshOranges > 0 {
		return -1
	} else {
		return round
	}
}

// 207. 课程表
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 每门课程的依赖关系, 一对多的关系，可能2，依赖1  [2,1]  [3,1], [2,3]
	// graph 为 1, [2,3],1没有依赖学完1。 1对应的2，和3，就可以减少对1的依赖了。 node--
	graph := make(map[int][]int, numCourses)
	// 每门课程依赖的数量
	node := make([]int, numCourses)
	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
		node[v[0]]++
	}
	var queue []int
	for i := 0; i < numCourses; i++ {
		// 可以学习的课程
		if node[i] == 0 {
			queue = append(queue, i)
		}
	}
	learned := 0
	for len(queue) > 0 {
		learned++
		q := queue[0]
		queue = queue[1:]
		// 遍历当前节点的邻接节点
		for _, v := range graph[q] {
			node[v]--
			if node[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return learned == numCourses
}
