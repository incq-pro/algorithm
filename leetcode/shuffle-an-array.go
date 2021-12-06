package main

import "math/rand"

func main() {

}

type Solution struct {
	nums []int
}

func NewSolution(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

func (s *Solution) Reset() []int {
	return s.nums
}

func (s *Solution) Shuffle() []int {
	ss := make([]int, len(s.nums))
	copy(ss, s.nums)
	rand.Shuffle(len(ss), func(i, j int) {
		ss[i], ss[j] = ss[j], ss[i]
	})
	return ss
}
