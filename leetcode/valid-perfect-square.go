package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(1))
	fmt.Println(isPerfectSquare(2))
	fmt.Println(isPerfectSquare(3))
	fmt.Println(isPerfectSquare(4))
	fmt.Println(isPerfectSquare(14))
	fmt.Println(isPerfectSquare(16))
}


func isPerfectSquare(num int) bool {
	lo, hi := 0, num
	for lo <= hi {
		mid := lo + (hi - lo) / 2
		midSq := mid * mid
		if midSq > num {
			hi = mid - 1
		} else if midSq < num {
			lo = mid + 1
		} else {
			return true
		}
	}
	return false
}