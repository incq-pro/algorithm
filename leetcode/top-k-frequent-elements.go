package main

import (
	"fmt"
	"math/rand"
	"time"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	counts := make([][2]int, 0, len(m))
	for k, v := range m {
		counts = append(counts, [2]int{k, v})
	}
	j := topK(counts, k-1)
	result := make([]int, 0, k)
	for i := 0; i <= j; i++ {
		result = append(result, counts[i][0])
	}
	return result
}

func topK(nums [][2]int, k int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		pivot := nums[l]
		i := l
		j := r + 1
		for {
			for i++; i <= r && nums[i][1] > pivot[1]; i++ {
			}
			for j--; j >= l && nums[j][1] < pivot[1]; j-- {
			}
			if i >= j {
				break
			}
			// swap
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[l], nums[j] = nums[j], pivot
		if j < k {
			l = j + 1
		} else if j > k {
			r = j - 1
		} else {
			return j
		}
	}
	return -1
}

func topK2(nums []int, k int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		pivot := nums[r]
		i := l
		for j := l; j < r; j++ {
			if nums[j] < pivot {
				if i != j {
					nums[i], nums[j] = nums[j], nums[i]
				}
				i++
			}
		}
		nums[i], nums[r] = pivot, nums[i]
		if i > k {
			r = i - 1
		} else if i < k {
			l = i + 1
		} else {
			return nums[i]
		}
	}
	return 0
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	fmt.Println(nums)
	for k := 0; k < len(nums); k++ {
		fmt.Println(k, topK2(nums, k))
	}
}
