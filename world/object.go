package world

type Object struct {
	PositionX float64
	PositionY float64
	Width     float64
	Height    float64
}

func (o Object) CheckCollision(another Object, direction string) bool {

	if direction == "down" {
		if between(o.PositionX, o.Width, another.PositionX, another.Width) &&
			between(o.PositionY, o.Height, another.PositionY, another.Height) {
			return true
		}
	}

	return false
}

func between(x float64, xw float64, ax float64, aw float64) bool {
	return (x <= ax+aw && x >= ax) || (x+xw <= ax+aw && x+xw >= ax)
}
