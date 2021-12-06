package main

import (
	"fmt"
)

func main() {
	graph := [][]int{
		{0, 1, 1, 1, 1, 0, 0},
		{1, 0, 0, 1, 0, 1, 0},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 1, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0},
	}
	DFS(graph)

	graph2 := []map[int]bool{
		{1: true, 2: true},
		{0: true, 2: true, 3: true},
		{0: true, 1: true, 3: true, 4: true},
		{4: true, 1: true},
		{2: true, 3: true},
	}
	BFS(graph2)
}

func DFS(graph [][]int) {
	n := len(graph)
	visited := map[int]bool{}
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(graph, i, visited)
		}
	}
}

func dfs(graph [][]int, v int, visited map[int]bool) {
	visited[v] = true
	fmt.Println(v)
	for i := 0; i < len(graph[v]); i++ {
		if graph[v][i] > 0 && !visited[i] {
			dfs(graph, i, visited)
		}
	}
}

func BFS(graph []map[int]bool) {
	n := len(graph)
	visited := map[int]bool{}
	for i := 0; i < n; i++ {
		if !visited[i] {
			bfs(graph, i, visited)
		}
	}
}

func bfs(graph []map[int]bool, v int, visited map[int]bool) {
	visited[v] = true
	queue := []int{v}

	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		fmt.Println("BFS", x)
		for vv, _ := range graph[x] {
			if !visited[vv] {
				visited[vv] = true
				queue = append(queue, vv)
			}
		}
	}
}
