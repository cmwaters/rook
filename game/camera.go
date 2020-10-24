package game

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

const scrollSensitivityFactor = 4

type Camera struct {
	posX, posY, targetX, targetY float64
	pGain float64
	speed float64
}

func NewCamera(x, y, gain, speed float64) *Camera {
	return &Camera{x, y, x, y, gain, speed}
}

func (c *Camera) ParseMovementKeys() {
	diffX, diffY := 0.0, 0.0
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		diffY -= c.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		diffX -= c.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		diffY += c.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		diffX += c.speed
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

func (c *Camera) MoveTo(x, y int) {

}

func (c *Camera) Update() *ebiten.DrawImageOptions {
	diffX, diffY := (c.targetX - c.posX), (c.targetY - c.posY)
	c.posX += (diffX * c.pGain)
	c.posY += (diffY * c.pGain)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.posX, c.posY)
	return op
}