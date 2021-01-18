package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
)

// nCk * p^k * (1 - p)^(n-k)
func B(n, p, k float64) *big.Float {
	// cr := nCr(big.NewFloat(n), big.NewFloat(k))
	cr := nCrEfficient(int64(n), int64(k))
	pPk := partialExp(big.NewFloat(p), big.NewInt(int64(k)))
	gPnk := partialExp(new(big.Float).Sub(big.NewFloat(1.0), big.NewFloat(p)), big.NewInt(int64(n-k)))
	return new(big.Float).Mul(new(big.Float).Mul(cr, pPk), gPnk)
}

// B(n, p, 1) + B(n, p, 2) + ... + B(n, p, n)
func F(n int64, p float64, k int64) (res *big.Float) {
	res = big.NewFloat(0.0)
	var i int64 = 1
	for ; i <= k; i++ {
		res = res.Add(res, B(float64(n), p, float64(i)))
	}
	return
}

func println(w io.Writer, v ...interface{}) {
	if w == nil {
		log.Println(v...)
	} else {
		_, _ = fmt.Fprintln(w, v...)
	}
}

func findUpperBound(w io.Writer, n int64, p float64, P float64) (*int64, *float64) {
	log.Println("+ Searching for P( X <=", P, ") [", n, "; ", p, "]")

	var boundL int64 = 0
	var boundR = n

	var i int64
	for i = 0; i <= n; i++ {
		var s = float64(boundL+boundR) / 2.0
		var half = int64(math.Floor(s))
		var val = F(n, p, half)
		f, _ := val.Float64()

		println(w, " ┌ #", i, "| s:", s, "half:", half, "val:", val, "f:", f)

		// check if value was found
		if f == P {
			return &half, &f
		}

		var dir Direction

		// update bounds
		if f < P {
			boundL = half
			dir = DirLeft

			println(w, " ├ Updated left bound to:", boundL)
		} else {
			boundR = half
			dir = DirRight
			println(w, " ├ Updated right bound to:", boundR)
		}

		drawLine(w, n, boundL, boundR, dir)

		// check if bounds too narrow
		if (boundR - boundL) <= 1 {
			i2 := int64(math.Ceil(s)) - 1
			println(w, " └ Bounds too narrow. Using:", i2)
			return &i2, &f
		}

		println(w, " └ Waiting for next ...")
	}

	return nil, nil
}

// lower-equals
func findUpperBoundLe(w io.Writer, n int64, p float64, P float64) *int64 {
	bound, val := findUpperBound(w, n, p, P)
	if bound == nil || val == nil {
		return nil
	}
	if *val > P {
		i := *bound - 1
		return &i
	} else {
		return bound
	}
}

// lower-equals
func findUpperBoundGe(w io.Writer, n int64, p float64, P float64) *int64 {
	bound, val := findUpperBound(w, n, p, P)
	if bound == nil || val == nil {
		return nil
	}
	if *val > P {
		i := *bound
		return &i
	} else {
		i := *bound + 1
		return &i
	}
}
