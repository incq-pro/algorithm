package bruteforce

import "testing"

func TestPolynomial(t *testing.T) {
	type args struct {
		x float64
		a []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"1",
			args{
				2.0,
				[]float64{2.1, 3.1},
			},
			8.3,
		},
		{
			"2",
			args{
				3.0,
				[]float64{1.1, 2.1, 3.1},
			},
			35.3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Polynomial(tt.args.x, tt.args.a); (got - tt.want) > 0.001 {
				t.Errorf("Polynomial(%v, %v) = %v, want %v", tt.args.x, tt.args.a, got, tt.want)
			}
		})
	}
}
