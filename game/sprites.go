package game

import (
	"github.com/hajimehoshi/ebiten"
	util "github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	//landscapes
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

func init() {
	var err error
	farmSprite, _, err = util.NewImageFromFile("/home/callum/go/src/github.com/cmwaters/rook/game/assets/farm.png", ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}

	plainsSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = plainsSprite.Fill(whiteColor)

	rookSprite, _ = ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	_ = rookSprite.Fill(redColor)
}
