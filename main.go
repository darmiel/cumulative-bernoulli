package main

import (
	"log"
	"math/big"
)

func fac(n *big.Float) *big.Float {
	if n.String() == "0" || n.String() == "0" {
		return big.NewFloat(float64(1))
	}
	return new(big.Float).Mul(fac(new(big.Float).Sub(n, big.NewFloat(1))), n)
}

// n! / (r! * (n - r)!)
func nCr(n, r *big.Float) *big.Float {
	return new(big.Float).Quo(fac(n), new(big.Float).Mul(fac(r), fac(new(big.Float).Sub(n, r))))
}

func partialExp(x *big.Float, e *big.Int) (res *big.Float) {
	if e.Int64() == 0 {
		return big.NewFloat(1.0)
	}

	res = x // first iteration
	i := e

	for i.Int64() > 1 { // ignore 1
		i = i.Sub(i, big.NewInt(1))
		res = new(big.Float).Mul(res, x)
	}

	return
}

// nCk * p^k * (1 - p)^(n-k)
func B(n, p, k float64) *big.Float {
	cr := nCr(big.NewFloat(n), big.NewFloat(k))
	pPk := partialExp(big.NewFloat(p), big.NewInt(int64(k)))
	_1sP := new(big.Float).Sub(big.NewFloat(1.0), big.NewFloat(p))
	NsK := big.NewInt(int64(n - k))
	gPnk := partialExp(_1sP, NsK)
	x1 := new(big.Float).Mul(cr, pPk)
	x2 := new(big.Float).Mul(x1, gPnk)
	return x2
}

func main() {
	log.Println("B(100;0.2;20)", B(100, .2, 20))
	log.Println(partialExp(big.NewFloat(0.2), big.NewInt(100)))
}
