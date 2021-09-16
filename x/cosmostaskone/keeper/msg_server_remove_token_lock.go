package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dsorm/cosmostaskone/x/cosmostaskone/types"
)

func (k msgServer) RemoveTokenLock(goCtx context.Context, msg *types.MsgRemoveTokenLock) (*types.MsgRemoveTokenLockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// create new store from sdk context
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WithPrefix(""))

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
	address, err := sdk.AccAddressFromBech32(tokenLock.Creator)
	if err != nil {
		return nil, err
	}
	coins := types.DereferenceCoinSlice(tokenLock.Balances)

	// TODO use Module Accounts
	err = k.bankKeeper.AddCoins(ctx, address, coins)
	if err != nil {
		return nil, err
	}

	return &types.MsgRemoveTokenLockResponse{}, nil
}
