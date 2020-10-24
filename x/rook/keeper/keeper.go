package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cmwaters/rook/x/rook/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	cdc      codec.Marshaler
	storeKey sdk.StoreKey
	memKey   sdk.StoreKey

	games        map[string]*types.GameState
	pendingGames map[string]*types.PendingGameState
	players      map[string]string
}

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		games:        make(map[string]*types.GameState),
		pendingGames: make(map[string]*types.PendingGameState),
		players:      make(map[string]string),
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
		players := game.Participants()
		k.games[msg.Id] = types.NewGameState(players, game.Config())
		for _, player := range players {
			k.players[player.String()] = msg.Id
		}
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

func (k *Keeper) CheckPendingGames() {
	for id, game := range k.pendingGames {
		expired, quorum, players := game.Tick()
		if expired && quorum {
			k.games[id] = types.NewGameState(players, game.Config())
			for _, player := range players {
				k.players[player.String()] = id
			}
			delete(k.pendingGames, id)
		} else if expired {
			// not enough players voted in time, delete pending game
			delete(k.pendingGames, id)
		}
		// game hasn't expired yet so continue to wait for votes
	}
}

func (k Keeper) HandleBuildMessage(ctx sdk.Context, build types.MsgBuild) error {
	if gameId, ok := k.players[build.Creator.String()]; ok {
		game := k.games[gameId]
		return game.Build(game.Players[build.Creator.String()], build.Settlement, build.Position)
	}
	return fmt.Errorf("Player with address %s is not involved in any game currently", build.Creator.String())
}

func (k Keeper) HandleMoveMessage(ctx sdk.Context, move types.MsgMove) error {
	if gameId, ok := k.players[move.Creator.String()]; ok {
		game := k.games[gameId]
		return game.Move(game.Players[move.Creator.String()], move.Quantity, move.Position, move.Direction)
	}
	return fmt.Errorf("Player with address %s is not involved in any game currently", move.Creator.String())
}
