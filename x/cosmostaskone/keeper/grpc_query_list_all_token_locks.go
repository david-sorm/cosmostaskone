package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dsorm/cosmostaskone/x/cosmostaskone/types"
	"strings"
)

func (k Keeper) ListAllTokenLocks(goCtx context.Context, req *types.QueryListAllTokenLocksRequest) (*types.QueryListAllTokenLocksResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// create new store from sdk context
	// store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WithPrefix(""))

	tokensLockList := make([]*types.TokensLock, 0, 16)

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.WithPrefix(""))

	var tl types.TokenLockInternal
	bz := []byte("")
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		tl = types.TokenLockInternal{}
		// skip stored last hashes
		if len(strings.Split(string(iterator.Key()), "-")) == 2 {
			continue
		}
		bz = iterator.Value()
		k.cdc.MustUnmarshalBinaryBare(bz, &tl)
		if tl.Disabled {
			continue
		}
		tokensLockList = append(tokensLockList, &types.TokensLock{
			Id:       tl.ID,
			Creator:  tl.Creator,
			Balances: tl.Balances,
		})
	}

	return &types.QueryListAllTokenLocksResponse{TokensLockList: tokensLockList}, nil
}
