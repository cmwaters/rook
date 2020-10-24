package game

import (
	"github.com/hajimehoshi/ebiten"
)

var _ View = &GameView{}

type GameView struct {

}

func NewGameView() *GameView {
	return &GameView{}
}

func (g *GameView) Update(views map[string]View) (View, error) {
	return views[gameView], nil
}

func (g *GameView) Draw(screen *ebiten.Image) {
	
}