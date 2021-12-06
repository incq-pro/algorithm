package main

import "fmt"

func main() {
	results := [][]int{
		{0, 1, 0, -1},
		{-1, 0, -1, 1},
		{0, 1, 0, -1},
		{1, -1, 1, 0},
	}
	tidyResult(results)

	order := teamSort(results)
	fmt.Println(checkTeamSort(results, order), order)
}

func tidyResult(result [][]int) {
	n := len(result)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			result[j][i] = -result[i][j]
		}
	}
}
func teamSort(results [][]int) []int {
	n := len(results)
	order := make([]int, n)
	for i := 0; i < n; i++ {
		order[i] = i
	}
	for i := 1; i < n; i++ {
		j := i - 1
		for ; j >= 0 && results[i][order[j]] >= 0; j-- {
			order[j+1] = order[j]
		}
		order[j+1] = i
	}
	return order
}

func checkTeamSort(results [][]int, order []int) bool {
	n := len(results)
	for i := 0; i < n-1; i++ {
		a := order[i]
		b := order[i+1]
		if results[a][b] < 0 {
			return false
		}
	}
	return true
}
