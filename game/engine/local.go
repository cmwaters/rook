package engine

import (
	"github.com/cmwaters/rook/bot"
	"github.com/cmwaters/rook/x/rook/types"
)

// LocalGameEngine generates a map and runs the game locally with a set of bots
// It keeps full state of the game and updates the players state at a defined tick
// speed. It emulates the same interface as a network game engine
type LocalGameEngine struct {
	tickSpeed uint32 // the rate with which the board updates
	mapSeed int64
	bots []bot.Bot
	gameMap [][]types.Tile
}

func NewLocalGameEngine(config *types.GameConfig, bots int) *LocalGameEngine {
	return &LocalGameEngine{
		tickSpeed: 3, // every 3 seconds
		mapSeed: config.Map.Seed,
		bots: InitializeBots(bots),
		gameMap: GenerateMap(config.Map),
	}
}

func (l *LocalGameEngine) Init(state chan *types.PartialState) {

}

func (l *LocalGameEngine) Build(settlement types.Settlement, x, y int) error {
	return nil
}

func (l *LocalGameEngine) Move(quantity, x, y int, direction types.Direction) error {
	return nil
}





func GenerateMap(config *types.MapConfig) [][]types.Tile {
	
}

func InitializeBots(count int) []bot.Bot {
	bots := make([]bot.Bot, count)
	for i := 0; i < count; i++ {
		bots[i] = bot.NewComputer()
	}
	return bots
}