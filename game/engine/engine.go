package engine

import (
	"github.com/cmwaters/rook/x/rook/types"
)

// GameEngine can be run either locally or with the network. The engine handles all state transitions
// between parties and updates players state accordingly.
type GameEngine interface {
	Init(chan *types.PartialState)
	Build(settlement types.Settlement, x, y int) error
	Move(quantity uint32, x, y int, direction types.Direction) error
}

var _ GameEngine = &LocalGameEngine{}
var _ GameEngine = &NetworkGameEngine{}

type BuildSettlement struct {
	settlement types.Settlement
	x, y       int
}
