package main

import (
	"errors"
	"fmt"
	"strconv"
)

func topoSortDFS(graph [][]int) ([]int, error) {
	n := len(graph)
	inDegreeCount := map[int]int{}
	for _, vs := range graph {
		for _, v := range vs {
			inDegreeCount[v]++
		}
	}
	var starter []int
	for i := 0; i < n; i++ {
		if inDegreeCount[i] <= 0 {
			starter = append(starter, i)
		}
	}
	if len(starter) == 0 {
		return nil, errors.New("没有找到起点！")
	}
	fmt.Println(starter)
	var path []int
	visited := map[int]bool{}
	var err error
	history := map[int]bool{}
	for _, start := range starter {
		path, err = dfs(graph, start, visited, path, history)
	}
	return path, err
}

func dfs(graph [][]int, v int, visited map[int]bool, order []int, history map[int]bool) ([]int, error) {
	visited[v] = true
	history[v] = true
	var err error
	for _, vs := range graph[v] {
		if history[vs] {
			return nil, errors.New("有环" + strconv.Itoa(v) + " -> " + strconv.Itoa(vs))
		}
		if visited[vs] {
			continue
		}
		order, err = dfs(graph, vs, visited, order, history)
		if err != nil {
			return order, err
		}
	}
	order = append(order, v)
	delete(history, v)
	return order, nil
}
