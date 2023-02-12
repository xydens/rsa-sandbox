package lib

import (
	"fmt"
)

func ReverseMod(a int, mod int) (int, bool) {
	var x, y int
	gcd := ExtendedEuclidGCD(a, mod, &x, &y)
	if gcd != 1 {
		return 0, false
	}
	return x, false
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

func ExtendedEuclidGCD(a int, b int, x *int, y *int) int {
	if a == 0 {
		*x = 0
		*y = 1
		return b
	}
	var x1, y1 int
	d := ExtendedEuclidGCD(b%a, a, &x1, &y1)
	*x = y1 - (b/a)*x1
	*y = x1
	fmt.Printf("a: %d; b: %d; x: %d; y: %d x1: %d; y1: %d \n", a, b, *x, *y, x1, y1)
	return d
}
