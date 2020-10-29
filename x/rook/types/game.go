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

	// generates a random map from the config. seed cannot be 0
	game.Map = GenerateMap(game.Config.Map)

	// add capitals of each faction to the map. Note that we pass
	// a reference to the map so this function can change it directly
	game.PopulateFactions()

	return game
}

// Build validates a build msg from a faction and applies it to the current state
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
	isQuarry := func(tile *Tile) bool { return tile.Landscape == Landscape_MOUNTAINS }
	if settlement == Settlement_QUARRY && len(g.searchAdjacent(position, isQuarry)) == 0 {
		return errors.New("building a quarry must be adjacent to mountains")
	}

	// other settlements can't be in a forest
	if settlement != Settlement_LUMBERMILL && g.Map[position.X][position.Y].Landscape != Landscape_FOREST {
		return fmt.Errorf("building a %v can't be in a forest", settlement)
	}

	//TODO: add logic for cities and capitals which must be built on top of towns and cities respectively

	// we are good to build the settlement
	faction.Settlements[position.Index(g.Config.Map)] = settlement
	g.Map[position.X][position.Y].Settlement = settlement
	// and remove the resources from the faction
	faction.Resources.Subtract(ConstructionResources(g.Config.Construction, settlement))

	return nil
}

// Move validates a move msg and applies it to the current state
func (g *GameState) Move(faction *Faction, quantity uint32, origin *Position, direction Direction) error {
	// validate position
	if origin.X >= g.Config.Map.Width {
		return fmt.Errorf("origin X (%d) exceeds game boundaries (%d)", origin.X, g.Config.Map.Width)
	}
	if origin.Y >= g.Config.Map.Height {
		return fmt.Errorf("origin Y (%d) exceeds game boundaries (%d)", origin.Y, g.Config.Map.Height)
	}
	
	// check that there is enough population at that tile to make the move
	posIndex := (origin.Y * g.Config.Map.Width) + origin.X
	if populace, ok := faction.Population[posIndex]; ok {
		if quantity > populace {
			return fmt.Errorf("more populace requested to move than is actually there. Max %d", populace)
		}
	} else {
		return fmt.Errorf("no populace belonging to faction %s at pos x: %d, y: %d", 
		faction.Moniker, origin.X, origin.Y)
	}
	
	// ensure that the destination tile can be moved to i.e. not mountains or lake
	destination, destPos := g.neighborTile(origin, direction)
	if destination == nil {
		return fmt.Errorf("unknown direction %v requested ", direction)
	}
	if destination.Landscape == Landscape_MOUNTAINS || destination.Landscape == Landscape_LAKE {
		return fmt.Errorf("cannot move populace to %s", destination.Landscape)
	}
	destIndex := (destPos.Y * g.Config.Map.Width) + destPos.Y
	
	// check to see if the move is an attack
	if destination.Faction != nil && destination.Faction != faction {
		if destination.Population > quantity { // defenders win
			destination.Population -= quantity
			destination.Faction.Population[destIndex] -= quantity
		} else if quantity > destination.Population { // attackers win (we need to replace the tile)
			g.Map[destPos.X][destPos.Y] = Tile{
				Population: quantity - destination.Population,
				Faction: faction,
				Landscape: destination.Landscape,
				Settlement: destination.Settlement,
			}
			faction.Population[destIndex] = quantity - destination.Population
			delete(destination.Faction.Population, destIndex)
			if destination.Settlement != Settlement_NONE { // player has captured an opponents settlement
				delete(destination.Faction.Settlements, destIndex)
				faction.Settlements[destIndex] = destination.Settlement
				if destination.Settlement == Settlement_CAPITAL {
					// TODO: opponent might have been conquered. We need to add logic for this
				}
			}
		} else { // stalemate
			g.Map[destPos.X][destPos.Y] = Tile{
				Population: 0,
				Faction: nil,
				Landscape: destination.Landscape,
				Settlement: destination.Settlement,
			}
			delete(destination.Faction.Population, destIndex)
			delete(destination.Faction.Settlements, destIndex)
		}
	} else if destination.Faction == nil { // player now occupies this land
		destination.Faction = faction
		destination.Population = quantity
		faction.Population[destIndex] = quantity
	} else { // player is transferring troops between two lands that the player already occupies
		destination.Population += quantity
	}
	
	// update the tile that was left behind
	if quantity == g.Map[origin.X][origin.Y].Population { // all populace moved
		if g.Map[origin.X][origin.Y].Settlement != Settlement_NONE { // player is abandoning their own settlement!
			delete(faction.Settlements, posIndex)
		}
		delete(faction.Population, posIndex)
		land := g.Map[origin.X][origin.Y].Landscape
		settlement := g.Map[origin.X][origin.Y].Settlement
		g.Map[origin.X][origin.Y] = Tile{
			Population: 0,
			Faction: nil,
			Landscape: land,
			Settlement: settlement,
		}
	} else { // only some of the populace moved
		g.Map[origin.X][origin.Y].Population -= quantity
		faction.Population[posIndex] = g.Map[origin.X][origin.Y].Population
	}
	
	return nil
}

// Victor assesses if someone has won and if so which faction
func (g *GameState) Victor() *Faction {

	return nil
}

func (g *GameState) UpdateResources() {
	for _, faction := range g.Factions {
		faction.Reap(g.Config.Production)
		// reflect population changes on the map struct
		for pos, populace := range faction.Population {
			x, y := pos % g.Config.Map.Width, pos / g.Config.Map.Width
			g.Map[int(x)][int(y)].Population = populace
		}
	}
}

var directions = []Direction{Direction_LEFT, Direction_RIGHT, Direction_UP, Direction_DOWN}

func (g *GameState) searchAdjacent(position *Position, condition func(*Tile) bool) []*Tile {
	tiles := make([]*Tile, 0, 4)
	for _, direction := range directions {
		tile, _ := g.neighborTile(position, direction)
		if condition(tile) {
			tiles = append(tiles, tile)
		}
	}
	return tiles
}

func (g *GameState) neighborTile(position *Position, direction Direction) (*Tile, Position) {
	switch direction {
	case Direction_LEFT:
		if position.X == 0 {
			return &g.Map[g.Config.Map.Width-1][position.Y], Position{X: g.Config.Map.Width-1, Y: position.Y}
		}
		return &g.Map[position.X-1][position.Y], Position{X: position.X-1, Y: position.Y}
	case Direction_RIGHT:
		if position.X == g.Config.Map.Width-1 {
			return &g.Map[0][position.Y], Position{X: 0, Y: position.Y}
		}
		return &g.Map[position.X+1][position.Y], Position{X: position.X+1, Y: position.Y}
	case Direction_UP:
		if position.Y == g.Config.Map.Height-1 {
			return &g.Map[position.X][0], Position{X: position.X, Y: 0}
		}
		return &g.Map[position.X][position.Y+1], Position{X: position.X, Y: position.Y+1}
	case Direction_DOWN:
		if position.X == 0 {
			return &g.Map[position.X][g.Config.Map.Height-1], Position{X: position.X, Y: g.Config.Map.Height-1}
		}
		return &g.Map[position.X][position.Y-1], Position{X: position.X, Y: position.Y-1}
	default: //unknown direction
		return nil, Position{}
	}
}

// InitiatePlayers creates a faction for each player
//
// In the future we will need to account for teams.
func (g *GameState) initiatePlayers(players []*sdk.AccAddress) {
	for idx, player := range players {
		g.Factions[idx] = NewFaction(player.String(), *g.Config.Initial)
		g.Players[player.String()] = g.Factions[idx]
	}
}
