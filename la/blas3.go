// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package la

import (
	"github.com/cpmech/gosl/la/oblas"
)

// MatMatMul returns the matrix multiplication (scaled)
//
//  c := α⋅a⋅b    ⇒    cij := α * aik * bkj
//
func MatMatMul(c *Matrix, α float64, a, b *Matrix) {
	if c.M < 6 && c.N < 6 && a.N < 30 {
		for i := 0; i < c.M; i++ {
			for j := 0; j < c.N; j++ {
				c.Set(i, j, 0.0)
				for k := 0; k < a.N; k++ {
					c.Add(i, j, α*a.Get(i, k)*b.Get(k, j))
				}
			}
		}
		return
	}
	oblas.Dgemm(false, false, a.M, b.N, a.N, α, a.Data, a.M, b.Data, b.M, 0.0, c.Data, c.M)
}

// MatTrMatMul returns the matrix multiplication (scaled) with transposed(a)
//
//  c := α⋅aᵀ⋅b    ⇒    cij := α * aki * bkj
//
func MatTrMatMul(c *Matrix, α float64, a, b *Matrix) {
	if c.M < 6 && c.N < 6 && a.M < 30 {
		for i := 0; i < c.M; i++ {
			for j := 0; j < c.N; j++ {
				c.Set(i, j, 0.0)
				for k := 0; k < a.M; k++ {
					c.Add(i, j, α*a.Get(k, i)*b.Get(k, j))
				}
			}
		}
		return
	}
	oblas.Dgemm(true, false, a.N, b.N, a.M, α, a.Data, a.M, b.Data, b.M, 0.0, c.Data, c.M)
}

// MatMatTrMul returns the matrix multiplication (scaled) with transposed(b)
//
//  c := α⋅a⋅bᵀ    ⇒    cij := α * aik * bjk
//
func MatMatTrMul(c *Matrix, α float64, a, b *Matrix) {
	oblas.Dgemm(false, true, a.M, b.M, a.N, α, a.Data, a.M, b.Data, b.M, 0.0, c.Data, c.M)
}

// MatTrMatTrMul returns the matrix multiplication (scaled) with transposed(a) and transposed(b)
//
//  c := α⋅aᵀ⋅bᵀ    ⇒    cij := α * aki * bjk
//
func MatTrMatTrMul(c *Matrix, α float64, a, b *Matrix) {
	oblas.Dgemm(true, true, a.N, b.M, a.M, α, a.Data, a.M, b.Data, b.M, 0.0, c.Data, c.M)
}

// mat mul add ////////////////////////////////////////////////////////////////////////////////////

// MatMatMulAdd returns the matrix multiplication (scaled)
//
//  c += α⋅a⋅b    ⇒    cij += α * aik * bkj
//
func MatMatMulAdd(c *Matrix, α float64, a, b *Matrix) {
	oblas.Dgemm(false, false, a.M, b.N, a.N, α, a.Data, a.M, b.Data, b.M, 1.0, c.Data, c.M)
}

// MatTrMatMulAdd returns the matrix multiplication (scaled) with transposed(a)
//
//  c += α⋅aᵀ⋅b    ⇒    cij += α * aki * bkj
//
func MatTrMatMulAdd(c *Matrix, α float64, a, b *Matrix) {
	oblas.Dgemm(true, false, a.N, b.N, a.M, α, a.Data, a.M, b.Data, b.M, 1.0, c.Data, c.M)
}

// MatMatTrMulAdd returns the matrix multiplication (scaled) with transposed(b)
//
//  c += α⋅a⋅bᵀ    ⇒    cij += α * aik * bjk
//
func MatMatTrMulAdd(c *Matrix, α float64, a, b *Matrix) {
	oblas.Dgemm(false, true, a.M, b.M, a.N, α, a.Data, a.M, b.Data, b.M, 1.0, c.Data, c.M)
}

// MatTrMatTrMulAdd returns the matrix multiplication (scaled) with transposed(a) and transposed(b)
//
//  c += α⋅aᵀ⋅bᵀ    ⇒    cij += α * aki * bjk
//
func MatTrMatTrMulAdd(c *Matrix, α float64, a, b *Matrix) {
	oblas.Dgemm(true, true, a.N, b.M, a.M, α, a.Data, a.M, b.Data, b.M, 1.0, c.Data, c.M)
}
