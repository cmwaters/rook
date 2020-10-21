package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cmwaters/rook/x/rook/types"
	"github.com/cmwaters/rook/x/rook/proto"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) CreateBuild(ctx sdk.Context, build proto.MsgBuild) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuildKey))
	b := k.cdc.MustMarshalBinaryBare(&build)
	store.Set(types.KeyPrefix(types.BuildKey), b)
}

func (k Keeper) GetAllBuild(ctx sdk.Context) (msgs []proto.MsgBuild) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuildKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.BuildKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg proto.MsgBuild
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
