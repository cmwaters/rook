package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cmwaters/rook/x/rook/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) HandleBuildMessage(ctx sdk.Context, build types.MsgBuild) error {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuildKey))
	b := k.cdc.MustMarshalBinaryBare(&build)
	store.Set(types.KeyPrefix(types.BuildKey), b)
	return nil
}
