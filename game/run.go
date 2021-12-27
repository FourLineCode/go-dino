package game

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var files embed.FS

func Run(f embed.FS) {
	files = f
	if err := LoadEntities(); err != nil {
		log.Fatal("error loading sprites", err)
	}

	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle(GAME_TITLE)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
