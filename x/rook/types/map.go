package types


func GenerateMap(config *MapConfig) [][]Tile {
	// TODO: actually create a map
	return NewEmptyBoard(config)
}

func PopulateFactions(gameMap *[][]Tile, factions []*Faction, config *MapConfig) {

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