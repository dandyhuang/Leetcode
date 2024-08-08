package main

import "fmt"

func main() {
	n := 5 // 假设节点编号从 1 到 4
	// m := 4 // 边的数量

	// 初始化邻接表和入度数组
	adjList := make([][]int, n)
	inDegree := make([]int, n)

	// 图的边
	edges := [][2]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{3, 4},
	}

	// 填充邻接表和入度数组
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adjList[a] = append(adjList[a], b)
		inDegree[b]++
	}

	// 输出邻接表
	for i := 0; i < n; i++ {
		fmt.Printf("Node %d: %v\n", i, adjList[i])
	}
	fmt.Println("indegree:", inDegree)
	// 拓扑排序
	queue := []int{}
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	fmt.Println("queue:", queue)
}
