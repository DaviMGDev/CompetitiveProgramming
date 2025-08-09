package main

import (
	"fmt"
)

type Matrix struct {
	self [2][2]uint64
}

func (matrix Matrix)Mul(other Matrix, m int) *Matrix {
	r := &Matrix{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				r.self[i][j] += (matrix.self[i][k] * other.self[k][j]) % uint64(m)
			}
		}
	}
	return r
}

func (matrix Matrix)Pow(n uint64, m int) *Matrix {
	r := &Matrix{
		self: [2][2]uint64{
			{1, 0,},
			{0, 1,},
		},
	}
	base := matrix 
	for n = n; n > 0; n >>= 1{
		if n % 2 == 1 {
			r = r.Mul(base, m)
		}
		base = *base.Mul(base, m)
	}
	return r
}

func fib(n uint64, m int) uint64 {
	if n <= 1 {
		return 1
	}
	fb := Matrix{
		self: [2][2]uint64{
			{1, 1,},
			{1, 0,},
		},
	}

	mr := fb.Pow(n, m)
	return 1 + (((2 * mr.self[0][0]) - 2) % uint64(m))
}

func main() {
	var n uint64 = 1
	var b int    = 1
	fmt.Scanln(&n, &b)
	for i := 1; n != 0 || b != 0; i++ {
		fmt.Printf("Case %d: %d %d %d\n", i, n, b, fib(n, b))
		fmt.Scanln(&n, &b)
	}
}
