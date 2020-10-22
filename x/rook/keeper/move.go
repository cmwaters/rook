package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cmwaters/rook/x/rook/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) HandleMoveMessage(ctx sdk.Context, move types.MsgMove) error {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MoveKey))
	b := k.cdc.MustMarshalBinaryBare(&move)
	store.Set(types.KeyPrefix(types.MoveKey), b)
	return nil
}