package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Dino struct {
	X                int
	Y                int
	Counter          int
	Tick             int
	SpriteRun0       *ebiten.Image
	SpriteCurrent    *ebiten.Image
	SpriteRun1       *ebiten.Image
	SpriteLose       *ebiten.Image
	SpriteStationary *ebiten.Image
	Options          *ebiten.DrawImageOptions
}

func (d *Dino) Load() error {
	var err error
	d.SpriteRun0, _, err = ebitenutil.NewImageFromFile("assets/dino-run-0.png")
	if err != nil {
		return err
	}
	d.SpriteRun1, _, err = ebitenutil.NewImageFromFile("assets/dino-run-1.png")
	if err != nil {
		return err
	}
	d.SpriteLose, _, err = ebitenutil.NewImageFromFile("assets/dino-lose.png")
	if err != nil {
		return err
	}
	d.SpriteStationary, _, err = ebitenutil.NewImageFromFile("assets/dino-stationary.png")
	if err != nil {
		return err
	}

	d.SpriteCurrent = d.SpriteStationary

	d.X = 50
	d.Y = WINDOW_HEIGHT/2 - 20
	d.Counter = 0
	d.Tick = 10

	return nil
}

func (d *Dino) Draw(screen *ebiten.Image) {
	d.Options = &ebiten.DrawImageOptions{}
	d.Options.GeoM.Translate(float64(d.X), float64(d.Y))

	screen.DrawImage(d.SpriteCurrent, d.Options)
}

func (d *Dino) Update() {
	d.Counter++

	if d.Counter == 2000 {
		d.Counter = 0
		d.Tick = int(math.Max(float64(d.Tick-1), 5))
		println(d.Tick)
	}

	if d.Counter%d.Tick == 0 {
		if d.SpriteCurrent == d.SpriteRun0 {
			d.SpriteCurrent = d.SpriteRun1
		} else {
			d.SpriteCurrent = d.SpriteRun0
		}
	}
}
