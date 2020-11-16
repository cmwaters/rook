package bot

import (
	"github.com/cmwaters/rook/x/rook/types"
)

type Bot interface {
	Init(types.GameConfig, string)
	Update(types.PartialState) UpdateResponse
}

type UpdateResponse struct {
	Moves  []*types.MsgMove
	Builds []*types.MsgBuild
}
