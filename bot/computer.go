package bot

import (
	"github.com/cmwaters/rook/x/rook/types"
)

var _ Bot = &Computer{}

type Computer struct {
	Board   [][]types.Tile // 2D map is serialized (makes it easier to read UpdatedTiles)
	Faction types.Faction
	Config  types.GameConfig
	Step    uint32
}

func NewComputer() *Computer {
	return &Computer{} // uses all zero values which will be populated upon init
}

// TODO: stub out init and update functions
func (c *Computer) Init(config types.GameConfig) {
	c.Board = types.NewEmptyBoard(config.Map)
	c.Faction = *types.NewFaction("computer", *config.Initial)
	c.Config = config
	c.Step = 0
}

func (c *Computer) Update(newState types.PartialState) UpdateResponse {

	return UpdateResponse{}
}
