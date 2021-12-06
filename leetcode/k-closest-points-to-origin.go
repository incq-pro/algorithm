package main

import "fmt"

func main() {
	points := [][]int{
		{3, 3},
		{5, -1},
		{-2, 4},
	}
	for k := 1; k <= 3; k++ {
		fmt.Println(kClosest(points, k))
	}
}

func kClosest(points [][]int, k int) [][]int {
	points2 := make([][]int, len(points))
	for i, point := range points {
		points2[i] = []int{point[0]*point[0] + point[1]*point[1], i}
	}
	kk := k - 1
	l := 0
	r := len(points2) - 1
	for l <= r {
		m := partition(points2, l, r)
		if m > kk {
			r = m - 1
		} else if m < kk {
			l = m + 1
		} else {
			break
		}
	}
	result := make([][]int, k)
	for i, x := range points2[:k] {
		result[i] = points[x[1]]
	}
	return result
}

func partition(arr [][]int, l, r int) int {
	pivot := arr[l]
	i := l
	j := r + 1
	for {
		for i++; i <= r && arr[i][0] < pivot[0]; i++ {
		}
		for j--; arr[j][0] > pivot[0]; j-- {
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}
