package main

import (
	"fmt"
	"testing"
)

func TestWiggleSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "haha",
			nums: []int{1, 2, 3, 4},
		},
		{
			name: "h1",
			nums: []int{1},
		},
		{
			name: "h2",
			nums: []int{1, 2},
		},
		{
			name: "h3",
			nums: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WiggleSort(tt.nums)
			fmt.Println(tt.nums)
		})
	}
}
