package lib

import (
	"math/big"
)

type Primes struct {
	P *big.Int
	Q *big.Int
}

func (receiver Primes) GetN() *big.Int {
	return new(big.Int).Mul(receiver.P, receiver.Q)
}

func (receiver Primes) Euler() *big.Int {
	p := new(big.Int).Sub(receiver.P, big.NewInt(1))
	q := new(big.Int).Sub(receiver.Q, big.NewInt(1))
	return new(big.Int).Mul(p, q)
}

type PrivateKey struct {
	E *big.Int // >= 3; <= E(pq)
	N Primes
}

func (k PrivateKey) GetPublicKey() PublicKey {
	return PublicKey{
		E: k.E,
		N: k.N.GetN(),
	}
}

type PublicKey struct {
	E *big.Int
	N *big.Int
}

func Encrypt(message *big.Int, publicKey PublicKey) *big.Int {
	return new(big.Int).Exp(message, publicKey.E, publicKey.N)
}

func Decrypt(encryptedMsg *big.Int, privateKey PrivateKey) *big.Int {
	d := new(big.Int).ModInverse(privateKey.E, privateKey.N.Euler())
	return new(big.Int).Exp(encryptedMsg, d, privateKey.N.GetN())
}
