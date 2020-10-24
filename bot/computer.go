package bot

import (
	"github.com/cmwaters/rook/x/rook/types"
)

var _ Bot = &Computer{}

type Computer struct {
	Map     []types.Tile // 2D map is serialized (makes it easier to read UpdatedTiles)
	Faction types.Faction
	Config  types.GameConfig
	Step    uint32
}

func NewComputer() *Computer {
	return &Computer{} // uses all zero values which will be populated upon init
}

// TODO: stub out init and update functions
func (c *Computer) Init(req InitRequest) {
	c.Map = EmptyMap(req.Config.Map.Width, req.Config.Map.Height)
	c.Faction = req.Faction
	c.Config = req.Config
	c.Step = 0
}

func (c *Computer) Update(req UpdateRequest) UpdateResponse {

	return UpdateResponse{}
}

func EmptyMap(width, height uint32) []types.Tile {
	gameMap := make([]types.Tile, width*height)
	for index := uint32(0); index < (width * height); index++ {
		gameMap[index] = EmptyTile()
	}
	return gameMap
}

func EmptyTile() types.Tile {
	return types.Tile{
		Landscape:  types.Landscape_L_UNKNOWN,
		Population: 0,
		Faction:    nil,
	}
}
