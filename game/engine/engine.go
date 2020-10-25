package engine
// perhaps move this to it's own package and split out local and network engines

import (
	"github.com/cmwaters/rook/bot"
	"github.com/cmwaters/rook/x/rook/types"
)

type GameEngine interface {
	Init(chan *types.PartialState) 
	Build(settlement types.Settlement, x, y int) error
	Move(quantity, x, y int, direction types.Direction) error
}

var _ GameEngine = &LocalGameEngine{}
// var _ GameEngine = &NetworkGameEngine{}






