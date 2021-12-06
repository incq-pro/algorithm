package main

import "fmt"

func main() {
	nums := []int{6, 5, 7, 4, 8, 3, 9, 2, 10, 1, 0, -1}
	QuickSort(nums)
	fmt.Println(nums)
}

func QuickSort(nums []int) {
	doQuickSort(nums, 0, len(nums)-1)
}

func doQuickSort(nums []int, r, l int) {
	for r < l {
		k := partition(nums, r, l)
		doQuickSort(nums, k+1, l)
		l = k - 1
	}
}

func partition(nums []int, r, l int) int {
	p := nums[r]
	i := r + 1
	j := l
	for {
		for ; i <= l && nums[i] < p; i++ {
		}
		for ; j >= r && nums[j] > p; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			break
		}
	}
	nums[r], nums[j] = nums[j], p
	return j
}
