package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dsorm/cosmostaskone/x/cosmostaskone/types"
	"strings"
)

func (k msgServer) RemoveTokenLock(goCtx context.Context, msg *types.MsgRemoveTokenLock) (*types.MsgRemoveTokenLockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// create new store from sdk context
	store := ctx.KVStore(k.storeKey)

	// make sure the ID doesn't have the prefix
	msg.Id = strings.TrimPrefix(msg.Id, string(types.WithPrefix("")))

	tokenLock, exists := types.TokenLockLoadIfExists(store, k.cdc, msg.Id)

	// verify token lock's existence
	if !exists {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "The token lock doesn't exist")
	}

	// verify that the sender address is correct
	if !(tokenLock.Creator == msg.Creator) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "The sender's address is invalid")
	}

	// make sure the tokenlock isn't disabled already
	if tokenLock.Disabled {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "The token lock no longer exists")
	}

	tokenLock.Disabled = true
	tokenLock.Save(store, k.cdc)
	creatorAddress, err := sdk.AccAddressFromBech32(tokenLock.Creator)
	if err != nil {
		return nil, err
	}
	coins := types.DereferenceCoinSlice(tokenLock.Balances)

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAddress, coins)
	if err != nil {
		return nil, err
	}

	return &types.MsgRemoveTokenLockResponse{}, nil
}
