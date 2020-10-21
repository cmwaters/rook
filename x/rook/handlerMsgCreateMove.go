package rook

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cmwaters/rook/x/rook/types"
	"github.com/cmwaters/rook/x/rook/keeper"
)

func handleMsgCreateMove(ctx sdk.Context, k keeper.Keeper, move *types.MsgMove) (*sdk.Result, error) {
	k.CreateMove(ctx, *move)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
