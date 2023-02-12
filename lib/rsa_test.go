package lib

import (
	"math/big"
	"testing"
)

func TestEncodeAndDecodeInt(t *testing.T) {
	t.Skip()
	primes := Primes{
		P: 6,
		Q: 11,
	}
	privateKey := PrivateKey{
		N: primes,
		E: new(big.Int).SetInt64(7),
	}

	publicKey := privateKey.GetPublicKey()
	t.Logf("publicKey: %+v", publicKey)
	message := new(big.Int).SetInt64(20)
	encoded := Encrypt(message, publicKey)
	t.Logf("enocoded: %d", encoded)
	if encoded.Cmp(message) == 0 {
		t.Errorf("expected encoded not to be eqaul to the message")
	}
	decoded, decodeSuccess := Decrypt(encoded, privateKey)
	if !decodeSuccess {
		t.Errorf("Decrypt(encoded, privateKey) expecte to get result")
	}
	if decoded.Cmp(message) != 0 {
		t.Errorf("Decrypt(encoded, privateKey) = %d; want %d", decoded, message)
	}
}
