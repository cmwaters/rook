package types

import (
	"errors"
	"fmt"
	// "math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const defaultCountdownBlocks = 50

type GameState struct {
	Map      [][]Tile //x then y
	Players  map[string]*Faction
	Factions []*Faction
	Config   *GameConfig
}

func NewGameState(players []*sdk.AccAddress, config *GameConfig) *GameState {
	game := &GameState{
		Players:  make(map[string]*Faction, len(players)),
		Factions: make([]*Faction, len(players)),
		Config:   config,
	}
	// create teams and initialize factions
	game.initiatePlayers(players)

	// generate a random map
	game.GenerateMap()

	// add capitals of each faction to the map
	game.PopulateFactions()
	return game
}

// TODO: create the random map generator algorithm
func (g *GameState) GenerateMap() {
	// create a random generator from the seed
	// randGen := rand.New(rand.NewSource(g.Config.Map.Seed))

	// populate the map
}

// InitiatePlayers creates a faction for each player
//
// In the future we will need to account for teams.
func (g *GameState) initiatePlayers(players []*sdk.AccAddress) {
	for idx, player := range players {
		g.Factions[idx] = NewFaction(player.String())
		g.Players[player.String()] = g.Factions[idx]
	}
}

// TODO: populate the map with the factions capital
func (g *GameState) PopulateFactions() {

}

func (g *GameState) Build(faction *Faction, settlement Settlement, position *Position) error {
	// validate position
	if position.X >= g.Config.Map.Width {
		return fmt.Errorf("position X (%d) exceeds game boundaries (%d)", position.X, g.Config.Map.Width)
	}
	if position.Y >= g.Config.Map.Height {
		return fmt.Errorf("position Y (%d) exceeds game boundaries (%d)", position.Y, g.Config.Map.Height)
	}

	// check that the faction occupies this piece of land
	if g.Map[position.X][position.Y].Faction != faction {
		return fmt.Errorf("faction %s does not occupy land (x: %d, y: %d)",
			faction.Moniker, position.X, position.Y)
	}

	// check that the faction has the allocated resources
	if !ConstructionResources(g.Config.Construction, settlement).Less(faction.Resources) {
		return fmt.Errorf("not enough resources. Required: %s, got: %s",
			ConstructionResources(g.Config.Construction, settlement), faction.Resources)
	}

	// check all the settlement specific rules
	// lumbermills must be on a forest
	if settlement == Settlement_LUMBERMILL && g.Map[position.X][position.Y].Landscape != Landscape_FOREST {
		return errors.New("building a lumbermill must be over a forest")
	}

	// A quarry must be built adjacent to at least one mountain range
	isQuarry := func (tile *Tile) bool { return tile.Landscape == Landscape_MOUNTAINS}
	if settlement == Settlement_QUARRY && len(g.searchAdjacent(position, isQuarry)) == 0 {
		return errors.New("building a quarry must be adjacent to mountains")
	}

	// other settlements can't be in a forest
	if settlement != Settlement_LUMBERMILL && g.Map[position.X][position.Y].Landscape != Landscape_FOREST {
		return fmt.Errorf("building a %v can't be in a forest", settlement)
	}

	// we are good to build the settlement
	faction.Settlements[position.Index(g.Config.Map)] = settlement

	return nil
}

func (g *GameState) Move(faction *Faction, quantity uint32, origin *Position, direction Direction) error {
	return nil
}

var directions = []Direction{Direction_LEFT, Direction_RIGHT, Direction_UP, Direction_DOWN}

func (g *GameState) searchAdjacent(position *Position, condition func(*Tile) bool) []*Tile {
	tiles := make([]*Tile, 0, 4)
	for _, direction := range directions {
		tile := g.neighborTile(position, direction)
		if condition(tile) {
			tiles = append(tiles, tile)
		}
	}
	return tiles
}

func (g *GameState) neighborTile(position *Position, direction Direction) *Tile {
	switch direction {
	case Direction_LEFT:
		if position.X == 0 {
			return &g.Map[g.Config.Map.Width - 1][position.Y]
		}
		return &g.Map[position.X - 1][position.Y]
	case Direction_RIGHT:
		if position.X == g.Config.Map.Width - 1 {
			return &g.Map[0][position.Y]
		}
		return &g.Map[position.X + 1][position.Y]
	case Direction_UP:
		if position.Y == g.Config.Map.Height - 1 {
			return &g.Map[position.X][0]
		}
		return &g.Map[position.X][position.Y + 1]
	case Direction_DOWN:
		if position.X == 0 {
			return &g.Map[position.X][g.Config.Map.Height - 1]
		}
		return &g.Map[position.X][position.Y - 1]
	default: //unknown direction
		return nil
	}
}

type PendingGameState struct {
	playerVotes map[string]bool
	config      *GameConfig
	gameID      string

	total uint32
	count uint32
	timer uint32 // countdown where the pending game is available
}

func NewPendingGame(players []sdk.AccAddress, gameId string, config *GameConfig) *PendingGameState {
	playerMap := make(map[string]bool, len(players))
	for _, p := range players {
		playerMap[p.String()] = false
	}
	return &PendingGameState{
		playerVotes: playerMap,
		config:      config,
		gameID:      gameId,
		total:       uint32(len(players)),
		count:       0,
		timer:       defaultCountdownBlocks,
	}
}

func (p *PendingGameState) Tick() (expired bool, ready bool, players []*sdk.AccAddress) {
	p.timer--
	if p.timer == 0 {
		ready, players := p.Quorum()
		return true, ready, players
	}
	return false, false, nil
}

func (p *PendingGameState) AddVote(player *sdk.AccAddress) bool {
	if _, ok := p.playerVotes[player.String()]; !ok {
		p.playerVotes[player.String()] = true
		p.count++
		if p.count == p.total {
			return true
		}
	}
	return false
}

func (p *PendingGameState) Participants() []*sdk.AccAddress {
	players := make([]*sdk.AccAddress, 0, p.total)
	for playerString, vote := range p.playerVotes {
		if vote {
			player, _ := sdk.AccAddressFromBech32(playerString)
			players = append(players, &player)
		}
	}
	return players
}

func (p *PendingGameState) Quorum() (bool, []*sdk.AccAddress) {
	counter := uint32(0)
	players := make([]*sdk.AccAddress, 0, p.total)
	for playerString, vote := range p.playerVotes {
		if vote {
			counter++
			player, _ := sdk.AccAddressFromBech32(playerString)
			players = append(players, &player)
		}
	}
	if counter > (p.total / 2) {
		return true, players
	}
	return false, nil
}

func (p *PendingGameState) Config() *GameConfig {
	return p.config
}
