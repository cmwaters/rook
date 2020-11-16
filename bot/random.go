package bot

import (
	"math/rand"
	"github.com/cmwaters/rook/x/rook/types"
)

var _ Bot = &RandomBot{}

const (
	moveProb = 4
	buildProb = 8
)

var settlements = []types.Settlement{
	types.Settlement_CAPITAL,
	types.Settlement_CITY,
	types.Settlement_FARM, 
	types.Settlement_LUMBERMILL, 
	types.Settlement_QUARRY, 
	types.Settlement_ROOK, 
	types.Settlement_TOWN,
}

var directions = []types.Direction{
	types.Direction_LEFT,
	types.Direction_RIGHT,
	types.Direction_DOWN,
	types.Direction_UP,
}

// RandomBot makes random moves and builds. Should be used for testing purposes only.
type RandomBot struct {
	board   [][]types.Tile
	resources *types.ResourceSet
	config  types.GameConfig
	moniker string
	step    uint32
}

func NewRandomBot() *RandomBot {
	return &RandomBot{} // uses all zero values which will be populated upon init
}

// TODO: stub out init and update functions
func (r *RandomBot) Init(config types.GameConfig, moniker string) {
	r.board = types.NewEmptyBoard(config.Map)
	r.resources = types.NewResourceSet(config.Initial.Resources)
	r.config = config
	r.moniker = moniker
	r.step = 0
}

func (r *RandomBot) Update(newState types.PartialState) UpdateResponse {
	for posIdx, tile := range newState.Map {
		x, y := int(posIdx % r.config.Map.Width), int(posIdx / r.config.Map.Width)
		r.board[x][y] = *tile
	}

	resp := UpdateResponse{}

	// make a bunch of random builds and moves
	for x := 0; x < int(r.config.Map.Width); x++ {
		for y := 0; y < int(r.config.Map.Height); y++ {
			if r.board[x][y].Landscape != types.Landscape_UNKNOWN {
				// check if bot occupies this tile
				if r.board[x][y].Faction != nil && r.board[x][y].Faction.Moniker == r.moniker { 
					if rand.Intn(buildProb) % buildProb == 0 {
						settlement := settlements[rand.Intn(len(settlements))]
						msg := &types.MsgBuild{Settlement: settlement, Position: &types.Position{X: uint32(x), Y: uint32(y)}}
						resp.Builds = append(resp.Builds, msg)
					} else if rand.Intn(moveProb) % moveProb == 0 {
						direction := directions[rand.Intn(len(directions))]
						quantity := r.board[x][y].Population 
						if r.board[x][y].Settlement != types.Settlement_NONE {
							quantity-- // leave one populace at the settlement
						}
						msg := &types.MsgMove{Position: &types.Position{X: uint32(x), Y: uint32(y)}, Direction: direction, Quantity: quantity}
						resp.Moves = append(resp.Moves, msg)
					}
					// else this populace does nothing for this turn
				}
			}
		}
	}

	return resp
}