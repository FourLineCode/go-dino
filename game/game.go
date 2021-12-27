package game

import (
	"fmt"
	"image/color"
	"log"
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
	Score     int
	HighScore int
	State     GameState
	keys      []ebiten.Key
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	for _, key := range g.keys {
		if key == ebiten.KeyQ || key == ebiten.KeyEscape {
			os.Exit(0)
		}
		if key == ebiten.KeySpace && g.State == GameStateLost {
			if err := LoadEntities(); err != nil {
				log.Fatal("error loading sprites", err)
			}
			g.State = GameStateRunning
			if g.Score > g.HighScore {
				g.HighScore = g.Score
			}
			g.Score = 0
			return nil
		}
	}

	EntityGround.Update(g)
	EntityDino.Update(g)
	EntityCactus.Update(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	EntityGround.Draw(screen)
	EntityCactus.Draw(screen)
	EntityDino.Draw(screen)

	font := bitmapfont.Gothic12r

	if g.State == GameStateRunning {
		text.Draw(screen, fmt.Sprintf("SCORE: %v", g.Score), font, 30, 30, color.Black)
		text.Draw(screen, fmt.Sprintf("HIGH SCORE: %v", g.HighScore), font, 30, 50, color.Black)
	} else {
		text.Draw(screen, fmt.Sprintf("SCORE: %v", g.Score), font, 380, 100, color.Black)
		text.Draw(screen, "PRESS SPACE TO CONTINUE", font, 340, 120, color.Black)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
