package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	g "github.com/cmwaters/rook/game"
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

