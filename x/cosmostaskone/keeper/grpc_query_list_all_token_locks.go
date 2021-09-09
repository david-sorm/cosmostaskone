package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dsorm/cosmostaskone/x/cosmostaskone/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetTokensLockCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(keyPrefix))
	byteKey := types.KeyPrefix("Count")
	result := store.Get(byteKey)

	// no result, no element
	if result == nil {
		return 0
	}

	// bytes -> uint64
	count, err := strconv.ParseUint(string(result), 10, 64)
	if err != nil {
		panic("err while converting bytes to uint64")
	}

	return count
}

func (k Keeper) ListAllTokenLocks(goCtx context.Context, req *types.QueryListAllTokenLocksRequest) (*types.QueryListAllTokenLocksResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// create new store from sdk context
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(keyPrefix))

	tokenLockCount := k.GetTokensLockCount(ctx)

	tokensLockList := make([]*types.TokensLock, 0, tokenLockCount)

	var msg types.MsgAddTokensLock
	for i := uint64(1); i <= tokenLockCount; i++ {
		key := keyPrefix + strconv.FormatUint(i, 10)
		bz := store.Get([]byte(key))

		msg = types.MsgAddTokensLock{}
		k.cdc.MustUnmarshalBinaryBare(bz, &msg)
		tokensLockList = append(tokensLockList, &types.TokensLock{
			Id:       strconv.FormatUint(i, 10),
			Creator:  msg.Creator,
			Balances: msg.Balances,
		})
	}

	return &types.QueryListAllTokenLocksResponse{TokensLockList: tokensLockList}, nil
}
