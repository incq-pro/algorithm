package bruteforce

import (
	"incq.pro/golang-algo/common"
	"math"
)

// SequentialSearch 查找k, 返回index
func SequentialSearch(nums []int, k int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == k {
			return i
		}
	}
	return -1
}

func BytesMatch(t, p []byte) int {
	n := len(t)
	m := len(p)
	if m == 0 {
		return 0
	}
	for i := 0; i <= n-m; i++ {
		j := 0
		for ; j < m; j++ {
			if t[i+j] != p[j] {
				break
			}
		}
		if j == m {
			return i
		}
	}
	return -1
}

func ClosestPoints(points []*common.Point) float64 {
	n := len(points)
	min := math.MaxFloat64
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			min = math.Min(min, (points[i].X-points[j].X)*(points[i].X-points[j].X)+(points[i].Y-points[j].Y)*(points[i].Y-points[j].Y))
		}
	}
	return math.Sqrt(min)
}

func MinAvgDistance(x []float64) (float64, int) {
	n := len(x)
	min := math.MaxFloat64
	index := -1
	for i := 0; i < n; i++ {
		d := 0.0
		for j := 0; j < n; j++ {
			d += math.Abs(x[j] - x[i])
		}
		if min > d {
			min = d
			index = i
		}
	}
	return min, index
}

func ConvexHull(points []*common.Point) []*common.Point {
	result := make(map[int]bool)
	n := len(points)
	if n <= 3 {
		return points
	}
	for i := 0; i < n - 1 ; i++ {
		for j := i + 1; j < n; j++ {
			s1 := false
			s2 := false
			a, b, c := line(points[i], points[j])

			for k := 0; k < n; k ++ {
				if k == i || k == j {
					continue
				}
				ck := a * points[k].X + b * points[k].Y
				if ck > c {
					s1 = true
				} else if ck < c {
					s2 = true
				}
				if s1 && s2 {
					break
				}
			}
			if ! (s1 && s2) {
				result[i] = true
				result[j] = true
			}
		}
	}
	points2 := make([]*common.Point, 0, len(result))
	for p, _ := range result {
		points2 = append(points2, points[p])
	}
	return points2
}

func line(p1, p2 *common.Point) (a, b, c float64) {
	a = p2.Y - p1.Y
	b = p1.X - p2.X
	c = p1.X * p2.Y - p2.X * p1.Y
	return
}

func Tsp() float64 {
	return 0.0
}
