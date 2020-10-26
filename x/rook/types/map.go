package types

import (
	"math/rand"
)

func GenerateMap(config *MapConfig) [][]Tile {
	// TODO: actually create a map
	return NewEmptyBoard(config)
}

type tileInfo struct {
	Tile *Tile
	X, Y uint32
}

// PopulateFactions takes a map and adds the initial capital city somewhere randomly in the map
func (g *GameState) PopulateFactions() {
	randGen := rand.New(rand.NewSource(g.Config.Map.Seed))
	// make a list of all the possible places
	var possibleTiles []tileInfo
	for x, column := range g.Map {
		for y, tile := range column {
			if tile.Landscape == Landscape_PLAINS {
				possibleTiles = append(possibleTiles, tileInfo{Tile: &tile, X: uint32(x), Y: uint32(y)})
			}
		}
	}
	// by allocation each faction a portion of possible tiles that their capital will
	// reside we eliminate the possibility of collision. 
	// NOTE: This could be improved as capitals could technically start right next to one another
	allocation := len(possibleTiles)/len(g.Factions)
	for x, faction := range g.Factions {
		chosenTile := possibleTiles[x * allocation + randGen.Intn(allocation)]
		chosenTile.Tile.Faction = faction
		chosenTile.Tile.Population = g.Config.Initial.Population
		index := (chosenTile.Y * g.Config.Map.Width) + chosenTile.X
		faction.Settlements[index] = Settlement_CAPITAL
		faction.Population[index] = g.Config.Initial.Population
	}
}

// GetVisibleTilesFromMap loops through the position of all population and returns the visible tiles
//
// TODO: in the future we should only take the delta based on the movements that the player took
// in that turn.
func GetVisibleTilesFromMap(gameMap [][]Tile, faction Faction) map[uint32]*Tile {
	return make(map[uint32]*Tile)
}

func NewEmptyBoard(config *MapConfig) [][]Tile {
	board := make([][]Tile, config.Width)
	for x := 0; x < int(config.Width); x++ {
		board[x] = make([]Tile, config.Height) // essentially a column
		for y := 0; y < int(config.Height); y++ {
			board[x][y] = EmptyTile()
		}
	}
	return board
}

func EmptyTile() Tile {
	return Tile{
		Landscape:  Landscape_UNKNOWN,
		Settlement: Settlement_NONE,
		Population: 0,
		Faction:    nil,
	}
}
