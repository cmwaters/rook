package game

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

const scrollSensitivityFactor = 4

type Camera struct {
	posX, posY, targetX, targetY float64
	pGain                        float64
	speed                        float64
	width, height                float64
}

func NewCamera(x, y, gain, speed, width, height float64) *Camera {
	return &Camera{x, y, x, y, gain, speed, width, height}
}

func (c *Camera) ParseMovementKeys() {
	diffX, diffY := 0.0, 0.0
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		diffY += c.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		diffX += c.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		diffY -= c.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		diffX -= c.speed
	}
	// adjust if diagonal
	if diffX != 0 && diffY != 0 {
		diffX /= math.Sqrt(2)
		diffY /= math.Sqrt(2)
	}
	// check scrolling
	if diffX == 0 && diffY == 0 {
		x, y := ebiten.Wheel()
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			x, y = y, x // swap them
		}
		diffX += (x * c.speed * scrollSensitivityFactor)
		diffY += (y * c.speed * scrollSensitivityFactor)
	}
	c.targetX += diffX
	c.targetY += diffY
}

// moves the camera to an absolute point in the map. Used
// for centering on tiles
func (c *Camera) MoveTo(x, y int) {
	// the second part of the assignment center the square in
	// the middle of the screen as opposed to the top left
	c.targetX = -float64(x) + (c.width - float64(tileWidth))/2
	c.targetY = -float64(y) + (c.height - float64(tileHeight))/2
}

func (c *Camera) Update() *ebiten.DrawImageOptions {
	diffX, diffY := (c.targetX - c.posX), (c.targetY - c.posY)
	c.posX += (diffX * c.pGain)
	c.posY += (diffY * c.pGain)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.posX, c.posY)
	return op
}
