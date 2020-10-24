package game

import "github.com/hajimehoshi/ebiten"

type Tile struct {
	image *ebiten.Image
	x, y int
}

func NewTile(x, y, width, height int) *Tile {
	image, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	image.Fill(frameColor)
	return &Tile{image, x, y}
}

func (t *Tile) Draw(screen *ebiten.Image) {

}