package game

import (
	"fmt"

	e "github.com/cmwaters/rook/game/engine"
	"github.com/cmwaters/rook/x/rook/types"
	"github.com/hajimehoshi/ebiten"
)

var _ View = &GameView{}

var (
	tileWidth  = 64
	tileHeight = 64
	tileMargin = 2
)

type GameView struct {
	mapImage *ebiten.Image
	camera   *Camera
	resources *types.ResourceSet
	board    [][]types.Tile // this represents the players known map
	config   *types.GameConfig
	op       *ebiten.DrawImageOptions
	engine   e.GameEngine
	step		 uint32
	active   *types.Tile
	sprites  []TileSprite
}

func NewLocalGameView(config *types.GameConfig, bots int) *GameView {
	fmt.Println("Starting local game")
	g := &GameView{
		mapImage: NewMapImage(config.Map),
		camera:   NewCamera(0, 0, 0.3, 5),
		config:   config,
		resources: types.NewResourceSet(config.Initial.Resources),
		board:    types.NewEmptyBoard(config.Map),
		op:       &ebiten.DrawImageOptions{},
		engine:   e.NewLocalGameEngine(config, bots),
		step: 		0,
	}
	statec := make(chan *types.PartialState)
	go g.engine.Init(statec)
	go g.ApplyStateTransitions(statec)
	return g
}

func (g *GameView) Update(views map[string]View) (View, error) {
	g.camera.ParseMovementKeys()
	return views[gameView], nil
}

func (g *GameView) Draw(screen *ebiten.Image) {
	_ = screen.Fill(shadowColor)
	for _, ts := range g.sprites {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(ts.x, ts.y)
		_ = g.mapImage.DrawImage(ts.color, g.op)
		_ = g.mapImage.DrawImage(ts.sprite, g.op)
		if ts.population != 0 {
			// TODO: write number over the top
		}
	}
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
	for stateUpdate := range c {
		fmt.Printf("update state to %d", stateUpdate.Step)
		// update map first
		for posIdx, tile := range stateUpdate.Map {
			x, y := IndexToCoordinate(posIdx, g.config.Map.Width)
			g.board[x][y] = *tile
		} 
		// update resources
		g.resources = stateUpdate.Resources
		g.step = stateUpdate.Step
		g.refreshTileSprites(len(stateUpdate.Map))
	}
}

func (g *GameView) refreshTileSprites(knownTilesNum int) {
	g.sprites = make([]TileSprite, knownTilesNum)
	index := 0
	for x := 0; x < int(g.config.Map.Width); x++ {
		for y := 0; y < int(g.config.Map.Height); y++ {
			if g.board[x][y].Landscape != types.Landscape_UNKNOWN {
				posX := tileMargin + (x * (tileWidth + tileMargin))
				posY := tileMargin + (y * (tileHeight + tileMargin))
				g.sprites[index] = TileSprite{
					x: float64(posX), 
					y: float64(posY),
				}
				if g.board[x][y].Settlement != types.Settlement_NONE {
					g.sprites[index].sprite = SpriteFromSettlement(g.board[x][y].Settlement)
					g.sprites[index].color = FactionToColorSprite(g.board[x][y].Faction)
					// TODO: also need to add the population
				} else {
					g.sprites[index].sprite = SpriteFromLandscape(g.board[x][y].Landscape)
					g.sprites[index].color = greySprite
				}
			}
		}
	}
}

func NewMapImage(config *types.MapConfig) *ebiten.Image {
	width := tileMargin + int(config.Width)*(tileWidth+tileMargin)
	height := tileMargin + int(config.Height)*(tileHeight+tileMargin)
	image, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	image.Fill(shadowColor)
	return image
}

func IndexToCoordinate(index, width uint32) (x, y int) {
	return int(index % width), int(index / width)
}

type TileSprite struct {
	x, y float64
	sprite *ebiten.Image
	color *ebiten.Image
	population uint32
}
