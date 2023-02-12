package lib

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestReverseMod(t *testing.T) {
	got, _ := ReverseMod(new(big.Int).SetInt64(5), new(big.Int).SetInt64(7))
	assert.Equal(t, int64(3), got.Int64())
}
func TestEuclidGCD(t *testing.T) {
	got := EuclidGCD(5, 10)
	if got != 5 {
		t.Errorf("EuclidGCD(5, 10) = %d; want 5", got)
	}
}

func TestExtendedEuclidGCD(t *testing.T) {
	assert := assert.New(t)
	x := new(big.Int)
	y := new(big.Int)
	gcd := ExtendedEuclidGCD(new(big.Int).SetInt64(25), new(big.Int).SetInt64(15), x, y)
	assert.Equal(int64(5), gcd.Int64())
	assert.Equal(int64(-1), x.Int64())
	assert.Equal(int64(2), y.Int64())
}
