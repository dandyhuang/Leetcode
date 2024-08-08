package main

import (
	"fmt"
)

func main() {
	var n, m int
	// 表示有 n 个节点，m 条关系描述
	fmt.Scan(&n, &m)

	// 邻接表和入度数组
	adjList := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		adjList[a] = append(adjList[a], b)
	}
	fmt.Println("adjList:", adjList)
	// 找结束节点
	endingNodes := []int{}
	for i := 0; i < n; i++ {
		if len(adjList[i]) == 0 {
			endingNodes = append(endingNodes, i)
		}
	}
	fmt.Println("endingNodes:", endingNodes)

	if len(endingNodes) == 1 {
		fmt.Println(endingNodes[0])
	} else {
		maxDependencies := 0
		resultNode := 0
		dependencies := make([]int, n)
		for i := 0; i < n; i++ {
			for _, neighbor := range adjList[i] {
				dependencies[neighbor]++
			}
		}
		for i := 0; i < n; i++ {
			if dependencies[i] > maxDependencies || (dependencies[i] == maxDependencies && i > resultNode) {
				maxDependencies = dependencies[i]
				resultNode = i
			}
		}
		fmt.Println(resultNode)
	}
}
