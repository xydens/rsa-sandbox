package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseMod(t *testing.T) {
	got, _ := ReverseMod(5, 7)
	assert.Equal(t, got, 3)
}
func TestEuclidGCD(t *testing.T) {
	got := EuclidGCD(5, 10)
	if got != 5 {
		t.Errorf("EuclidGCD(5, 10) = %d; want 5", got)
	}
}

func TestExtendedEuclidGCD(t *testing.T) {
	var x, y int
	gcd := ExtendedEuclidGCD(25, 15, &x, &y)
	if gcd != 5 {
		t.Errorf("EuclidGCD(25, 15) = %d; want 5", gcd)
	}
	if x != -1 {
		t.Errorf("x = %d; want -1", x)
	}
	if y != 2 {
		t.Errorf("y = %d; want 2", y)
	}
}
