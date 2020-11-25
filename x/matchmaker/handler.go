package matchmaker

import (
	"fmt"

	"github.com/cmwaters/rook/x/matchmaker/keeper"
	"github.com/cmwaters/rook/x/matchmaker/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		var err error

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MsgBuild:
			err = k.HandleBuildMessage(ctx, *msg)
		case *types.MsgMove:
			err = k.HandleMoveMessage(ctx, *msg)
		case *types.MsgInitialize:
			err = k.InitGame(ctx, *msg)
		case *types.MsgReady:
			err = k.AddGameVote(ctx, *msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
		if err != nil {
			return nil, err
		}
		return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
	}
}
