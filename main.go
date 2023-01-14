package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"test/utils"
)

func main() {
	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("Hell wanker")

	game := &utils.Game{State: &MenuState{1}}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
