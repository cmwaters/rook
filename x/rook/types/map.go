package types

import (
	"fmt"
	"math/rand"
)

func GenerateMap(config *MapConfig) [][]Tile {
	// TODO: actually create a map
	board := make([][]Tile, config.Width)
	for x := 0; x < int(config.Width); x++ {
		board[x] = make([]Tile, config.Height) // essentially a column
		for y := 0; y < int(config.Height); y++ {
			board[x][y] = Tile{
				Landscape:  Landscape_PLAINS,
				Settlement: Settlement_NONE,
				Population: 0,
				Faction:    nil,
			}
		}
	}
	return board
}

// PopulateFactions takes a map and adds the initial capital city somewhere randomly in the map
func (g *GameState) PopulateFactions() {
	randGen := rand.New(rand.NewSource(g.Config.Map.Seed))
	// make a list of all the possible places
	var possibleTiles []Position
	for x, column := range g.Map {
		for y, tile := range column {
			if tile.Landscape == Landscape_PLAINS {
				possibleTiles = append(possibleTiles, Position{X: uint32(x), Y: uint32(y)})
			}
		}
	}
	// by allocation each faction a portion of possible tiles that their capital will
	// reside we eliminate the possibility of collision. 
	// NOTE: This could be improved as capitals could technically start right next to one another
	allocation := len(possibleTiles)/len(g.Factions)
	for x, faction := range g.Factions {
		chosenTile := possibleTiles[x * allocation + randGen.Intn(allocation)]
		g.Map[chosenTile.X][chosenTile.Y].Settlement = Settlement_CAPITAL
		g.Map[chosenTile.X][chosenTile.Y].Faction = faction
		g.Map[chosenTile.X][chosenTile.Y].Population = g.Config.Initial.Population
		index := (chosenTile.Y * g.Config.Map.Width) + chosenTile.X
		faction.Settlements[index] = Settlement_CAPITAL
		faction.Population[index] = g.Config.Initial.Population
		if (g.Map[chosenTile.X][chosenTile.Y].Settlement != Settlement_CAPITAL) {
			panic("hello")
		}
	}
}

// GetVisibleTilesFromMap loops through the position of all population and returns the visible tiles
//
// TODO: in the future we should only take the delta based on the movements that the player took
// in that turn.
func GetVisibleTilesFromMap(gameMap [][]Tile, faction Faction, config *MapConfig) map[uint32]*Tile {
	var positions = []Position{}
	for x := 0; x < len(gameMap); x++ {
		for y := 0; y < len(gameMap[x]); y++ {
			// fmt.Println(gameMap[x][y].Faction)
			if gameMap[x][y].Faction != nil && gameMap[x][y].Faction.Moniker == faction.Moniker {
				fmt.Printf("High jack\n")
				positions = append(positions, Position{X: uint32(x), Y: uint32(y)})
			}
		}
	}
	fmt.Printf("We have %d position\n", positions)
	visibleTiles := make(map[uint32]*Tile)
	for _, pos := range positions {
		lineOfSight := config.LineOfSight
		if gameMap[int(pos.X)][int(pos.Y)].Settlement == Settlement_ROOK {
			lineOfSight = config.RookLineOfSight
		}
		tiles := getNeighborTiles(gameMap, pos, int(lineOfSight))
		for index, tile := range tiles {
			visibleTiles[index] = tile
		}
	}
  return visibleTiles
}

func getNeighborTiles(gameMap [][]Tile, pos Position, size int) map[uint32]*Tile {
	width := len(gameMap)
	height := len(gameMap[0])
	tiles := make(map[uint32]*Tile)
	for deltaX := -size; deltaX <= size; deltaX++ {
		for deltaY := -size; deltaY <= size; deltaY++ {
			x := int(pos.X) + deltaX
			y := int(pos.Y) + deltaY
			// check for positions that are off the map
			if x < 0 || y < 0 || x >= width || y >= height {
				continue
			}
			tiles[uint32((y * width) + x)] = &gameMap[x][y]
		}
	}
	return tiles
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
