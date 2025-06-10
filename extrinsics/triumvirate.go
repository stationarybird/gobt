package extrinsics

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/extrinsic"
	"github.com/subtrahend-labs/gobt/client"
)

func SetMembersCall(c *client.Client, newMembers []types.AccountID, prime *types.AccountID, oldCount uint32) (types.Call, error) {
	call, err := types.NewCall(
		c.Meta,
		"Triumvirate.set_members",
		newMembers,
		prime,
		oldCount,
	)
	if err != nil {
		return types.Call{}, err
	}
	return call, err
}

func SetMembersExt(c *client.Client, newMembers []types.AccountID, prime *types.AccountID, oldCount uint32) (*extrinsic.Extrinsic, error) {
	call, err := SetMembersCall(c, newMembers, prime, oldCount)
	if err != nil {
		return nil, err
	}
	ext := extrinsic.NewExtrinsic(call)
	return &ext, nil
}

// func ExecuteCall(c *client.Client, proposal types.Call, lengthBound uint32) (types.Call, error) {
// 	call, err := types.NewCall(
// 		c.Meta,
// 		"Triumvirate.execute",
// 		proposal,
// 		lengthBound,
// 	)
// 	if err != nil {
// 		return types.Call{}, err
// 	}
// 	return call, nil
// }

// func ExecuteExt(c *client.Client, proposal types.Call, lengthBound uint32) (*extrinsic.Extrinsic, error) {
// 	call, err := ExecuteCall(c, proposal, lengthBound)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ext := extrinsic.NewExtrinsic(call)
// 	return &ext, nil
// }
