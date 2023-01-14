package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type State interface {
	Draw(screen *ebiten.Image, g *Game)
}

type Game struct {
	Count          float32
	State          State
	Keys           []ebiten.Key
	KeyLastPressed map[ebiten.Key]bool
}

func (g *Game) Update() error {
	g.KeyLastPressed = make(map[ebiten.Key]bool)
	for _, p := range g.Keys {
		g.KeyLastPressed[p] = true
	}
	g.Keys = inpututil.AppendPressedKeys(g.Keys[:0])
	g.Count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.State.Draw(screen, g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
