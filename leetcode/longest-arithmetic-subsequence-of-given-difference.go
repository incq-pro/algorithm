package main

import "fmt"

func main() {
	longestSubsequence([]int{1, 2, 3, 4}, 1)
	longestSubsequence([]int{1, 3, 5, 7}, 1)
	longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2)
	longestSubsequence([]int{1, 5, 7, 1, 1, 3, 4, 2, 1}, 0)
}

func longestSubsequence(arr []int, difference int) int {
	max := 0
	counter := make(map[int]int, len(arr))
	for _, v := range arr {
		c := counter[v-difference] + 1
		counter[v] = c
		if c > max {
			max = c
		}
	}
	return max
}


func longestSubsequence2(arr []int, difference int) int {
	max := 0
	for i, v := range arr {
		l := 1
		for _, v2 := range arr[i:] {
			if v2-v == difference {
				l++
				v = v2
			}
		}
		if l > max {
			max = l
		}
	}
	fmt.Println(arr, difference, max)
	return max
}
