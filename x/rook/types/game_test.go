package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
	// "github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	
)

func TestNewGameState(t *testing.T) {
	N := 4
	game := &GameState{
		Players:  make(map[string]*Faction, N), // we won't really be using this for testing
		Factions: make([]*Faction, N),
		Config:   DefaultGameConfig(),
	}
	// create teams and initialize factions
	for i := 0; i < N; i++ {
		game.Factions[i] = NewFaction(tmrand.Str(10), *game.Config.Initial)
	}

	// generates a random map from the config. seed cannot be 0
	game.Map = GenerateMap(game.Config.Map)

	// add capitals of each faction to the map. Note that we pass
	// a reference to the map so this function can change it directly
	game.PopulateFactions()

	width := game.Config.Map.Width
	for _, faction := range game.Factions {
		if assert.EqualValues(t, 1, len(faction.Settlements)) {
			for pos, settlement := range faction.Settlements {
				x, y := pos % width, pos / width
				assert.EqualValues(t, game.Map[x][y].Settlement, settlement)
			}
		}
		assert.Equal(t, 9, len(GetVisibleTilesFromMap(game.Map, *faction, game.Config.Map)))
	}

}