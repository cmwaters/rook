package game

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/cmwaters/rook/x/rook/types"
	"github.com/hajimehoshi/ebiten"
	util "github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	// settlements
	settlementSprites map[types.Settlement]*ebiten.Image
	capitalSprite,
	citySprite,
	farmSprite,
	lumbermillSprite,
	quarrySprite,
	rookSprite,
	townSprite *ebiten.Image
	
	//landscapes
	landSprites map[types.Landscape]*ebiten.Image
	forestSprite,
	lakeSprite,
	mountainsSprite,
	plainsSprite *ebiten.Image

	// faction colours
	colorSprites map[string]*ebiten.Image
	redSprite,
	lightRedSprite,
	blueSprite,
	lightBlueSprite,
	greySprite,
	lightGreySprite *ebiten.Image
	colors = []string{"red", "blue"}
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	farmSprite, _, err = util.NewImageFromFile(filepath.Join(pwd, "../../game/assets/Farm-01.png"), ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	capitalSprite, _, err = util.NewImageFromFile(filepath.Join(pwd, "../../game/assets/Capital-01.png"), ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}

	plainsSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)

	redSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = redSprite.Fill(redColor)
	lightRedSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = lightRedSprite.Fill(lightRedColor)
	blueSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = blueSprite.Fill(blueColor)
	lightBlueSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = lightBlueSprite.Fill(lightBlueColor)
	greySprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = greySprite.Fill(greyColor)
	lightGreySprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = lightGreySprite.Fill(lightGreyColor)
	// add all the colors to an array
	colorSprites = map[string]*ebiten.Image{
		"red": redSprite, 
		"blue": blueSprite,
		"grey": greySprite,
	}
	
	landSprites = map[types.Landscape]*ebiten.Image{
		types.Landscape_FOREST: forestSprite,
		types.Landscape_LAKE: lakeSprite,
		types.Landscape_MOUNTAINS: mountainsSprite,
		types.Landscape_PLAINS: plainsSprite,
	}
	
	settlementSprites = map[types.Settlement]*ebiten.Image{
	types.Settlement_CAPITAL: capitalSprite,
	types.Settlement_CITY: citySprite,
	types.Settlement_FARM: farmSprite,
	types.Settlement_LUMBERMILL: lumbermillSprite,
	types.Settlement_QUARRY: quarrySprite,
	types.Settlement_ROOK: rookSprite,
	types.Settlement_TOWN: townSprite,
	}
	
}

func SpriteFromLandscape(l types.Landscape) *ebiten.Image {
	switch l {
	case types.Landscape_FOREST:
		return forestSprite
	case types.Landscape_MOUNTAINS:
		return mountainsSprite
	case types.Landscape_LAKE:
		return lakeSprite
	case types.Landscape_PLAINS:
		return plainsSprite
	default:
		return &ebiten.Image{}
	}
}

func SpriteFromSettlement(s types.Settlement) *ebiten.Image {
	switch s {
	case types.Settlement_CAPITAL:
		return capitalSprite
	case types.Settlement_CITY:
		return citySprite
	case types.Settlement_FARM:
		return farmSprite
	case types.Settlement_LUMBERMILL:
		return lumbermillSprite
	case types.Settlement_QUARRY:
		return quarrySprite
	case types.Settlement_ROOK:
		return rookSprite
	case types.Settlement_TOWN:
		return townSprite
	default:
		return &ebiten.Image{}
	}
}

func FactionToColorSprite(f *types.Faction) string {
	if f == nil {
		return "grey"
	}
	number, err := strconv.Atoi(f.Moniker[len(f.Moniker)-1:])
	if err != nil {
		panic(fmt.Sprintf("converting faction moniker: %s to number: %s", f.Moniker, err))
	}
	return colors[number%len(colors)]
}

func toActivatedColor(color *ebiten.Image) *ebiten.Image {
	switch color {
	case redSprite:
		return lightRedSprite
	case blueSprite:
		return lightBlueSprite
	case greySprite:
		return lightGreySprite
	default:
		panic(fmt.Sprintf("%v is not a recognized color image in rook", color))
	}
}

func toDeactivatedColor(color *ebiten.Image) *ebiten.Image {
	switch color {
	case lightRedSprite:
		return redSprite
	case lightBlueSprite:
		return blueSprite
	case lightGreySprite:
		return greySprite
	default:
		panic(fmt.Sprintf("%v is not a recognized color image in rook", color))
	}
}
