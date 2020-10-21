package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cmwaters/rook/x/rook/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) CreateMove(ctx sdk.Context, move types.MsgMove) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MoveKey))
	b := k.cdc.MustMarshalBinaryBare(&move)
	store.Set(types.KeyPrefix(types.MoveKey), b)
}

func (k Keeper) GetAllMove(ctx sdk.Context) (msgs []types.MsgMove) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MoveKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.MoveKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgMove
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
