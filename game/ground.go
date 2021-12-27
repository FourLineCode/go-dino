package game

import (
	"bytes"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ground struct {
	X1       int
	X2       int
	Y        int
	Speed    int
	Counter  int
	Sprite   *ebiten.Image
	Options1 *ebiten.DrawImageOptions
	Options2 *ebiten.DrawImageOptions
}

func (g *Ground) Load() error {
	var err error
	imageByte, err := files.ReadFile("assets/ground.png")
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(imageByte)
	g.Sprite, _, err = ebitenutil.NewImageFromReader(reader)
	if err != nil {
		return err
	}

	g.X1 = 0
	g.X2 = g.Sprite.Bounds().Max.X
	g.Y = WINDOW_HEIGHT/2 + 50

	g.Speed = PLATFORM_MIN_SPEED
	g.Counter = 0

	return nil
}

func (g *Ground) Draw(screen *ebiten.Image) {
	g.Options1 = &ebiten.DrawImageOptions{}
	g.Options2 = &ebiten.DrawImageOptions{}
	g.Options1.GeoM.Translate(float64(g.X1), float64(g.Y))
	g.Options2.GeoM.Translate(float64(g.X2), float64(g.Y))

	screen.DrawImage(g.Sprite, g.Options1)
	screen.DrawImage(g.Sprite, g.Options2)
}

func (g *Ground) Update(game *Game) {
	if game.State == GameStateLost {
		return
	}

	g.Counter++
	if g.Counter == SPEED_INCREASE_INTERVAL {
		g.Counter = 0
		g.Speed = int(math.Min(float64(g.Speed+1), PLATFORM_MAX_SPEED))
	}

	g.X1 -= g.Speed
	g.X2 -= g.Speed

	if g.X1 <= -g.Sprite.Bounds().Max.X {
		g.X1 = g.Sprite.Bounds().Max.X
	}
	if g.X2 <= -g.Sprite.Bounds().Max.X {
		g.X2 = g.Sprite.Bounds().Max.X
	}
}
