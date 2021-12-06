package main

import "fmt"

func main() {
	nums1 := []int{1, 5, 1, 1, 6, 4, 7}
	wiggleSort(nums1)
	fmt.Println(nums1)

	nums2 := []int{4,5,5,6}
	wiggleSort(nums2)
	fmt.Println(nums2)
}

func wiggleSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	h := splitHalf(nums)
	fmt.Println("1:", h, nums)
	move(nums, h)
	fmt.Println("2:", h, nums)
	compose(nums, h)
	fmt.Println("3:", h, nums)
}

func splitHalf(nums []int) int {
	k := len(nums) / 2
	if len(nums)%2 == 0 {
		k = k - 1
	}
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := partation(nums, l, r)
		if m > k {
			r = m - 1
		} else if m < k {
			l = m + 1
		} else {
			break
		}
	}
	return k
}

func partation(nums []int, l, r int) int {
	i := l
	j := r + 1
	pivot := nums[l]
	for {
		for i++; i <= r && nums[i] < pivot; i++ {
		}
		for j--; nums[j] > pivot; j-- {
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}

func move(nums []int, x int) {
	pivot := nums[x]
	j := x - 1
	for i := x - 1; i >= 0; i-- {
		if nums[i] == pivot {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}
	j = x + 1
	for i := x + 1; i < len(nums); i++ {
		if nums[i] == pivot {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func compose(nums []int, x int) {
	nums2 := make([]int, len(nums))
	copy(nums2, nums)
	n1 := nums2[:x + 1]
	n2 := nums2[x + 1:]
	i := 0
	for k := len(n1) - 1; k >= 0; k-- {
		nums[i*2] = n1[k]
		i++
	}
	i = 0
	for k := len(n2) - 1; k >= 0 ;k-- {
		nums[i*2+1] = n2[k]
		i++
	}
}
