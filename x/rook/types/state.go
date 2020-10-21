package rook


type State struct {
	Board   [][]Tile
	Players map[string]Civilization

	Step int16
}



type Tile struct {
	landscape  Landscape
	population int
	occupied   *Civilization
}

