//go:build integration
// +build integration

package extrinsics

import (
	"testing"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/stretchr/testify/require"
	"github.com/subtrahend-labs/gobt/testutils"
)

func TestTriumvirateExtrinsics(t *testing.T) {
    t.Parallel()
    t.Run("SetMembers", func(t *testing.T) {
        t.Parallel()
        env := setup(t)
        setupSubnet(t, env)

		newMembers := []types.AccountID{
            *env.Alice.Hotkey.AccID,
            *env.Bob.Hotkey.AccID,
            *env.Charlie.Hotkey.AccID,
        }

        prime := env.Alice.Hotkey.AccID

        oldCount := uint32(0)

        call, err := SetMembersCall(env.Client, newMembers, prime, oldCount)
        require.NoError(t, err, "Failed to create set members call")

        sudoExt, err := NewSudoExt(env.Client, &call)
        require.NoError(t, err, "Failed to create sudo extrinsic")

        hash, err := testutils.SignAndSubmit(t, env.Client, sudoExt, env.Alice.Coldkey.Keypair, uint32(env.Alice.Coldkey.AccInfo.Nonce))
        require.NoError(t, err, "Failed to submit transaction")
        require.NotNil(t, hash, "Transaction hash should not be nil")

        updateUserInfo(t, &env.Alice, env, false)

        require.Greater(t, env.Alice.Coldkey.AccInfo.Nonce, uint32(0), "Alice's nonce should be incremented after successful transaction")
    })
}
