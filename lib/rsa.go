package lib

import "math/big"

type Primes struct {
	P int64
	Q int64
}

func (receiver Primes) GetN() *big.Int {
	p := new(big.Int)
	p.SetInt64(receiver.P)

	q := new(big.Int)
	q.SetInt64(receiver.Q)

	z := new(big.Int)
	return z.Mul(p, q)
}

func (receiver Primes) Euler() *big.Int {
	p := new(big.Int)
	p.SetInt64(receiver.P - 1)

	q := new(big.Int)
	q.SetInt64(receiver.Q - 1)

	z := new(big.Int)
	return z.Mul(p, q)
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

func Decrypt(encryptedMsg *big.Int, privateKey PrivateKey) (*big.Int, bool) {
	euler := privateKey.N.Euler()
	exp := new(big.Int).Sub(euler, new(big.Int).SetInt64(2))
	d := new(big.Int).Exp(privateKey.E, exp, euler)

	return new(big.Int).Exp(encryptedMsg, d, privateKey.N.GetN()), true
}
