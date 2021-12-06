package main

import (
	"fmt"
)

func main() {
	graph1 := [][]int{
		0: {1, 2},
		1: {4, 6},
		2: {5},
		3: {0, 1, 2, 5, 6},
		4: {},
		5: {},
		6: {4, 5},
	}
	fmt.Println(topoSortDFS(graph1))
	fmt.Println(topoReduce(graph1))

	graph2 := [][]int{
		0: {1},
		1: {2},
		2: {3},
		3: {6},
		4: {0},
		5: {1, 2, 4, 6},
		6: {4},
	}
	fmt.Println(topoSortDFS(graph2))
	fmt.Println(topoReduce(graph2))
}
