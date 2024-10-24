/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package suite

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/trustbloc/vc-go/internal/mock/kmscrypto"
	"github.com/trustbloc/vc-go/signature/api"
)

func TestNewCryptoSigner(t *testing.T) {
	cryptoSigner := NewCryptoWrapperSigner(&kmscrypto.MockFixedKeyCrypto{
		SignVal: []byte("signature"),
	})
	require.NotNil(t, cryptoSigner)

	signature, err := cryptoSigner.Sign([]byte("msg"))
	require.NoError(t, err)
	require.Equal(t, []byte("signature"), signature)
}

func TestNewCryptoVerifier(t *testing.T) {
	kc := &kmscrypto.MockKMSCrypto{
		VerifyErr: errors.New("verify error"),
	}

	cryptoVerifier := NewCryptoVerifier(kc)
	require.NotNil(t, cryptoVerifier)

	err := cryptoVerifier.Verify(&api.PublicKey{}, []byte("msg"), []byte("signature"))
	require.Error(t, err)
	require.EqualError(t, err, "verify error")
}
