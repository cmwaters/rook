package bot

import (
	"github.com/cmwaters/rook/x/rook/types"
)

var _ Bot = &Computer{}

type Computer struct {
	Board     [][]types.Tile // 2D map is serialized (makes it easier to read UpdatedTiles)
	Faction types.Faction
	Config  types.GameConfig
	Step    uint32
}

func NewComputer() *Computer {
	return &Computer{} // uses all zero values which will be populated upon init
}

// TODO: stub out init and update functions
func (c *Computer) Init(req InitRequest) {
	c.Board = types.NewEmptyBoard(req.Config.Map)
	c.Faction = req.Faction
	c.Config = req.Config
	c.Step = 0
}

func (c *Computer) Update(req UpdateRequest) UpdateResponse {

	return UpdateResponse{}
}
