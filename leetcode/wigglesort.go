package main

import "sort"

func WiggleSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	sort.Ints(nums)
	mid := (n + 1) / 2
	nums2 := make([]int, n)
	for i := 0; i < mid; i++ {
		nums2[i*2] = nums[i]
		if 2*i+1 < n {
			nums2[i*2+1] = nums[n-i-1]
		}
	}
	for i := 0; i < n; i++ {
		nums[i] = nums2[i]
	}
}

type xx []int

func (a xx) Len() int {
	return len(a)
}
