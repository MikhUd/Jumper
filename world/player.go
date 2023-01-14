package world

import "image/color"

type Player struct {
	PositionX    float64
	PositionY    float64
	Color        color.Color
	VectorSpeedY float64
	VectorSpeedX float64
}

func GetNewPlayer() Player {
	playerColor := color.RGBA64{128 * 100, 255 * 100, 100 * 100, 65535}

	return Player{150, 10, playerColor, 1, 0}
}

func (p *Player) updateVelocity() {
	if !p.isPlayerOnGround() {
		p.VectorSpeedY += 1
	} else {
		p.VectorSpeedX /= 2
		p.VectorSpeedY = 0
	}
}

func (p *Player) isPlayerOnGround() bool {
	if p.PositionY == 160 {
		return true
	}

	return false
}
