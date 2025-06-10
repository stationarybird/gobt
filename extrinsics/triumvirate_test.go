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

        // Create new members list with Alice, Bob, and Charlie
        newMembers := []types.AccountID{
            *env.Alice.Hotkey.AccID,
            *env.Bob.Hotkey.AccID,
            *env.Charlie.Hotkey.AccID,
        }

        // Set Alice as prime member
        prime := env.Alice.Hotkey.AccID

        // Get current member count (0 for initial setup)
        oldCount := uint32(0)

        // Create the set members call
        call, err := SetMembersCall(env.Client, newMembers, prime, oldCount)
        require.NoError(t, err, "Failed to create set members call")

        // Create sudo extrinsic to execute the call
        sudoExt, err := NewSudoExt(env.Client, &call)
        require.NoError(t, err, "Failed to create sudo extrinsic")

        // Sign and submit the transaction
        hash, err := testutils.SignAndSubmit(t, env.Client, sudoExt, env.Alice.Coldkey.Keypair, uint32(env.Alice.Coldkey.AccInfo.Nonce))
        require.NoError(t, err, "Failed to submit transaction")
        require.NotNil(t, hash, "Transaction hash should not be nil")

        // Update account info after transaction
        updateUserInfo(t, &env.Alice, env, false)

        // Verify nonce was incremented
        require.Greater(t, env.Alice.Coldkey.AccInfo.Nonce, uint32(0), "Alice's nonce should be incremented after successful transaction")
    })
}
