package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func Run() {
	if err := LoadEntities(); err != nil {
		log.Fatal("error loading sprites", err)
	}

	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle(GAME_TITLE)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
