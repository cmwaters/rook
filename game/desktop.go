package game

import (
	"fmt"

	"github.com/cmwaters/rook/x/rook/types"
	"github.com/hajimehoshi/ebiten"
)

const (
	menuView            = "menu"
	gameView            = "game"
	defaultScreenWidth  = 640
	defaultScreenHeight = 480
)

// RookDesktop implements ebiten.Game interface.
type RookDesktop struct {
	ScreenWidth, ScreenHeight int

	views   map[string]View // the desktop switches between multiple views
	current View
}

func NewRookDesktop() *RookDesktop {
	config := types.DefaultGameConfig()
	r := &RookDesktop{
		ScreenWidth:  defaultScreenWidth,
		ScreenHeight: defaultScreenHeight,
		views: map[string]View{
			gameView: NewLocalGameView(config, 3),
		},
	}
	// start with game view (in the future we will start with the menu)
	r.current = r.views[gameView]
	return r
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (r *RookDesktop) Update(screen *ebiten.Image) error {
	// Write your game's logical update.
	view, err := r.current.Update(r.views)
	if err != nil {
		return err
	}
	if view != nil {
		r.current = view
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (r *RookDesktop) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	r.current.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (r *RookDesktop) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	r.ScreenWidth, r.ScreenHeight = outsideWidth, outsideHeight
	if game, ok := r.current.(*GameView); ok {
		game.camera.width = float64(outsideWidth)
		game.camera.height = float64(outsideHeight)
	}
	return r.ScreenWidth, r.ScreenHeight
}

func (r *RookDesktop) changeView(viewStr string) error {
	if view, ok := r.views[viewStr]; ok {
		r.current = view
		return nil
	}
	return fmt.Errorf("view %s not found", viewStr)
}
