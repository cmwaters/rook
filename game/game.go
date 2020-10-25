package game

import (
	e "github.com/cmwaters/rook/game/engine"
	"github.com/cmwaters/rook/x/rook/types"
	"github.com/hajimehoshi/ebiten"
)

var _ View = &GameView{}

var (
	tileWidth				= 64
	tileHeight			= 64
	tileMargin		 	= 2
)


type GameView struct {
	mapImage *ebiten.Image
	camera *Camera
	board [][]types.Tile // this represents the players known map
	config *types.GameConfig
	op *ebiten.DrawImageOptions
	engine e.GameEngine
}

func NewLocalGameView(config *types.GameConfig, bots int) *GameView {
	g := &GameView{
		mapImage: NewMapImage(config.Map),
		camera: NewCamera(0, 0, 0.3, 5),
		config: config,
		board: types.NewEmptyBoard(config.Map),
		op: &ebiten.DrawImageOptions{},
		engine: e.NewLocalGameEngine(config, bots),
	}
	statec := make(chan *types.PartialState)
	g.engine.Init(statec)
	go g.ApplyStateTransitions(statec)
	return g
}

func (g *GameView) Update(views map[string]View) (View, error) {
	g.camera.ParseMovementKeys()
	return views[gameView], nil
}

func (g *GameView) Draw(screen *ebiten.Image) {
	_ = screen.Fill(shadowColor)
	for x := 0; x < int(g.config.Map.Width); x++ {
		for y := 0; y < int(g.config.Map.Height); y++ {
			posX := tileMargin + (x * (tileWidth + tileMargin))
			posY := tileMargin + (y * (tileHeight + tileMargin))
			g.op.GeoM.Reset()
			g.op.GeoM.Translate(float64(posX), float64(posY))
			_ = g.mapImage.DrawImage(plainsSprite, g.op)
		}
	}
	_ = screen.DrawImage(g.mapImage, g.camera.Update())
}

func (g *GameView) ApplyStateTransitions(c chan *types.PartialState) {
	for {

	}
}




func NewMapImage(config *types.MapConfig) *ebiten.Image {
	width := tileMargin + int(config.Width) * (tileWidth + tileMargin)
	height := tileMargin + int(config.Height) * (tileHeight + tileMargin)
	image, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	image.Fill(shadowColor)
	return image
}
