type Matrix struct {
	self [2][2]int
}

func (matrix Matrix)Mul(other Matrix) *Matrix {
	r := &Matrix{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				r.self[i][j] += matrix.self[i][k] * other.self[k][j];
			}
		}
	}
	return r;
}

func (matrix Matrix)Pow(n int) *Matrix {
	r := &Matrix{self: [2][2]int{{1, 0}, {0, 1}}}
	base := matrix
	for n > 0 {
		if n % 2 == 1 {
			r = r.Mul(base)
		}
		base = *base.Mul(base)
		n >>= 1
	}
	return r
}

func fib(n int) int {
	if (n <= 0) {
		return 0
	}
	if (n <= 1) {
		return 1
	}

	fb := Matrix{
        self: [2][2]int{
            {1, 1,},
            {1, 0,},
        },
	}

	r := fb.Pow(n - 1)
	return r.self[0][0]
}
