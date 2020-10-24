package types


func (p *Position) Sum(p2 *Position) *Position {
	return &Position{X: p.X + p2.X, Y: p.Y + p2.Y}
}

func (p *Position) Wrap(c *MapConfig) {
	p.X = p.X % c.Width
	p.Y = p.Y % c.Height
}

func (p *Position) Index(c *MapConfig) uint32 {
	return p.Y * c.Width + p.X
}