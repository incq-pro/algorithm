package main

import "fmt"

func main() {
	for k := 1; k < 11; k++ {
		fmt.Println(getLeastNumbers([]int{4, 5, 1, 6, 2, 7, 3, 8, 6, 6, 6}, k))
	}
}

func getLeastNumbers(arr []int, k int) []int {
	l := 0
	r := len(arr) - 1
	kk := k - 1
	for l <= r {
		m := p2(arr, l, r)
		if m == kk {
			return arr[:m+1]
		} else if m > kk {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return nil
}

func p1(arr []int, l, r int) int {
	i := l + 1
	s := l
	pivot := arr[l]
	for ; i <= r; i++ {
		if arr[i] < pivot {
			s++
			if s != i {
				arr[s], arr[i] = arr[i], arr[s]
			}
		}
	}
	arr[l], arr[s] = arr[s], arr[l]
	return s
}

func p2(arr []int, l, r int) int {
	i := l
	j := r + 1
	pivot := arr[l]
	for {
		for i++; i <= r && arr[i] < pivot; i++ {
		}
		for j--; arr[j] > pivot; j-- {
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}
