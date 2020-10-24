package game

import (
	"image/color"

	"github.com/cmwaters/rook/x/rook/types"
	"github.com/hajimehoshi/ebiten"
)

var _ View = &GameView{}

var (
	backgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	frameColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
	tileWidth				= 64
	tileHeight			= 64
	tileMargin		 	= 2
)


type GameView struct {
	mapImage *ebiten.Image
	camera *Camera
	board *Board
	tile *Tile
	config *types.GameConfig
}

func NewGameView(config *types.GameConfig) *GameView {
	return &GameView{
		mapImage: NewMapImage(config.Map),
		tile: NewTile(0, 0, 64, 64),
		camera: NewCamera(0, 0, 0.3, 5),
		config: config,
	}
}

func (g *GameView) Update(views map[string]View) (View, error) {
	g.camera.ParseMovementKeys()
	return views[gameView], nil
}

func (g *GameView) Draw(screen *ebiten.Image) {
	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(float64(g.tile.x), float64(g.tile.y))
	g.board.Draw(g.mapImage)
	_ = screen.DrawImage(g.mapImage, g.camera.Update())
}




func NewMapImage(config *types.MapConfig) *ebiten.Image {
	width := tileMargin + int(config.Width) * (tileWidth + tileMargin)
	height := tileMargin + int(config.Height) * (tileHeight + tileMargin)
	image, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	return image
}
