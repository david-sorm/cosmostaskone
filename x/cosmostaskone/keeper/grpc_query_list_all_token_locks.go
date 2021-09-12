package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dsorm/cosmostaskone/x/cosmostaskone/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListAllTokenLocks(goCtx context.Context, req *types.QueryListAllTokenLocksRequest) (*types.QueryListAllTokenLocksResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// create new store from sdk context
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WithPrefix(""))

	tokensLockList := make([]*types.TokensLock, 0, 16)

	for tokenLock := types.TokenLockStartNode(store, k.cdc); len(tokenLock.NextNode) != 0; tokenLock.Next(store, k.cdc) {
		tokensLockList = append(tokensLockList, &types.TokensLock{
			Id:       tokenLock.ID,
			Creator:  tokenLock.Creator,
			Balances: tokenLock.Balances,
		})
	}

	return &types.QueryListAllTokenLocksResponse{TokensLockList: tokensLockList}, nil
}
