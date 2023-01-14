package world

import (
	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"test/utils"
	"time"
)

type WorldState struct {
	MainHero      Player
	OnlinePlayers []OnlinePlayer
	Con           *websocket.Conn
}

var secretNum int = -1

func (state *WorldState) Draw(screen *ebiten.Image, g *utils.Game) {
	state.HandleEvent(g)
	backGround := color.RGBA64{128 * 170, 166 * 170, 255 * 170, 65535}
	ground := color.RGBA64{128 * 200, 255 * 200, 100 * 200, 65535}

	ebitenutil.DrawRect(screen, 0, 0, 320, 240, backGround)
	ebitenutil.DrawRect(screen, 0, 200, 320, 240, ground)

	ebitenutil.DrawRect(screen, state.MainHero.PositionX, state.MainHero.PositionY, 20, 40, state.MainHero.Color)

	for _, onlinePlayer := range state.OnlinePlayers {
		if onlinePlayer.PositionY != 0 {
			ebitenutil.DrawRect(screen, onlinePlayer.PositionX, onlinePlayer.PositionY, 20, 40, state.MainHero.Color)
		}
	}
}

func (state *WorldState) HandleEvent(g *utils.Game) {

	for _, p := range g.Keys {
		if g.KeyLastPressed[p] && p.String() == "Space" {
			break
		}

		switch p.String() {
		case "Space":
			if state.MainHero.isPlayerOnGround() {
				state.MainHero.PositionY = 159
				state.MainHero.VectorSpeedY -= 10
			}
		}
	}

	state.MainHero.updateVelocity()

	for _, p := range g.Keys {
		switch p.String() {
		case "ArrowLeft":
			state.MainHero.VectorSpeedX = -5
		case "ArrowRight":
			state.MainHero.VectorSpeedX = +5
		}

	}

	state.MainHero.PositionX += state.MainHero.VectorSpeedX

	if state.MainHero.PositionY+state.MainHero.VectorSpeedY > 160 {
		state.MainHero.PositionY = 160
	} else {
		state.MainHero.PositionY += state.MainHero.VectorSpeedY
	}

}

func (state *WorldState) HandleNetwork(g *utils.Game) {
	var mainHero OnlinePlayer
	state.Con.WriteJSON(OnlinePlayer{state.MainHero.PositionX, state.MainHero.PositionY, 0})
	state.Con.ReadJSON(&mainHero)
	secretNum = mainHero.N
	for true {
		state.Con.WriteJSON(OnlinePlayer{state.MainHero.PositionX, state.MainHero.PositionY, mainHero.N})
		state.Con.ReadJSON(&state.OnlinePlayers)
		time.Sleep(time.Millisecond * 30)
	}

}
