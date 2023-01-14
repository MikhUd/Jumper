package main

import (
	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"os"
	"test/utils"
	"test/world"
)

type MenuState struct {
	currentChose int
}

func (state *MenuState) Draw(screen *ebiten.Image, g *utils.Game) {
	state.HandleEvent(g)
	img, _, _ := ebitenutil.NewImageFromFile("images/pngegg.png")
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(200, 0)
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(img, op)

	//posX, posY := ebiten.CursorPosition()
	switch state.currentChose {
	case 1:
		ebitenutil.DrawRect(screen, 35, 50, 80, 15, color.White)
	case 2:
		ebitenutil.DrawRect(screen, 35, 70, 80, 15, color.White)
	case 3:
		ebitenutil.DrawRect(screen, 35, 90, 80, 15, color.White)
	}
	ebitenutil.DebugPrintAt(screen, "Play", 40, 50)
	ebitenutil.DebugPrintAt(screen, "Credentials", 40, 70)
	ebitenutil.DebugPrintAt(screen, "Exit", 40, 90)
}

func (state *MenuState) HandleEvent(g *utils.Game) {
	for _, p := range g.Keys {
		if g.KeyLastPressed[p] {
			break
		}

		switch p.String() {
		case "ArrowDown":
			state.currentChose++
		case "ArrowUp":
			state.currentChose--
		case "Enter":
			state.DoAction(g)
		}

	}

	if state.currentChose == 0 {
		state.currentChose = 3
	}

	if state.currentChose == 4 {
		state.currentChose = 1
	}
}

func (state *MenuState) DoAction(g *utils.Game) {
	switch state.currentChose {
	case 3:
		os.Exit(0)
	case 1:
		ws, _, er := websocket.DefaultDialer.Dial("ws://91.149.232.36:6554", nil)

		if er != nil {
			break
		}

		state := &world.WorldState{
			MainHero:      world.GetNewPlayer(),
			OnlinePlayers: []world.OnlinePlayer{},
			Con:           ws,
		}
		g.State = state
		go state.HandleNetwork(g)
	}
}
