package main

import "fmt"

func main() {
	prerequisites := [][]int{{1, 3}, {2, 4}}
	graph := make(map[int][]int, 2)
	for _, prerequisite := range prerequisites {
		course, prereq := prerequisite[0], prerequisite[1]
		graph[prereq] = append(graph[prereq], course)
	}
	for k, v := range graph[3] {
		fmt.Println(k, v)
	}
}
