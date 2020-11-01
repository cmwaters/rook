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
	greenSprite,
	lightGreenSprite,
	yellowSprite,
	lightYellowSprite,
	purpleSprite,
	lightPurpleSprite,
	tealSprite,
	lightTealSprite,
	greySprite,
	lightGreySprite *ebiten.Image
	colors = []string{"red", "green", "yellow", "blue", "purple", "teal"}
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	capitalSprite, _, err = util.NewImageFromFile(filepath.Join(pwd, "../../game/assets/Capital.png"), ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	citySprite, _, err = util.NewImageFromFile(filepath.Join(pwd, "../../game/assets/City.png"), ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	farmSprite, _, err = util.NewImageFromFile(filepath.Join(pwd, "../../game/assets/Farm.png"), ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	lumbermillSprite, _, err = util.NewImageFromFile(filepath.Join(pwd, "../../game/assets/Lumbermill.png"), ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	quarrySprite, _, err = util.NewImageFromFile(filepath.Join(pwd, "../../game/assets/Quarry.png"), ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	rookSprite, _, err = util.NewImageFromFile(filepath.Join(pwd, "../../game/assets/Rook.png"), ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	townSprite, _, err = util.NewImageFromFile(filepath.Join(pwd, "../../game/assets/Town.png"), ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}

	forestSprite, _, err = util.NewImageFromFile(filepath.Join(pwd, "../../game/assets/Forest.png"), ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	lakeSprite, _, err = util.NewImageFromFile(filepath.Join(pwd, "../../game/assets/Lake.png"), ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	mountainsSprite, _, err = util.NewImageFromFile(filepath.Join(pwd, "../../game/assets/Mountains.png"), ebiten.FilterDefault)
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
	greenSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = greenSprite.Fill(greenColor)
	lightGreenSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = lightGreenSprite.Fill(lightGreenColor)
	yellowSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = yellowSprite.Fill(yellowColor)
	lightYellowSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = lightYellowSprite.Fill(lightYellowColor)
	purpleSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = purpleSprite.Fill(purpleColor)
	lightPurpleSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = lightPurpleSprite.Fill(lightPurpleColor)
	tealSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = tealSprite.Fill(tealColor)
	lightTealSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = lightTealSprite.Fill(lightTealColor)
	greySprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = greySprite.Fill(greyColor)
	lightGreySprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = lightGreySprite.Fill(lightGreyColor)

	// add all the colors to an array
	colorSprites = map[string]*ebiten.Image{
		"red":    redSprite,
		"blue":   blueSprite,
		"green":  greenSprite,
		"yellow": yellowSprite,
		"purple": purpleSprite,
		"teal":   tealSprite,
		"grey":   greySprite,
	}

	landSprites = map[types.Landscape]*ebiten.Image{
		types.Landscape_FOREST:    forestSprite,
		types.Landscape_LAKE:      lakeSprite,
		types.Landscape_MOUNTAINS: mountainsSprite,
		types.Landscape_PLAINS:    plainsSprite,
	}

	settlementSprites = map[types.Settlement]*ebiten.Image{
		types.Settlement_CAPITAL:    capitalSprite,
		types.Settlement_CITY:       citySprite,
		types.Settlement_FARM:       farmSprite,
		types.Settlement_LUMBERMILL: lumbermillSprite,
		types.Settlement_QUARRY:     quarrySprite,
		types.Settlement_ROOK:       rookSprite,
		types.Settlement_TOWN:       townSprite,
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

func FactionToColorSprite(f *types.Faction) *ebiten.Image {
	if f == nil {
		return colorSprites["grey"]
	}
	number, err := strconv.Atoi(f.Moniker[len(f.Moniker)-1:])
	if err != nil {
		panic(fmt.Sprintf("converting faction moniker: %s to number: %s", f.Moniker, err))
	}
	return colorSprites[colors[number%len(colors)]]
}

func toActivatedColor(color *ebiten.Image) *ebiten.Image {
	switch color {
	case redSprite:
		return lightRedSprite
	case blueSprite:
		return lightBlueSprite
	case greenSprite:
		return lightGreenSprite
	case yellowSprite:
		return lightYellowSprite
	case purpleSprite:
		return lightPurpleSprite
	case tealSprite:
		return lightTealSprite
	case greySprite:
		return lightGreySprite
	default:
		panic(fmt.Sprintf("%v is not a recognized color image in rook", color))
	}
}
