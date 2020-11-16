package bot

import (
	"github.com/cmwaters/rook/x/rook/types"
)

var _ Bot = &Computer{}

type Computer struct {
	board   [][]types.Tile // 2D map is serialized (makes it easier to read UpdatedTiles)
	faction types.Faction
	config  types.GameConfig
	moniker string
	step    uint32
}

func NewComputer() *Computer {
	return &Computer{} // uses all zero values which will be populated upon init
}

// TODO: stub out init and update functions
func (c *Computer) Init(config types.GameConfig, moniker string) {
	c.board = types.NewEmptyBoard(config.Map)
	c.faction = *types.NewFaction("computer", *config.Initial)
	c.config = config
	c.step = 0
}

func (c *Computer) Update(newState types.PartialState) UpdateResponse {

	return UpdateResponse{}
}
