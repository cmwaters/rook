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
	moniker   string
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
		moniker:   "player3",
	}
	statec := make(chan *types.PartialState)
	go g.engine.Init(statec)
	go g.ApplyStateTransitions(statec)
	return g
}

func (g *GameView) Update(views map[string]View) (View, error) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.handleClick(ebiten.CursorPosition())
	}
	g.parseMoveKeys() // arrow keys && WASD for move actions and navigation
	g.parseBuildKeys() // build keys for building settlements
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
						_ = g.mapImage.DrawImage(toActivatedColor(FactionToColorSprite(g.board[x][y].Faction)), g.op)
					} else { 
						_ = g.mapImage.DrawImage(FactionToColorSprite(g.board[x][y].Faction), g.op)
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
		fmt.Sprintf("TPS: %0.2f\nCamera position %.2f, %.2f\nActive tile %d, %d\nResources Food: %d, Stone: %d, Wood: %d", 
		ebiten.CurrentTPS(), -g.camera.posX, -g.camera.posY, g.active.x, g.active.y, 
		g.resources.Food, g.resources.Stone, g.resources.Wood),
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
	}
}

func (g *GameView) handleClick(clickX, clickY int) {
	clickX, clickY = clickX-int(g.camera.posX), clickY-int(g.camera.posY)
	fmt.Printf("\nadjusted click at x: %d, y: %d", clickX, clickY)
	for x := 0; x < int(g.config.Map.Width); x++ {
		for y := 0; y < int(g.config.Map.Height); y++ {
			if g.board[x][y].Landscape != types.Landscape_UNKNOWN {
				posX := tileMargin + (x * (tileWidth + tileMargin))
				posY := tileMargin + (y * (tileHeight + tileMargin))
				if clickX > posX && clickX < posX+tileWidth && clickY > posY && clickY < posY+tileHeight {
					quantity, direction := g.attemptToMovePopulation(x, y)
					if quantity != 0 {
						fmt.Printf("Sending move request %v\n", direction)
						g.engine.Move(quantity, g.active.x, g.active.y, direction)
					}
					g.deactivateCurrentTile()
					g.activateTile(x, y, posX, posY)
				}
			}
		}
	}
}

func (g *GameView) deactivateCurrentTile() {
	fmt.Printf("\ndeactivating sprite at x %d y %d", g.active.x, g.active.y)
	prior := g.board[g.active.x][g.active.y]
	priorPosX := tileMargin + (g.active.x * (tileWidth + tileMargin))
	priorPosY := tileMargin + (g.active.y * (tileHeight + tileMargin))
	g.op.GeoM.Reset()
	g.op.GeoM.Translate(float64(priorPosX), float64(priorPosY))
	_ = g.mapImage.DrawImage(FactionToColorSprite(prior.Faction), g.op)
	if prior.Settlement != types.Settlement_NONE {
		_ = g.mapImage.DrawImage(settlementSprites[prior.Settlement], g.op)
	} else {
		_ = g.mapImage.DrawImage(landSprites[prior.Landscape], g.op)
	}
	if prior.Population != 0 {
		text.Draw(g.mapImage, strconv.Itoa(int(prior.Population)), mFont, priorPosX + (tileWidth/2 - 6), priorPosY + (tileHeight/2 + 6), whiteColor)
	}
}

func (g *GameView) activateTile(x, y, posX, posY int) {
	g.active = position{x: x, y: y}
	g.op.GeoM.Reset()
	g.op.GeoM.Translate(float64(posX), float64(posY))
	fmt.Printf("\nactivating sprite at x %d y %d", x, y)
	_ = g.mapImage.DrawImage(toActivatedColor(FactionToColorSprite(g.board[x][y].Faction)), g.op)
	if g.board[x][y].Settlement != types.Settlement_NONE {
		_ = g.mapImage.DrawImage(settlementSprites[g.board[x][y].Settlement], g.op)
	} else {
		_ = g.mapImage.DrawImage(landSprites[g.board[x][y].Landscape], g.op)
	}
	if g.board[x][y].Population != 0 {
		text.Draw(g.mapImage, strconv.Itoa(int(g.board[x][y].Population)), mFont, posX + (tileWidth/2 - 6), posY + (tileHeight/2 + 6), whiteColor)
	}
}

func (g *GameView) parseMoveKeys() {
	x, y := g.active.x, g.active.y
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if g.active.y == 0 {
			y = int(g.config.Map.Height)-1
		} else {
			y = g.active.y-1
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if g.active.y == int(g.config.Map.Height)-1 {
			y = 0
		} else {
			y = g.active.y+1
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if g.active.x == 0 {
			x = int(g.config.Map.Width) - 1
		} else {
			x = g.active.x-1
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.active.x == int(g.config.Map.Width) {
			x = 0
		} else {
			x = g.active.x+1
		}
	}
	if x != g.active.x || y != g.active.y { // we have moved
		if g.board[x][y].Landscape != types.Landscape_UNKNOWN {
			posX := tileMargin + (x * (tileWidth + tileMargin))
			posY := tileMargin + (y * (tileHeight + tileMargin))
			if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyLeft) || 
			inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyDown) { // action keys
				if (g.board[g.active.x][g.active.y].Population > 0 && g.board[g.active.x][g.active.y].Faction.Moniker == g.moniker) {
					quantity, direction := g.attemptToMovePopulation(x, y)
					if quantity != 0 {
						fmt.Printf("Sending move request %v\n", direction)
						g.engine.Move(quantity, g.active.x, g.active.y, direction)
					}
				}
			}
			g.deactivateCurrentTile()
			g.activateTile(x, y, posX, posY)
			
			if g.nearingBoundary(posX, posY) {
				g.camera.MoveTo(posX, posY)
			}
		}
	}
}

func (g *GameView) parseBuildKeys() {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) { // build a farm
		if types.ConstructionResources(g.config.Construction, types.Settlement_FARM).Less(g.resources) {
			g.engine.Build(types.Settlement_FARM, g.active.x, g.active.y)
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyQ) { // build a quarry
		if types.ConstructionResources(g.config.Construction, types.Settlement_QUARRY).Less(g.resources) {
			g.engine.Build(types.Settlement_QUARRY, g.active.x, g.active.y)
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyT) { // build a town
		if types.ConstructionResources(g.config.Construction, types.Settlement_TOWN).Less(g.resources) {
			g.engine.Build(types.Settlement_TOWN, g.active.x, g.active.y)
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyC) { // build a city/capital
		if types.ConstructionResources(g.config.Construction, types.Settlement_CITY).Less(g.resources) {
			g.engine.Build(types.Settlement_CITY, g.active.x, g.active.y)
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyR) { // build a rook
		if types.ConstructionResources(g.config.Construction, types.Settlement_ROOK).Less(g.resources) {
			g.engine.Build(types.Settlement_ROOK, g.active.x, g.active.y)
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyL) { // build a lumbermill
		if types.ConstructionResources(g.config.Construction, types.Settlement_LUMBERMILL).Less(g.resources) {
			g.engine.Build(types.Settlement_LUMBERMILL, g.active.x, g.active.y)
		}
	}
}

func (g *GameView) attemptToMovePopulation(targetX, targetY int) (quantity uint32, direction types.Direction) {
	quantity = g.board[g.active.x][g.active.y].Population
	if g.board[g.active.x][g.active.y].Settlement != types.Settlement_NONE {
		// leave one population behind to occupy the settlement
		quantity--
	}
	if quantity == 0 {
		return 0, types.Direction_NOWHERE // we have no population to move
	}

	if targetX == g.active.x { // attempt was on the same vertical plane
		// check if moving up
		if g.active.y == 0  && targetY == int(g.config.Map.Height) - 1 || targetY == g.active.y - 1 {
			direction = types.Direction_UP
		} else if g.active.y == int(g.config.Map.Height) - 1 && targetY == 0 || targetY == g.active.y + 1 {
			// moving down
			direction = types.Direction_DOWN
		} else {
			return 0, types.Direction_NOWHERE
		}
	} else if targetY == g.active.y { // attempt was on the same horizontal plane
		// check if moving to the right
		if g.active.x == int(g.config.Map.Width) - 1 && targetX == 0 || targetX == g.active.x + 1 {
			direction = types.Direction_RIGHT
		} else if g.active.x == 0 && targetX == int(g.config.Map.Width) - 1 || targetX == g.active.x - 1 {
			// moving to the left
			direction = types.Direction_LEFT
		} else {
			return 0, types.Direction_NOWHERE
		}
	} else { // on neither planes
		return 0, types.Direction_NOWHERE
	}
	
	return
}

// perhaps we should be adding this to camera
func (g *GameView) nearingBoundary(x, y int) bool {
	adjX, adjY := float64(x) + g.camera.posX, float64(y) + g.camera.posY
	return adjX < viewBoundaryMargin || adjX > g.camera.width - viewBoundaryMargin || 
		adjY < viewBoundaryMargin || adjY > g.camera.height - viewBoundaryMargin
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
