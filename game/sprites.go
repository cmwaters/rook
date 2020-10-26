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
	//landscapes and settlements
	capitalSprite,
	citySprite,
	farmSprite,
	forestSprite,
	lakeSprite,
	lumbermillSprite,
	mountainsSprite,
	plainsSprite,
	quarrySprite,
	rookSprite,
	townSprite *ebiten.Image
)

var (
	// colors for each faction
	redSprite,
	lightRedSprite,
	blueSprite,
	lightBlueSprite,
	greySprite,
	lightGreySprite *ebiten.Image
)

var colorSprites = []*ebiten.Image{redSprite, blueSprite}

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	farmSprite, _, err = util.NewImageFromFile(filepath.Join(pwd, "../../game/assets/farm.png"), ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}

	plainsSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = plainsSprite.Fill(greyColor)

	redSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = redSprite.Fill(redColor)
	lightRedSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = lightRedSprite.Fill(lightRedColor)
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
	number, err := strconv.Atoi(f.Moniker[len(f.Moniker) - 1:])
	if err != nil {
		panic(fmt.Errorf("converting faction moniker: %s to number: %w", f.Moniker, err))
	}
	return colorSprites[number % len(colorSprites)]
}