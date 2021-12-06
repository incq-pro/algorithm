package main

import "fmt"

func main() {
	x := permutation(4)
	fmt.Println(len(x), x)

	y := BRGC(4)
	for _, yy := range y {
		fmt.Printf("%04b ", yy)
	}
	fmt.Println()
}

func permutation(n int) [][]int {
	if n == 0 {
		return nil
	}
	var result = [][]int{{0}}
	for i := 1; i < n; i++ {
		direction := true
		var next [][]int
		for _, pre := range result {
			if direction {
				direction = false
				pre = append(pre, i)
				order := make([]int, len(pre))
				copy(order, pre)
				next = append(next, order)
				for j := len(pre) - 2; j >= 0; j-- {
					pre[j], pre[j+1] = pre[j+1], pre[j]
					order = make([]int, len(pre))
					copy(order, pre)
					next = append(next, order)
				}
			} else {
				direction = true
				order := make([]int, 0, len(pre) + 1)
				order = append(order, i)
				order = append(order, pre...)
				next = append(next, order)
				pre = make([]int, len(order))
				copy(pre, order)
				for j := 0; j < len(pre)-1; j++ {
					pre[j], pre[j+1] = pre[j+1], pre[j]
					order = make([]int, len(pre))
					copy(order, pre)
					next = append(next, order)
				}
			}
		}
		result = next
	}
	return result
}

func BRGC(n int) []int {
	result := make([]int, 0, 1<<n)
	result = append(result, 0)
	for i := 1; i <= n; i++ {
		x := 1 << (i - 1)
		for j := x - 1; j >= 0; j-- {
			result = append(result, result[j]|x)
		}
	}
	return result
}