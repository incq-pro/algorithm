package main

import "fmt"

func main() {
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
	fmt.Println(getHint("1", "0"))
	fmt.Println(getHint("1", "1"))
}

func getHint(secret string, guess string) string {
	s := []byte(secret)
	g := []byte(guess)
	a := 0
	gg := []byte{}
	count := make(map[byte]int, len(s))
	for i, c := range s {
		if c == g[i] {
			a++
		} else {
			gg = append(gg, g[i])
			count[c]++
		}
	}
	b := 0
	for _, c := range gg {
		if count[c] > 0 {
			count[c]--
			b++
		}
	}
	return fmt.Sprintf("%dA%dB", a, b)
}
