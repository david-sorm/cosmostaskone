package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dsorm/cosmostaskone/x/cosmostaskone/types"
)

func (k msgServer) AddTokensLock(goCtx context.Context, msg *types.MsgAddTokensLock) (*types.MsgAddTokensLockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddTokensLockResponse{}, nil
}
