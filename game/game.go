package game

import (
	"image/color"
	"github.com/hajimehoshi/ebiten"
)

var _ View = &GameView{}

var (
	backgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	frameColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
)


type GameView struct {
	boardImage *ebiten.Image
}

func NewGameView() *GameView {
	return &GameView{}
}

func (g *GameView) Update(views map[string]View) (View, error) {
	return views[gameView], nil
}

func (g *GameView) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		w, h := 64, 64
		g.boardImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
		g.boardImage.Fill(frameColor)
	}
	screen.Fill(backgroundColor)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := g.boardImage.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}
