package hot100_2

// 200. 岛屿数量
// 输入：grid = [
//
//	["1","1","0","0","0"],
//	["1","1","0","0","0"],
//	["0","0","1","0","0"],
//	["0","0","0","1","1"]
//
// ]
// 输出：3
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
func numIslands(grid [][]byte) int {

}

// 994. 腐烂的橘子 bfs 层序遍历思想
// 在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
//
// 值 0 代表空单元格；
// 值 1 代表新鲜橘子；
// 值 2 代表腐烂的橘子。
// 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
//
// 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1
// 输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
// 输出：4
func orangesRotting(grid [][]int) int {
}

// 207. 课程表
// 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
// 输出：false
// 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
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
