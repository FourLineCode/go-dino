package game

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/bitmapfont"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type GameState int

const (
	GameStateRunning GameState = iota
	GameStateLost    GameState = iota
)

type Game struct {
	Score int
	State GameState
	keys  []ebiten.Key
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	for _, key := range g.keys {
		if key == ebiten.KeyQ {
			os.Exit(0)
		}
	}

	EntityGround.Update()
	EntityDino.Update(g.keys)
	EntityCactus.Update(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	EntityGround.Draw(screen)
	EntityCactus.Draw(screen)
	EntityDino.Draw(screen)

	font := bitmapfont.Gothic12r
	text.Draw(screen, fmt.Sprintf("SCORE: %v", g.Score), font, 30, 30, color.Black)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
