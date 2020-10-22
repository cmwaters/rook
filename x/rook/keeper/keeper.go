package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cmwaters/rook/x/rook/types"
)

type Keeper struct {
	cdc      codec.Marshaler
	storeKey sdk.StoreKey
	memKey   sdk.StoreKey

	games map[string]*types.GameState
	pendingGames map[string]*types.PendingGameState
	players map[string]string
}

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) AddGameVote(ctx sdk.Context, msg types.MsgReady) error {
	game, ok := k.pendingGames[msg.Id]
	if !ok {
		return fmt.Errorf("game with id %s does not exist", msg.Id)
	}
	ready := game.AddVote(&msg.Creator)
	if ready {
		k.games[msg.Id] = types.NewGameState(game.Participants(), game.Config()) 
		delete(k.pendingGames, msg.Id)
	}
	return nil
}

func (k Keeper) InitGame(ctx sdk.Context, msg types.MsgInitialize) error {
	if _, ok := k.pendingGames[msg.Id]; ok {
		return fmt.Errorf("game with id %s already exists", msg.Id)
	}
	k.pendingGames[msg.Id] = types.NewPendingGame(msg.Players, msg.Id, msg.Config)
	return nil
}