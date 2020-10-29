package game

import (
	"fmt"
	"strconv"

	e "github.com/cmwaters/rook/game/engine"
	"github.com/cmwaters/rook/x/rook/types"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
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
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.activateTileAt(ebiten.CursorPosition())
	}
	g.camera.ParseMovementKeys()
	return views[gameView], nil
}

func (g *GameView) Draw(screen *ebiten.Image) {
	_ = screen.Fill(shadowColor)
	if g.refresh {
		g.mapImage.Clear()
		for x := 0; x < int(g.config.Map.Width); x++ {
			for y := 0; y < int(g.config.Map.Height); y++ {
				if g.board[x][y].Landscape != types.Landscape_UNKNOWN {
					posX := tileMargin + (x * (tileWidth + tileMargin))
					posY := tileMargin + (y * (tileHeight + tileMargin))
					g.op.GeoM.Reset()
					g.op.GeoM.Translate(float64(posX), float64(posY))
					if g.board[x][y].Settlement == types.Settlement_CAPITAL && g.step == 0 {
						g.camera.MoveTo(posX, posY)
						g.active = position{x: x, y: y}
					}
					if g.active.x == x && g.active.y == y {
						_ = g.mapImage.DrawImage(toActivatedColor(colorSprites[FactionToColorSprite(g.board[x][y].Faction)]), g.op)
					} else { 
						_ = g.mapImage.DrawImage(colorSprites[FactionToColorSprite(g.board[x][y].Faction)], g.op)
					}
					if g.board[x][y].Settlement != types.Settlement_NONE {
						_ = g.mapImage.DrawImage(settlementSprites[g.board[x][y].Settlement], g.op)
					} else {
						_ = g.mapImage.DrawImage(landSprites[g.board[x][y].Landscape], g.op)
					}
					// TODO: A more dynamic text positioning algorithm
					if g.board[x][y].Population != 0 {
						text.Draw(g.mapImage, strconv.Itoa(int(g.board[x][y].Population)), mFont, posX + (tileWidth/2 - 6), posY + (tileHeight/2 + 6), whiteColor)
					}
				}
			}
		}
		g.refresh = false
	}
	_ = screen.DrawImage(g.mapImage, g.camera.Update())
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f\nCamera position %.2f, %.2f\nActive tile %d, %d", 
		ebiten.CurrentTPS(), -g.camera.posX, -g.camera.posY, g.active.x, g.active.y),
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
		g.refresh = true
		// g.refreshTileSprites(len(stateUpdate.Map))
	}
}

// DEPRECATE
// func (g *GameView) refreshTileSprites(knownTilesNum int) {
// 	g.sprites = make([]TileSprite, knownTilesNum)
// 	index := 0
// 	for x := 0; x < int(g.config.Map.Width); x++ {
// 		for y := 0; y < int(g.config.Map.Height); y++ {
// 			if g.board[x][y].Landscape != types.Landscape_UNKNOWN {
// 				posX := tileMargin + (x * (tileWidth + tileMargin))
// 				posY := tileMargin + (y * (tileHeight + tileMargin))
// 				g.sprites[index] = TileSprite{
// 					x:    x,
// 					y:    y,
// 					posX: float64(posX),
// 					posY: float64(posY),
// 				}
// 				if g.board[x][y].Settlement != types.Settlement_NONE {
// 					// if this is the start of the game we should move the camera
// 					// to the players capital
// 					if g.board[x][y].Settlement == types.Settlement_CAPITAL && g.step == 0 {
// 						fmt.Printf("\nTILE POS, X: %d, Y: %d", posX, posY)
// 						g.camera.MoveTo(posX, posY)
// 						g.active = position{x: x, y: y, index: index}
// 					}

// 					g.sprites[index].sprite = SpriteFromSettlement(g.board[x][y].Settlement)
// 					g.sprites[index].color = FactionToColorSprite(g.board[x][y].Faction)
// 					g.sprites[index].population = g.board[x][y].Population
// 				} else {
// 					g.sprites[index].sprite = SpriteFromLandscape(g.board[x][y].Landscape)
// 					if g.board[x][y].Population != 0 {
// 						g.sprites[index].color = FactionToColorSprite(g.board[x][y].Faction)
// 						g.sprites[index].population = g.board[x][y].Population
// 					} else {
// 						g.sprites[index].color = greySprite
// 					}
// 				}
// 				if g.active.x == x && g.active.y == y {
// 					g.sprites[index].color = toActivatedColor(FactionToColorSprite(g.board[x][y].Faction))
// 				}
// 				if g.sprites[index].population != 0 {
// 					text.Draw(g.sprites[index].color, strconv.Itoa(int(g.board[x][y].Population)), mFont, (tileWidth/2 - 6), (tileHeight/2 + 6), greyColor)
// 				}
// 				index++
// 			}
// 		}
// 	}
// 	g.refresh = true
// }

func (g *GameView) activateTileAt(clickX, clickY int) {
	clickX, clickY = clickX-int(g.camera.posX), clickY-int(g.camera.posY)
	fmt.Printf("\nadjusted click at x: %d, y: %d", clickX, clickY)
	for x := 0; x < int(g.config.Map.Width); x++ {
		for y := 0; y < int(g.config.Map.Height); y++ {
			if g.board[x][y].Landscape != types.Landscape_UNKNOWN {
				posX := tileMargin + (x * (tileWidth + tileMargin))
				posY := tileMargin + (y * (tileHeight + tileMargin))
				if clickX > posX && clickX < posX+tileWidth && clickY > posY && clickY < posY+tileHeight {
					fmt.Printf("\ndeactivating sprite at x %d y %d", g.active.x, g.active.y)
					prior := g.board[g.active.x][g.active.y]
					priorPosX := tileMargin + (g.active.x * (tileWidth + tileMargin))
					priorPosY := tileMargin + (g.active.y * (tileHeight + tileMargin))
					g.op.GeoM.Reset()
					g.op.GeoM.Translate(float64(priorPosX), float64(priorPosY))
					_ = g.mapImage.DrawImage(colorSprites[FactionToColorSprite(prior.Faction)], g.op)
					if prior.Settlement != types.Settlement_NONE {
						_ = g.mapImage.DrawImage(settlementSprites[prior.Settlement], g.op)
					} else {
						_ = g.mapImage.DrawImage(landSprites[prior.Landscape], g.op)
					}
					if prior.Population != 0 {
						text.Draw(g.mapImage, strconv.Itoa(int(prior.Population)), mFont, priorPosX + (tileWidth/2 - 6), priorPosY + (tileHeight/2 + 6), whiteColor)
					}
					
					// t := &g.board[g.active.x][g.active.y]
					g.active = position{x: x, y: y}
					g.op.GeoM.Reset()
					g.op.GeoM.Translate(float64(posX), float64(posY))
					// if toActivatedColor(colorSprites[FactionToColorSprite(g.board[x][y].Faction)]) == lightGreySprite {
					// 	panic("Hello")
					// }
					// _ = g.mapImage.DrawImage(lightGreySprite, g.op)
					fmt.Printf("\nactivating sprite at x %d y %d", x, y)
					_ = g.mapImage.DrawImage(toActivatedColor(colorSprites[FactionToColorSprite(g.board[x][y].Faction)]), g.op)
					if g.board[x][y].Settlement != types.Settlement_NONE {
						_ = g.mapImage.DrawImage(settlementSprites[g.board[x][y].Settlement], g.op)
					} else {
						_ = g.mapImage.DrawImage(landSprites[g.board[x][y].Landscape], g.op)
					}
					if g.board[x][y].Population != 0 {
						text.Draw(g.mapImage, strconv.Itoa(int(g.board[x][y].Population)), mFont, posX + (tileWidth/2 - 6), posY + (tileHeight/2 + 6), whiteColor)
					}
					// _ = g.mapImage.DrawImage(lightGreySprite, g.op)
					// g.sprites[g.active.index] = TileSprite{
					// 	x: t.x,
					// 	y: t.y,
					// 	posX: t.posX,
					// 	posY: t.posY,
					// 	color: toDeactivatedColor(t.color),
					// 	sprite: t.sprite,
					// 	population: t.population,
					// }
					// if t.population != 0 {
					// 	text.Draw(g.sprites[g.active.index].color, strconv.Itoa(int(t.population)), mFont, (tileWidth/2 - 6), (tileHeight/2 + 6), greyColor)
					// }
					// g.op.GeoM.Reset()
					// g.op.GeoM.Translate(t.posX, t.posY)
					// _ = g.mapImage.DrawImage(t.color, g.op)
					// _ = g.mapImage.DrawImage(t.sprite, g.op)
					// fmt.Printf("\nactivating sprite at x %d y %d", sprite.x, sprite.y)
					// g.active = position{sprite.x, sprite.y, idx}
					// if sprite.color != greySprite {
					// 	panic("why are you not grey")
					// } 
					// g.sprites[idx] = TileSprite{
					// 	x: sprite.x,
					// 	y: sprite.y,
					// 	posX: sprite.posX,
					// 	posY: sprite.posY,
					// 	color: toActivatedColor(sprite.color),
					// 	sprite: sprite.sprite,
					// 	population: sprite.population,
					// }
					// if sprite.population != 0 {
					// 	text.Draw(g.sprites[idx].color, strconv.Itoa(int(sprite.population)), mFont, (tileWidth/2 - 6), (tileHeight/2 + 6), greyColor)
					// }
					// g.op.GeoM.Reset()
					// g.op.GeoM.Translate(sprite.posX, sprite.posY)
					// _ = g.mapImage.DrawImage(sprite.color, g.op)
					// _ = g.mapImage.DrawImage(sprite.sprite, g.op)
					// g.camera.MoveTo(int(sprite.posX), int(sprite.posY))
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
	x, y       int
	posX, posY float64
	sprite     string //reference to the image
	color      string //reference to the image
	population uint32
}

type position struct {
	x, y int
}
