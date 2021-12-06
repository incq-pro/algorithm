package main

import "errors"

func topoReduce(graph [][]int) ([]int, error) {
	n := len(graph)
	count := make(map[int]int, n)
	nodes := make([]int, n)
	for i, vs := range graph {
		nodes[i] = i
		for _, v := range vs {
			count[v]++
		}
	}
	var sort []int
	for {
		found := false
		i := 0
		v := 0
		for i, v = range nodes {
			if count[v] <= 0 {
				found = true
				break
			}
		}
		if !found {
			break
		}
		sort = append(sort, v)
		nodes[i] = nodes[len(nodes)-1]
		nodes = nodes[0 : len(nodes)-1]
		for _, vv := range graph[v] {
			count[vv]--
		}
	}
	if len(sort) < n {
		return nil, errors.New("有坏！")
	}
	// 返回结果
	return sort, nil
}
