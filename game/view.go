package game

import (
	"github.com/hajimehoshi/ebiten"
)

type View interface {
	Update(map[string]View) (View, error)
	Draw(*ebiten.Image)
}