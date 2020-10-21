package matchmaker

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cmwaters/rook/x/matchmaker/types"
	"github.com/cmwaters/rook/x/matchmaker/keeper"
)

func handleMsgCreateFindGame(ctx sdk.Context, k keeper.Keeper, findGame *types.MsgFindGame) (*sdk.Result, error) {
	k.CreateFindGame(ctx, *findGame)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
