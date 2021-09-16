package bruteforce

import "testing"

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			"empty",
			[]int{},
		},
		{
			"one",
			[]int{1},
		},
		{
			"two",
			[]int{2, 1},
		},
		{
			"three",
			[]int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.nums)
			t.Log(tt.nums)
		})
	}
}

func TestBubbleSort(t *testing.T) {

	tests := []struct {
		name string
		nums []int
	}{
		{
			"empty",
			[]int{},
		},
		{
			"one",
			[]int{1},
		},
		{
			"two",
			[]int{2, 1},
		},
		{
			"three",
			[]int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.nums)
			t.Log(tt.nums)
		})
	}
}
