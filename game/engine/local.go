package engine

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cmwaters/rook/bot"
	"github.com/cmwaters/rook/x/rook/types"
)

// LocalGameEngine generates a map and runs the game locally with a set of bots
// It keeps full state of the game and updates the players state at a defined tick
// speed. It emulates the same interface as a network game engine
type LocalGameEngine struct {
	tickSpeed  time.Duration // the rate with which the board updates
	bots       []bot.Bot
	state      types.GameState
	buildQueue map[int][]*types.MsgBuild
	moveQueue  map[int][]*types.MsgMove
	step       uint32
}

func NewLocalGameEngine(config *types.GameConfig, bots int) *LocalGameEngine {
	state := types.GameState{Config: config, Factions: make([]*types.Faction, bots+1)}
	state.Map = types.GenerateMap(config.Map)
	for i := 0; i < bots+1; i++ {
		state.Factions[i] = types.NewFaction("player"+strconv.Itoa(i), *config.Initial)
	}
	types.PopulateFactions(&state.Map, state.Factions, state.Config.Map)

	return &LocalGameEngine{
		tickSpeed:  3, // every 3 seconds
		bots:       InitializeBots(bots),
		state:      state,
		buildQueue: make(map[int][]*types.MsgBuild),
		moveQueue:  make(map[int][]*types.MsgMove),
		step:       0,
	}
}

func (l *LocalGameEngine) Init(state chan *types.PartialState) {
	for _, b := range l.bots {
		b.Init(*l.state.Config)
	}
	// start the main game loop
	fmt.Println("Starting engine loop")
	ticker := time.NewTicker(l.tickSpeed * time.Second)
	for range ticker.C {
		fmt.Println("tick")
		// first update the state based on all the received messages in that step
		victor := l.processMessages()
		if victor != nil { //termination condition
			// we have a victor. Close the channel and end the game
			fmt.Printf("%s has won", victor.Moniker)
			close(state)
			ticker.Stop()
			return
		}
		// display state changes to all bots and add responses to the queue
		for idx, b := range l.bots {
			tiles := types.GetVisibleTilesFromMap(l.state.Map, *l.state.Factions[idx])
			resp := b.Update(types.PartialState{Map: tiles, Resources: l.state.Factions[idx].Resources, Step: l.step})
			l.buildQueue[idx] = resp.Builds
			l.moveQueue[idx] = resp.Moves
		}
		// get the visible area for the player and send it through the state channel
		// TODO: in the future we could look at just sending the delta than the entire visible map
		playerTiles := types.GetVisibleTilesFromMap(l.state.Map, *l.state.Factions[len(l.bots)])
		state <- &types.PartialState{
			Map:       playerTiles,
			Resources: l.state.Factions[len(l.bots)].Resources,
			Step:      l.step,
		}
		// lastly update the step

		l.step++
	}
}

func (l *LocalGameEngine) Build(settlement types.Settlement, x, y int) error {
	msg := &types.MsgBuild{Settlement: settlement, Position: &types.Position{X: uint32(x), Y: uint32(y)}}
	l.buildQueue[len(l.bots)] = append(l.buildQueue[len(l.bots)], msg)
	return nil
}

func (l *LocalGameEngine) Move(quantity, x, y int, direction types.Direction) error {
	msg := &types.MsgMove{
		Quantity:  uint32(quantity),
		Position:  &types.Position{X: uint32(x), Y: uint32(y)},
		Direction: direction,
	}
	l.moveQueue[len(l.bots)] = append(l.moveQueue[len(l.bots)], msg)
	return nil
}

// process messages updates state for each message and flushes message queue
func (l *LocalGameEngine) processMessages() *types.Faction {
	for idx := 0; idx < len(l.bots)+1; idx++ {
		for _, buildMessage := range l.buildQueue[idx] {
			_ = l.state.Build(l.state.Factions[idx], buildMessage.Settlement, buildMessage.Position)
		}
		l.buildQueue[idx] = []*types.MsgBuild{}
		for _, moveMessage := range l.moveQueue[idx] {
			_ = l.state.Move(l.state.Factions[idx], moveMessage.Quantity, moveMessage.Position, moveMessage.Direction)
		}
		l.moveQueue[idx] = []*types.MsgMove{}
	}
	return l.state.Victor()
}

func InitializeBots(count int) []bot.Bot {
	bots := make([]bot.Bot, count)
	for i := 0; i < count; i++ {
		bots[i] = bot.NewComputer()
	}
	return bots
}
