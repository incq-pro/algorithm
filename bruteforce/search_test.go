package bruteforce

import (
	"incq.pro/golang-algo/common"
	"testing"
)

func TestBytesMatch(t *testing.T) {
	tests := []struct {
		name string
		t    []byte
		p    []byte
		want int
	}{
		{
			name: "empty",
			t:    []byte(""),
			p:    []byte(""),
			want: 0,
		},
		{
			name: "2",
			t:    []byte{},
			p:    []byte{},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesMatch(tt.t, tt.p); got != tt.want {
				t.Errorf("BytesMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClosestPoints(t *testing.T) {

	tests := []struct {
		name   string
		points []*common.Point
		want   float64
	}{
		{
			"1",
			[]*common.Point{{0, 0}, {3, 4}, {5, 12}},
			5.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClosestPoints(tt.points); got != tt.want {
				t.Errorf("ClosestPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinAvgDistance(t *testing.T) {

	tests := []struct {
		name string
		x    []float64
		want float64
	}{
		{
			"2",
			[]float64{0.0, 1.0},
			0.0,
		},
		{
			"2",
			[]float64{0.0, 1.0, 2.0},
			0.0,
		},
		{
			"2",
			[]float64{0.0, 1.0, 2.0, 1000.0},
			0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, i := MinAvgDistance(tt.x); got != tt.want {
				t.Logf("MinAvgDistance() = %v, index %v", got, i)
			}
		})
	}
}

func TestConvexHull(t *testing.T) {
	tests := []struct {
		name   string
		points []*common.Point
		want   []*common.Point
	}{
		{
			"1",
			[]*common.Point{{1, 1}, {2, 3}, {3, 3}, {4, 3.5}, {5, 0.5}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvexHull(tt.points)
			for _, x := range got {
				t.Log(x)
			}
		})
	}
}
