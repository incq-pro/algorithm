/*
Package bruteforce algorithm design
*/
package bruteforce

// Polynomial3 多项式求和
func Polynomial3(x float64, a []float64) float64 {
	n := len(a)
	var result float64
	xn := make([]float64, n, n)
	xn[0] = 1
	xn[1] = x
	for i := 2; i < n; i++ {
		xn[i] = xn[i-1] * x
	}
	for i := 0; i < n; i++ {
		result += a[i] * xn[i]
	}
	return result
}

// Polynomial 多项式求和
func Polynomial(x float64, a []float64) float64 {
	var result float64
	n := len(a)
	for i := n - 1; i >= 0; i-- {
		result = result*x + a[i]
	}
	return result
}
