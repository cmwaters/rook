package types

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GameState struct {
	Map [][]Tile
	
	players map[string]*Faction
	factions []*Faction
	config *GameConfig
}


func InitGame(players []*sdk.AccAddress, config *GameConfig) *GameState {
	game := &GameState{
		players: make(map[string]*Faction, len(players)),
		factions: make([]*Faction, len(players)),
		config: config,
	}

	// generate a random map
	game.GenerateMap()

	// create teams and initialize factions
	game.InitiatePlayers(players)

	// add capitals of each faction to the map
	game.PopulateFactions()
}

func (g *GameState) GenerateMap()  {
	// create a random generator from the seed
	randGen := rand.New(rand.NewSource(g.config.Map.Seed))

	// populate the map
}

// InitiatePlayers creates a faction for each player
//
// In the future we will need to account for teams. 
func (g *GameState) InitiatePlayers(players []*sdk.AccAddress) {
	for idx, player := range players {
		g.factions[idx] = NewFaction(player.String())
		g.players[player.String()] = g.factions[idx]
	}
}

func (g *GameState) PopulateFactions() {



}

