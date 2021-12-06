package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7, 2, 4, 6, 8, 9, 10}
	InsertionSort(nums)
	fmt.Println(nums)
}

func InsertionSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		v := nums[i]
		j := i - 1
		for ; j >= 0 && nums[j] > v; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = v
	}
}
