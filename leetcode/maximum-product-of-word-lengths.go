package main

func main() {
	maxProduct([]string{"abcdefgg", "aaaaaaa", "bbbbbbb", "ccccccc", "dddddd", "fefewfwe"})
}


func maxProduct(words []string) int {
	masks := make(map[int]int, len(words))
	for _, w := range words {
		mask := 0
		for _, c := range w {
			mask |= 1 << int(c - 'a')
		}
		if masks[mask] < len(w) {
			masks[mask] = len(w)
		}
	}
	maxLen := 0
	for x, lenX := range masks {
		delete(masks, x)
		for y, lenY := range masks {
			if (x & y) == 0 && lenX * lenY > maxLen {
				maxLen = lenX * lenY
			}
		}
	}
	return maxLen
}
