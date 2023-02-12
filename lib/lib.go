package lib

import (
	"math/big"
)

func ReverseMod(a, mod *big.Int) (*big.Int, bool) {
	x := new(big.Int)
	y := new(big.Int)
	gcd := ExtendedEuclidGCD(a, mod, x, y)
	if gcd.Cmp(new(big.Int).SetInt64(1)) != 0 {
		return new(big.Int), false
	}
	return x, true
}

func EuclidGCD(x int, y int) int {
	if x == 0 {
		return y
	}

	var i int
	for y != 0 {
		i = y
		y = x % y
		x = i
	}
	return x
}

func ExtendedEuclidGCD(a, b *big.Int, x, y *big.Int) *big.Int {
	if a.Cmp(new(big.Int).SetInt64(0)) == 0 {
		x.SetInt64(0)
		y.SetInt64(1)
		return b
	}
	x1 := new(big.Int)
	y1 := new(big.Int)
	d := ExtendedEuclidGCD(new(big.Int).Mod(b, a), a, x1, y1)
	//fmt.Printf("a: %d; b: %d; x1: %d; y1: %d \n", a, b, x1, y1)
	x.Sub(
		y1,
		new(big.Int).Mul(
			new(big.Int).Div(b, a),
			x1,
		),
	) // y1 - (b/a)*x1
	y.Set(x1)
	//fmt.Printf("a: %d; b: %d; x: %d; y: %d x1: %d; y1: %d \n", a, b, x, y, x1, y1)
	return d
}
