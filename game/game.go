package game

import (
	"fmt"
	"strconv"

	e "github.com/cmwaters/rook/game/engine"
	"github.com/cmwaters/rook/x/rook/types"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
)

var _ View = &GameView{}

var (
	tileWidth  = 64
	tileHeight = 64
	tileMargin = 2
)

type GameView struct {
	mapImage  *ebiten.Image
	camera    *Camera
	resources *types.ResourceSet
	board     [][]types.Tile // this represents the players known map
	config    *types.GameConfig
	op        *ebiten.DrawImageOptions
	engine    e.GameEngine
	step      uint32
	active    position
	sprites   []TileSprite
	refresh   bool
}

func NewLocalGameView(config *types.GameConfig, bots int) *GameView {
	fmt.Println("Starting local game")
	g := &GameView{
		mapImage:  NewMapImage(config.Map),
		camera:    NewCamera(0, 0, 0.3, 5, defaultScreenWidth, defaultScreenHeight),
		config:    config,
		resources: types.NewResourceSet(config.Initial.Resources),
		board:     types.NewEmptyBoard(config.Map),
		op:        &ebiten.DrawImageOptions{},
		engine:    e.NewLocalGameEngine(config, bots),
		step:      0,
	}
	statec := make(chan *types.PartialState)
	go g.engine.Init(statec)
	go g.ApplyStateTransitions(statec)
	return g
}

func (g *GameView) Update(views map[string]View) (View, error) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.activateTileAt(ebiten.CursorPosition())
	}
	g.camera.ParseMovementKeys()
	return views[gameView], nil
}

func (g *GameView) Draw(screen *ebiten.Image) {
	_ = screen.Fill(shadowColor)
	if g.refresh {
		for _, ts := range g.sprites {
			g.op.GeoM.Reset()
			g.op.GeoM.Translate(ts.posX, ts.posY)
			_ = g.mapImage.DrawImage(ts.color, g.op)
			_ = g.mapImage.DrawImage(ts.sprite, g.op)
			if ts.population != 0 {
				// TODO: write number over the top
			}
		}
		g.refresh = false
	}
	_ = screen.DrawImage(g.mapImage, g.camera.Update())
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f\nCamera position %.2f, %.2f", ebiten.CurrentTPS(), g.camera.posX, g.camera.posY),
	)
}

func (g *GameView) ApplyStateTransitions(c chan *types.PartialState) {
	for stateUpdate := range c {
		fmt.Printf("update state to %d\n", stateUpdate.Step)
		// update map first
		for posIdx, tile := range stateUpdate.Map {
			x, y := IndexToCoordinate(posIdx, g.config.Map.Width)
			fmt.Printf("%v at x: %d and y: %d\n", tile, x, y)
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
					x:    x,
					y:    y,
					posX: float64(posX),
					posY: float64(posY),
				}
				if g.board[x][y].Settlement != types.Settlement_NONE {
					// if this is the start of the game we should move the camera
					// to the players capital
					if g.board[x][y].Settlement == types.Settlement_CAPITAL && g.step == 0 {
						fmt.Printf("\nTILE POS, X: %d, Y: %d", posX, posY)
						g.camera.MoveTo(posX, posY)
						g.active = position{x: x, y: y, index: index}
					}

					g.sprites[index].sprite = SpriteFromSettlement(g.board[x][y].Settlement)
					g.sprites[index].color = FactionToColorSprite(g.board[x][y].Faction)
					g.sprites[index].population = g.board[x][y].Population
				} else {
					g.sprites[index].sprite = SpriteFromLandscape(g.board[x][y].Landscape)
					if g.board[x][y].Population != 0 {
						g.sprites[index].color = FactionToColorSprite(g.board[x][y].Faction)
						g.sprites[index].population = g.board[x][y].Population
					} else {
						g.sprites[index].color = greySprite
					}
				}
				if g.active.x == x && g.active.y == y {
					g.sprites[index].color = toActivatedColor(FactionToColorSprite(g.board[x][y].Faction))
				}
				if g.sprites[index].population != 0 {
					text.Draw(g.sprites[index].color, strconv.Itoa(int(g.board[x][y].Population)), mFont, (tileWidth/2 - 6), (tileHeight/2 + 6), greyColor)
				}
				index++
			}
		}
	}
	g.refresh = true
}

func (g *GameView) activateTileAt(x, y int) {
	fmt.Printf("\nclicked at x: %d, y: %d", x, y)
	x, y = x-int(g.camera.posX), y-int(g.camera.posY)
	fmt.Printf("\nadjusted click at x: %d, y: %d", x, y)
	for idx, sprite := range g.sprites {
		if x > int(sprite.posX) && x < int(sprite.posX)+tileWidth && y < int(sprite.posY) && y > int(sprite.posY)-tileHeight {
			previouslyActiveTile := &g.sprites[g.active.index]
			previouslyActiveTile.color = toDeactivatedColor(previouslyActiveTile.color)
			if previouslyActiveTile.population != 0 {
				text.Draw(previouslyActiveTile.color, strconv.Itoa(int(previouslyActiveTile.population)), mFont, (tileWidth/2 - 6), (tileHeight/2 + 6), greyColor)
			}
			g.active = position{sprite.x, sprite.y, idx}
			sprite.color = toActivatedColor(sprite.color)
			if sprite.population != 0 {
				text.Draw(sprite.color, strconv.Itoa(int(sprite.population)), mFont, (tileWidth/2 - 6), (tileHeight/2 + 6), greyColor)
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
	x, y       int
	posX, posY float64
	sprite     *ebiten.Image
	color      *ebiten.Image
	population uint32
}

type position struct {
	x, y, index int
}
