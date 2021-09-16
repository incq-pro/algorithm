package main

import "fmt"

const m = 1337

func superPow(a int, b []int) int {
	if a > m {
		a = a % m
	}
	if a == 0 || a == 1 {
		return a
	}
	if len(b) == 0 {
		return 1
	}
	c, r := divideSlice(b, 2)
	if r != 0 {
		return (superPow(a*a, c) * a) % m
	} else {
		return superPow(a*a, c) % m
	}
}

func divideSlice(s []int, factor int) ([]int, int) {
	firstNonZero := 0
	foundNonZero := true
	left := 0
	for i := 0; i < len(s); i++ {
		value := s[i] + left*10
		v := value / factor
		left = value % factor
		s[i] = v
		if v != 0 {
			foundNonZero = false
		} else {
			if foundNonZero {
				firstNonZero = i + 1
			}
		}
	}
	return s[firstNonZero:], left
}

func main() {
	fmt.Println(superPow(2, []int{1, 0, 0}))
	fmt.Println(superPow(3, []int{1, 0, 1}))
	fmt.Println(superPow(5, []int{1, 0, 2}))
	fmt.Println(superPow(6, []int{1, 0, 3}))
	fmt.Println(superPow(7, []int{1, 0, 4}))
}
