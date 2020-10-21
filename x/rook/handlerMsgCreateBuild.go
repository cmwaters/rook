package rook

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cmwaters/rook/x/rook/types"
	"github.com/cmwaters/rook/x/rook/keeper"
)

func handleMsgCreateBuild(ctx sdk.Context, k keeper.Keeper, build *types.MsgBuild) (*sdk.Result, error) {
	k.CreateBuild(ctx, *build)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
