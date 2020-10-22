package types

import (
	// "math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const defaultCountdownBlocks = 50;

type GameState struct {
	Map [][]Tile
	Players map[string]*Faction
	Factions []*Faction
	Config *GameConfig
}

func NewGameState(players []*sdk.AccAddress, config *GameConfig) *GameState {
	game := &GameState{
		Players: make(map[string]*Faction, len(players)),
		Factions: make([]*Faction, len(players)),
		Config: config,
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
func (g *GameState) GenerateMap()  {
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


type PendingGameState struct {
	playerVotes map[string]bool
	config  *GameConfig
	gameID string

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
		config: config,
		gameID: gameId,
		total: uint32(len(players)),
		count: 0,
		timer: defaultCountdownBlocks,
	}
}

func (p *PendingGameState) Tick() (bool, []*sdk.AccAddress) { 
	p.timer--
	if p.timer == 0 {
		return p.Quorum()
	}
	return false, nil
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


