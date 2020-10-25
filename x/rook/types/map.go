package types

func GenerateMap(config *MapConfig) [][]Tile {
	// TODO: actually create a map
	return NewEmptyBoard(config)
}

func PopulateFactions(gameMap *[][]Tile, factions []*Faction, config *MapConfig) {
	// TODO: populate an existing map with the capital cities of each faction
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
