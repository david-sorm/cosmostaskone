package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dsorm/cosmostaskone/x/cosmostaskone/types"
)

const keyPrefix = "TokenLock-"

func (k msgServer) SetTokensLockCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(keyPrefix))
	byteKey := types.KeyPrefix("Count")
	countString := strconv.FormatUint(count, 10)

	store.Set(byteKey, []byte(countString))
}

func (k msgServer) AddTokensLock(goCtx context.Context, msg *types.MsgAddTokensLock) (*types.MsgAddTokensLockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// create new store from sdk context
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(keyPrefix))

	// generate the ID, (last id) + 1
	count := k.GetTokensLockCount(ctx)
	id := count + uint64(1)

	// create unique key from prefix and id
	key := keyPrefix + strconv.FormatUint(id, 10)

	// convert msg to binary
	value := k.cdc.MustMarshalBinaryBare(msg)
	store.Set([]byte(key), value)

	k.SetTokensLockCount(ctx, count+1)

	idString := strconv.FormatUint(id, 10)
	return &types.MsgAddTokensLockResponse{Id: idString}, nil
}
