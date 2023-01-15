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
	MainHero         Player
	OnlinePlayers    []OnlinePlayer
	Con              *websocket.Conn
	collisionObjects []Object
}

func GenerateWorldState(ws *websocket.Conn) *WorldState {
	return &WorldState{
		MainHero:      GetNewPlayer(),
		OnlinePlayers: []OnlinePlayer{},
		Con:           ws,
		collisionObjects: []Object{
			{-1000, 0, 2320, 16},
			{170, 50, 100, 16},
			{270, 140, 100, 16},
		},
	}
}

func (state *WorldState) Draw(screen *ebiten.Image, g *utils.Game) {
	state.HandleEvent(g)
	backGround := color.RGBA64{128 * 170, 166 * 170, 255 * 170, 65535}
	ebitenutil.DrawRect(screen, 0, 0, 320, 240, backGround)

	objectColor := color.RGBA64{128 * 200, 255 * 200, 100 * 200, 65535}
	for _, object := range state.collisionObjects {
		state.drawTransferCoords(&object, screen, objectColor)
	}

	state.drawTransferCoords(&state.MainHero.Object, screen, state.MainHero.Color)

	for _, onlinePlayer := range state.OnlinePlayers {
		if onlinePlayer.PositionY != 0 {
			state.drawTransferCoords(&Object{
				onlinePlayer.PositionX,
				onlinePlayer.PositionY,
				20,
				40,
			}, screen, state.MainHero.Color)

		}
	}
}

func (w *WorldState) drawTransferCoords(object *Object, screen *ebiten.Image, colorToPaint color.Color) {
	ebitenutil.DrawRect(screen,
		object.PositionX-w.MainHero.PositionX+160,
		240-object.PositionY-object.Height,
		object.Width,
		object.Height,
		colorToPaint,
	)
}

func (state *WorldState) HandleEvent(g *utils.Game) {

	for _, p := range g.Keys {
		switch p.String() {
		case "Space":
			state.MainHero.addAction(Action{"jump", 0, 10})
		case "ArrowLeft":
			state.MainHero.VectorSpeedX = -5
		case "ArrowRight":
			state.MainHero.VectorSpeedX = +5
		}

	}

	state.MainHero.updateVelocity(&state.collisionObjects)
	state.MainHero.updatePosition(&state.collisionObjects)
}

func (state *WorldState) HandleNetwork(g *utils.Game) {
	var mainHero OnlinePlayer
	state.Con.WriteJSON(OnlinePlayer{state.MainHero.PositionX, state.MainHero.PositionY, 0})
	state.Con.ReadJSON(&mainHero)

	for true {
		state.Con.WriteJSON(OnlinePlayer{state.MainHero.PositionX, state.MainHero.PositionY, mainHero.N})
		state.Con.ReadJSON(&state.OnlinePlayers)
		time.Sleep(time.Millisecond * 30)
	}

}
