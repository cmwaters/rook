package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cmwaters/rook/x/matchmaker/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) CreateFindGame(ctx sdk.Context, findGame types.MsgFindGame) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FindGameKey))
	b := k.cdc.MustMarshalBinaryBare(&findGame)
	store.Set(types.KeyPrefix(types.FindGameKey), b)
}

func (k Keeper) GetAllFindGame(ctx sdk.Context) (msgs []types.MsgFindGame) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FindGameKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.FindGameKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgFindGame
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
