package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 6, 7, 8, 9, 10}
	//for i := -1; i < 12; i++ {
	//	fmt.Println(BinarySearch2(nums, i))
	//}

	nums = []int{1, 2, 4, 5, 6, 7, 8}
	fmt.Println(FindLost(nums))
}

func BinarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (r-l)>>1 + l
		if target > nums[mid] {
			l = mid + 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func BinarySearch2(nums []int, target int) int {
	n := len(nums)
	l := 0
	r := n - 1
	for l < r {
		mid := (r-l)/2 + l
		if target <= nums[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l == r && nums[l] == target {
		return l
	}
	return -1
}

func FindLost2(nums []int) int {
	for i, x := range nums {
		if x-i > 1 {
			return i + 1
		}
	}
	return -1
}

func FindLost(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid]-mid == 1 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l + 1
}
