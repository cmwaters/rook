package game

import (
	"github.com/hajimehoshi/ebiten"
)

// RookDesktop implements ebiten.Game interface.
type RookDesktop struct{
	Screens map[string]*View // the desktop switches between multiple views
}

func NewRookDesktop() *RookDesktop {
	return &RookDesktop{}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (r *RookDesktop) Update(screen *ebiten.Image) error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (r *RookDesktop) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (r *RookDesktop) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}



