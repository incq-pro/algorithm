package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxSubArray([]int{1, 2, -3, 4, 5}))
	fmt.Println(maxSubArray([]int{1, 2, -3, -4, 5}))
	fmt.Println(maxSubArray([]int{1, 2, -3, -4, -5}))
	fmt.Println(maxSubArray([]int{-1}))
	fmt.Println(maxSubArray([]int{-1, 1}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	return maxSubArray2(nums)
}

func maxSubArray1(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	result := math.MinInt
	s := 0
	for _, num := range nums {
		s += num
		if s > result {
			result = s
		}
		if s < 0 {
			s = 0
		}
	}
	return result
}

func doMaxSubArray(nums []int, l, r int) (int, int, int, int) {
	if l > r {
		panic(fmt.Sprintf("l > r: %d > %d", l, r))
	} else if l == r {
		return nums[l], nums[l], nums[l], nums[l]
	}
	mid := l + (r-l)/2
	m1, l1, r1, a1 := doMaxSubArray(nums, l, mid)
	m2, l2, r2, a2 := doMaxSubArray(nums, mid+1, r)
	m := max(max(m1, m2), r1+l2)
	ll := max(l1, a1+l2)
	rr := max(r2, a2+r1)
	a := a1 + a2
	return m, ll, rr, a
}

func maxSubArray2(nums []int) int {
	m, _, _, _ := doMaxSubArray(nums, 0, len(nums)-1)
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
