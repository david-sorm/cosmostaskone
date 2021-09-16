package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	cosmosTypes "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dsorm/cosmostaskone/x/cosmostaskone/types"
)

func (k msgServer) AddTokensLock(goCtx context.Context, msg *types.MsgAddTokensLock) (*types.MsgAddTokensLockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// create new store from sdk context
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WithPrefix(""))

	// check if the account has the balance specified
	address, err := cosmosTypes.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	for _, v := range msg.Balances {
		if !k.bankKeeper.HasBalance(ctx, address, *v) {
			return nil,
				sdkerrors.Wrapf(
					sdkerrors.ErrInsufficientFunds,
					"The account (%v) doesn't have %v %v\n",
					msg.Creator,
					v.Amount,
					v.Denom)
		}
	}

	// convert []*cosmosTypes.Coin to []cosmosTypes.Coin
	coins := types.DereferenceCoinSlice(msg.Balances)

	// TODO use Module Accounts
	err = k.bankKeeper.SubtractCoins(ctx, address, coins)
	if err != nil {
		return nil, err
	}

	// update the data structures
	lastNode := types.TokenLockStartNode(store, k.cdc).Last(store, k.cdc)

	currentNode := types.TokenLockInternal{
		ID:       "",
		Creator:  msg.Creator,
		Balances: msg.Balances,
		NextNode: "",
	}
	currentNode.GenerateUniqueID(store)
	lastNode.NextNode = currentNode.ID

	currentNode.Save(store, k.cdc)
	lastNode.Save(store, k.cdc)

	return &types.MsgAddTokensLockResponse{Id: currentNode.ID}, nil
}
