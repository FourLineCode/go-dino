package game

import (
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
	options1 *ebiten.DrawImageOptions
	options2 *ebiten.DrawImageOptions
}

func (g *Ground) Load() error {
	var err error
	g.Sprite, _, err = ebitenutil.NewImageFromFile("assets/ground.png")
	if err != nil {
		return err
	}

	g.X1 = 0
	g.X2 = g.Sprite.Bounds().Max.X
	g.Y = WINDOW_HEIGHT / 2

	g.Speed = 5
	g.Counter = 0

	return nil
}

func (g *Ground) Draw(screen *ebiten.Image) {
	g.options1 = &ebiten.DrawImageOptions{}
	g.options2 = &ebiten.DrawImageOptions{}
	g.options1.GeoM.Translate(float64(g.X1), float64(g.Y))
	g.options2.GeoM.Translate(float64(g.X2), float64(g.Y))

	screen.DrawImage(g.Sprite, g.options1)
	screen.DrawImage(g.Sprite, g.options2)
}

func (g *Ground) Update() {
	g.Counter++
	if g.Counter == 1000 {
		g.Counter = 0
		g.Speed = int(math.Min(float64(g.Speed+1), 12))
		println(g.Speed)
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
