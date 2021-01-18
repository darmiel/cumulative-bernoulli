package main

import "math/big"

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
