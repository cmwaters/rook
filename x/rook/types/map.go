package types

import (
	"fmt"
	"math/rand"
)

var (
	// neighbor weightings. We want the heighest waiting to go to
	// tiles with a single neighbor because this increases exploration
	firstBucket  = 8
	secondBucket = 3
)

func GenerateMap(config *MapConfig) [][]Tile {
	// first populate the entire map with mountains, lakes and forests
	randGen := rand.New(rand.NewSource(config.Seed))
	terrainDensityTotal := config.MountainsDensity + config.ForestDensity + config.LakeDensity
	board := make([][]Tile, config.Width)
	for x := 0; x < int(config.Width); x++ {
		board[x] = make([]Tile, config.Height) // essentially a column
		for y := 0; y < int(config.Height); y++ {
			landIndex := randGen.Intn(int(terrainDensityTotal))
			var chosenLand Landscape
			if uint32(landIndex) >= (config.MountainsDensity + config.ForestDensity) {
				chosenLand = Landscape_LAKE
			} else if uint32(landIndex) >= config.MountainsDensity {
				chosenLand = Landscape_FOREST
			} else {
				chosenLand = Landscape_MOUNTAINS
			}
			board[x][y] = Tile{
				Landscape:  chosenLand,
				Settlement: Settlement_NONE,
				Population: 0,
				Faction:    nil,
			}
		}
	}
	// amount of tiles we want to be plains
	plainsCount := (float64(config.PlainsDensity) / float64(terrainDensityTotal+config.PlainsDensity)) *
		float64(config.Width*config.Height)
	fmt.Println(plainsCount)
	// use a path generation algorithm to inlaid the plains. This ensures
	// that all plains are connected and that all civilizations can reach one another
	startingPos := Position{config.Width / 2, config.Height / 2}
	firstBucketSet := make([]Position, 0, int(plainsCount))
	secondBucketSet := make([]Position, 0, int(plainsCount))
	isNotPlains := func(land Landscape) bool {
		return land != Landscape_PLAINS
	}
	// setup the first two tiles
	board[startingPos.X][startingPos.Y].Landscape = Landscape_PLAINS
	positions := possibleDirections(board, startingPos, config.Width, config.Height, isNotPlains)
	// select a new pos out of the possible positions
	newPos := positions[randGen.Intn(len(positions))]
	board[newPos.X][newPos.Y].Landscape = Landscape_PLAINS
	firstBucketSet = []Position{startingPos, newPos}
	for num := 0; num < int(plainsCount); num++ {
		// select from a random bucket
		bucketRandIndex := randGen.Intn(firstBucket + secondBucket)
		if bucketRandIndex > firstBucket && len(secondBucketSet) != 0 { // use second bucket
			chosenIndex := randGen.Intn(len(secondBucketSet))
			chosenPos := secondBucketSet[chosenIndex]
			positions := possibleDirections(board, chosenPos, config.Width, config.Height, isNotPlains)
			if len(positions) < 2 {
				// this tile is surrounded so we should remove it from the neighbors set and start again
				secondBucketSet = remove(secondBucketSet, chosenIndex)
				num--
				continue
			}
			newPos := positions[randGen.Intn(len(positions))]
			board[newPos.X][newPos.Y].Landscape = Landscape_PLAINS
			firstBucketSet = append(firstBucketSet, newPos)
			secondBucketSet = remove(secondBucketSet, chosenIndex)
		} else if len(firstBucketSet) != 0 { // use first bucket
			fmt.Println(len(firstBucketSet))
			chosenIndex := randGen.Intn(len(firstBucketSet))
			chosenPos := firstBucketSet[chosenIndex]
			positions := possibleDirections(board, chosenPos, config.Width, config.Height, isNotPlains)
			if len(positions) < 2 {
				// this tile is surrounded so we should remove it from the neighbors set and start again
				firstBucketSet = remove(firstBucketSet, chosenIndex)
				num--
				continue
			}
			newPos := positions[randGen.Intn(len(positions))]
			board[newPos.X][newPos.Y].Landscape = Landscape_PLAINS
			firstBucketSet = append(firstBucketSet, newPos)
			firstBucketSet = remove(firstBucketSet, chosenIndex)
			secondBucketSet = append(secondBucketSet, chosenPos)
		} else {
			num--
		}
	}
	return board
}

func remove(s []Position, i int) []Position {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func possibleDirections(board [][]Tile, pos Position, width, height uint32, condition func(Landscape) bool) []Position {
	positions := make([]Position, 0, 4)
	// check left
	if pos.X == 0 && condition(board[width-1][pos.Y].Landscape) {
		positions = append(positions, Position{width - 1, pos.Y})
	} else if pos.X > 0 && condition(board[pos.X-1][pos.Y].Landscape) {
		positions = append(positions, Position{pos.X - 1, pos.Y})
	}
	// check right
	if pos.X == width-1 && condition(board[0][pos.Y].Landscape) {
		positions = append(positions, Position{0, pos.Y})
	} else if pos.X < width-1 && condition(board[pos.X+1][pos.Y].Landscape) {
		positions = append(positions, Position{pos.X + 1, pos.Y})
	}
	// check above
	if pos.Y == 0 && condition(board[pos.X][height-1].Landscape) {
		positions = append(positions, Position{pos.X, height - 1})
	} else if pos.Y > 0 && condition(board[pos.X][pos.Y-1].Landscape) {
		positions = append(positions, Position{pos.X, pos.Y - 1})
	}
	// check below
	if pos.Y == height-1 && board[pos.X][0].Landscape != Landscape_PLAINS {
		positions = append(positions, Position{pos.X, height - 1})
	} else if pos.Y < height-1 && board[pos.X][pos.Y+1].Landscape != Landscape_PLAINS {
		positions = append(positions, Position{pos.X, pos.Y + 1})
	}
	return positions
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
	allocation := len(possibleTiles) / len(g.Factions)
	for x, faction := range g.Factions {
		chosenTile := possibleTiles[x*allocation+randGen.Intn(allocation)]
		g.Map[chosenTile.X][chosenTile.Y].Settlement = Settlement_CAPITAL
		g.Map[chosenTile.X][chosenTile.Y].Faction = faction
		g.Map[chosenTile.X][chosenTile.Y].Population = g.Config.Initial.Population
		index := (chosenTile.Y * g.Config.Map.Width) + chosenTile.X
		faction.Settlements[index] = Settlement_CAPITAL
		faction.Population[index] = g.Config.Initial.Population
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
				positions = append(positions, Position{X: uint32(x), Y: uint32(y)})
			}
		}
	}
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
			tiles[uint32((y*width)+x)] = &gameMap[x][y]
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
