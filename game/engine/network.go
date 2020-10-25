package engine

import (
	"github.com/cmwaters/rook/x/rook/types"
)

// NetworkGameEngine acts as the interface between the player and the
// server parsing user actions into rest txs and queries to update
// state accordingly
type NetworkGameEngine struct {
	endpoint string
	playerAddress string
}

func (l *NetworkGameEngine) Init(state chan *types.PartialState) {

}

func (l *NetworkGameEngine) Build(settlement types.Settlement, x, y int) error {
	return nil
}

func (l *NetworkGameEngine) Move(quantity, x, y int, direction types.Direction) error {
	return nil
}