package bot

import (
	"github.com/cmwaters/rook/x/rook/types"
)

type Bot interface {
	Init(InitRequest)
	Update(UpdateRequest) UpdateResponse
}

type InitRequest struct {
	Config  types.GameConfig
	Faction types.Faction
}

type UpdateRequest struct {
	UpdatedTiles map[uint32]types.Tile
}

type UpdateResponse struct {
	Moves  []*types.MsgMove
	Builds []*types.MsgBuild
}
