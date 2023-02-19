package lib

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestProducingCommonSecret(t *testing.T) {
	commons := &Commons{
		M: big.NewInt(17),
		G: big.NewInt(3),
	}

	alicePrivate := big.NewInt(15)
	alicePeerNumber := getPeerNumber(commons, alicePrivate)

	bobPrivate := big.NewInt(13)
	bobPeerNumber := getPeerNumber(commons, bobPrivate)

	aliceSecret := getCommonSecret(commons, alicePrivate, bobPeerNumber)
	bobSecret := getCommonSecret(commons, bobPrivate, alicePeerNumber)

	assert.Equal(t, aliceSecret, bobSecret, "Expected Alice secret %d to equal Bob secret %d", aliceSecret, bobSecret)
}
