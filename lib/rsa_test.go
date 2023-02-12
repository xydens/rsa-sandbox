package lib

import (
	"math/big"
	"testing"
)

func TestEncodeAndDecodeInt(t *testing.T) {
	primes := Primes{
		P: 6,
		Q: 11,
	}
	privateKey := PrivateKey{
		N: primes,
		E: new(big.Int).SetInt64(47),
	}

	publicKey := privateKey.GetPublicKey()
	t.Logf("publicKey: %+v", publicKey)
	message := new(big.Int).SetInt64(65)
	encoded := Encrypt(message, publicKey)
	t.Logf("enocoded: %d", encoded)
	//if encoded.Cmp(message) == 0 {
	//	t.Errorf("expected encoded not to be eqaul to the message")
	//}
	decoded := Decrypt(encoded, privateKey)
	if decoded == nil {
		t.Errorf("Decrypt(encoded, privateKey) expecte to get result")
		return
	}
	t.Logf("decoded: %d, message: %d", decoded, message)
	if decoded.Cmp(message) != 0 {
		t.Errorf("Decrypt(encoded, privateKey) = %d; want %d", decoded, message)
	}
}
