package main

import (
	"log"

	g "github.com/cmwaters/rook/game"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	game := g.NewRookDesktop()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Rook")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
