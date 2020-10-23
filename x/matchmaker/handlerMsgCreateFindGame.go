package matchmaker

import (
	"github.com/cmwaters/rook/x/matchmaker/keeper"
	"github.com/cmwaters/rook/x/matchmaker/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleMsgCreateFindGame(ctx sdk.Context, k keeper.Keeper, findGame *types.MsgFindGame) (*sdk.Result, error) {
	k.CreateFindGame(ctx, *findGame)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
