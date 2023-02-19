package lib

import "math/big"

type Commons struct {
	G *big.Int
	M *big.Int
}

func getPeerNumber(c *Commons, private *big.Int) *big.Int {
	return new(big.Int).Exp(c.G, private, c.M)
}

func getCommonSecret(c *Commons, private, peerNumber *big.Int) *big.Int {
	return new(big.Int).Exp(peerNumber, private, c.M)
}
