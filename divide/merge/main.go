package main

import (
	"fmt"
)

func main() {
	nums := []int{6, 5, 7, 4, 8, 3, 9, 2, 0, 1, -1}
	MergeSort(nums)
	fmt.Println(nums)
}

func MergeSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
}
