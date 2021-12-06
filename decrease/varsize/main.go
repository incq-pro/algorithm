package main

import "fmt"

func main() {
	//for k := 2; k < 7; k++ {
	//	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0}
	//	fmt.Println(k, QuickSelect(nums, k))
	//}
	fmt.Println(InterpolationSearch(nil, -1))
	nums := []int{0, 1, 4, 7, 9, 10, 12, 16, 17, 18, 100}
	for _, target := range []int{0, 1, 2, 17, 18, 100, 101, 200} {
		fmt.Println(target, InterpolationSearch(nums, target))
	}
}

func QuickSelect(nums []int, k int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		k1 := LomutoPartition(nums, l, r)
		if k > k1 {
			l = k1 + 1
		} else if k < k1 {
			r = k1 - 1
		} else {
			return k1
		}
	}
	return -1
}

func LomutoPartition(nums []int, l, r int) int {
	s := l
	pivot := nums[l]
	for i := l + 1; i <= r; i++ {
		if nums[i] < pivot {
			s++
			if i != s {
				nums[s], nums[i] = nums[i], nums[s]
			}
		}
	}
	nums[l], nums[s] = nums[s], nums[l]
	return s
}

// InterpolationSearch search target in nums, return -1 if not found
func InterpolationSearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	v := 0
	for l < r {
		v++
		if target < nums[l] || target > nums[r] {
			return -1
		}
		x := l + (target-nums[l])*(r-l)/(nums[r]-nums[l])
		if target > nums[x] {
			l = x + 1
		} else if target < nums[x] {
			r = x - 1
		} else {
			return x
		}
	}
	if l == r && target == nums[l] {
		return l
	}
	return -1
}
