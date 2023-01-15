package world

import (
	"image/color"
)

type Player struct {
	Object
	Color        color.Color
	VectorSpeedY float64
	VectorSpeedX float64
	ForceActions []Action
}

func GetNewPlayer() Player {
	playerColor := color.RGBA64{128 * 100, 255 * 100, 100 * 100, 65535}

	actions := []Action{
		{"jump", 0, 10},
	}

	collisionObject := Object{
		10, 10, 20, 40,
	}

	return Player{collisionObject, playerColor, 1, 0, actions}
}

func (p *Player) updateVelocity(objects *[]Object) {

	playerOnGround, _ := p.isPlayerOnGround(objects)

	if !playerOnGround {
		p.VectorSpeedY -= 1
	} else {
		p.VectorSpeedX /= 1.5
		p.VectorSpeedY = 0
	}

	for _, action := range p.ForceActions {
		if action.Name == "jump" {
			if playerOnGround {
				p.VectorSpeedY += action.VecY
				p.VectorSpeedX += action.VecX
			}
		}
	}

	if p.VectorSpeedY < -16 {
		p.VectorSpeedY = -16
	}
	p.ForceActions = make([]Action, 2)
}

func (p *Player) addAction(a Action) {
	p.ForceActions = append(p.ForceActions, a)
}

func (p *Player) updatePosition(objects *[]Object) {
	p.PositionY += p.VectorSpeedY
	p.PositionX += p.VectorSpeedX

	plOnG, object := p.isPlayerOnGround(objects)

	if p.VectorSpeedY < 0 && plOnG {
		p.PositionY = object.PositionY + object.Height
		p.VectorSpeedY = 0
	}
}

func (p *Player) isPlayerOnGround(objects *[]Object) (bool, *Object) {
	for _, object := range *objects {
		if p.CheckCollision(object, "down") {
			return true, &object
		}
	}

	return false, nil
}
